package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	RoomID        string
	UserID        string
	UserName      string
	UserAvatarKey string
	Conn          *websocket.Conn
	Send          chan []byte
}

type BroadcastMessage struct {
	ID              string    `json:"id"`
	RoomID          string    `json:""`
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

func (c *Client) ReadPump(onMessage func(*Client, []byte), onClose func(*Client)) {

	c.Conn.SetReadLimit(1024)

	defer func() {
		c.Conn.Close()
		onClose(c)
	}()

	for {

		_, msg, err := c.Conn.ReadMessage()

		if err != nil {
			break
		}

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
