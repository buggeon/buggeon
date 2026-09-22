// Buggeon - SelfHosted service for bug and task tracking
// Copyright (C) 2026 DEVE corp.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package handlers

import (
	"buggeon/internal/dto"
	"buggeon/internal/hub"
	"buggeon/internal/services"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatHandler struct {
	mu             sync.Mutex
	hubs           map[string]*hub.Hub
	messageService *services.MessageService
	tokenService   *services.TokenService
	userService    *services.UserService
	upgrader       websocket.Upgrader
}

func NewChatHandler(
	messageService *services.MessageService,
	tokenService *services.TokenService,
	userService *services.UserService,
) *ChatHandler {

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return &ChatHandler{
		tokenService:   tokenService,
		hubs:           make(map[string]*hub.Hub),
		messageService: messageService,
		upgrader:       upgrader,
		userService:    userService,
	}
}

// ServeWS godoc
// @Summary      WebSocket chat for a card
// @Description  Opens a WebSocket connection to receive and send messages in a card. Token is passed as a query parameter because browsers cannot set custom headers on WebSocket connections.
// @Tags         websocket
// @Param        cardId  path   string  true  "Card ID"
// @Param        token   query  string  true  "Access token (JWT)"
// @Success      101     {string}  string  "Switching Protocols"
// @Failure      401     {object}  map[string]string  "Unauthorized"
// @Failure      500     {object}  map[string]string  "Failed to upgrade connection"
// @Router       /ws/{cardId} [get]
func (h *ChatHandler) ServeWS(c *gin.Context) {

	cardID := c.Param("cardId")
	token := c.Query("token")

	claims, err := h.tokenService.ValidateAccessToken(token)

	if err != nil {
		c.Status(401)
		return
	}

	userData, err := h.userService.GetUser(c, claims.UserID)

	if err != nil {
		c.Status(400)
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.Status(500)
		log.Println("Failed to upgrade connection")
		return
	}

	client := hub.NewClient(conn, claims.UserID, userData.Name, userData.AvatarKey)

	room := h.getOrCreateHub(cardID)

	room.Register(client)

	go client.WritePump()
	go client.ReadPump(room, func(cl *hub.Client, msg []byte) {
		_, err := h.messageService.CreateMessage(c, dto.NewMessageDto{
			SenderID: cl.UserID,
			CardID:   cardID,
			Content:  string(msg),
		})

		if err != nil {
			fmt.Println(err)
		}
	})

}

func (h *ChatHandler) getOrCreateHub(cardID string) *hub.Hub {

	h.mu.Lock()
	defer h.mu.Unlock()

	if room, ok := h.hubs[cardID]; ok {
		return room
	}

	room := hub.NewHub(cardID)
	h.hubs[cardID] = room

	go room.Run()

	return room

}
