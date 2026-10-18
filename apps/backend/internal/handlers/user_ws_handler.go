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
	"buggeon/internal/hub"
	"buggeon/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type UserWSHandler struct {
	hub          *hub.UserHub
	tokenService *services.TokenService
	userService  *services.UserService
	upgrader     websocket.Upgrader
}

func NewUserWSHandler(
	tokenService *services.TokenService,
	userService *services.UserService,
	hub *hub.UserHub,
) *UserWSHandler {

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return &UserWSHandler{
		userService:  userService,
		tokenService: tokenService,
		hub:          hub,
		upgrader:     upgrader,
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
func (h *UserWSHandler) ServeWS(c *gin.Context) {

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

	h.hub.Join(client, userData.ID.Hex())

	go client.WritePump()
	// go client.ReadPump(func(cl *hub.Client, msg []byte) {
	// 	notification, err := h.notificationService.NewNotification(
	// 		c,
	// 		client.UserID,

	// 	)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	messageObject := hub.BroadcastMessage{
	// 		ID:              notification.ID.Hex(),
	// 		RoomID:          cl.RoomID,
	// 		SenderName:      cl.UserName,
	// 		SenderID:        cl.UserID,
	// 		SenderAvatarKey: cl.UserAvatarKey,
	// 		Content:         string(msg),
	// 		SentAt:          time.Now(),
	// 	}

	// 	messageObjectBytes, err := json.Marshal(messageObject)

	// 	h.hub.Broadcast(messageObjectBytes, userData.ID.Hex())

	// }, func(c *hub.Client) {
	// 	h.hub.Leave(c)
	// })

}
