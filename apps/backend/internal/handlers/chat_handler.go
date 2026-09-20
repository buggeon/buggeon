package handlers

import (
	"buggeon/internal/cache"
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
	upgrader       websocket.Upgrader
	cache          *cache.UserCache
}

func NewChatHandler(
	messageService *services.MessageService,
	tokenService *services.TokenService,
	cache *cache.UserCache,
) *ChatHandler {

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return &ChatHandler{
		tokenService:   tokenService,
		hubs:           make(map[string]*hub.Hub),
		messageService: messageService,
		upgrader:       upgrader,
		cache:          cache,
	}
}

func (h *ChatHandler) ServeWS(c *gin.Context) {

	cardID := c.Param("cardId")
	token := c.Query("token")

	claims, err := h.tokenService.ValidateAccessToken(token)

	if err != nil {
		c.Status(401)
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.Status(500)
		log.Println("Failed to upgrade connection")
		return
	}

	userData := h.cache.Get(claims.UserID)

	client := hub.NewClient(conn, claims.UserID, userData.Name, userData.AvatarUrl)

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
