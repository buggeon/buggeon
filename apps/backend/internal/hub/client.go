package hub

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID        string
	UserName      string
	UserAvatarKey string
	Conn          *websocket.Conn
	Send          chan []byte
}

type BroadcastMessage struct {
	ID              string    `json:"id"`
	SenderName      string    `json:"senderName"`
	SenderID        string    `json:"senderId"`
	SenderAvatarKey string    `json:"senderAvatarKey"`
	SentAt          time.Time `json:"sendAt"`
	Content         string    `json:"content"`
}

func NewClient(conn *websocket.Conn, userID, userName, userAvatarKey string) *Client {
	return &Client{
		Conn:          conn,
		UserID:        userID,
		UserName:      userName,
		UserAvatarKey: userAvatarKey,
		Send:          make(chan []byte, 256),
	}
}

func (c *Client) ReadPump(hub *Hub, onMessage func(*Client, []byte)) {

	c.Conn.SetReadLimit(1024)

	defer func() {
		c.Conn.Close()
		hub.unregister <- c
	}()

	for {

		_, msg, err := c.Conn.ReadMessage()

		if err != nil {
			break
		}

		messageObject := BroadcastMessage{
			ID:              hub.name,
			SenderName:      c.UserName,
			SenderID:        c.UserID,
			SenderAvatarKey: c.UserAvatarKey,
			Content:         string(msg),
			SentAt:          time.Now(),
		}

		messageObjectBytes, err := json.Marshal(messageObject)
		if err != nil {
			break
		}

		hub.Broadcast(messageObjectBytes)
		onMessage(c, msg)

	}

}

func (c *Client) WritePump() {

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}

}
