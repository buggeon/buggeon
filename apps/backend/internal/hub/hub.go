package hub

import (
	"sync"
)

type Hub struct {
	mu    sync.RWMutex
	rooms map[string]map[*Client]struct{}

	register   chan *Client
	unregister chan *Client
	broadcast  chan broadcastMessage
}

type broadcastMessage struct {
	RoomID  string
	Payload []byte
	Except  *Client
}

func NewHub() *Hub {
	return &Hub{
		rooms:      make(map[string]map[*Client]struct{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan broadcastMessage),
	}
}

func (h *Hub) Register(c *Client)   { h.register <- c }
func (h *Hub) Unregister(c *Client) { h.unregister <- c }
func (h *Hub) Broadcast(roomID string, msg []byte) {
	h.broadcast <- broadcastMessage{
		RoomID:  roomID,
		Payload: msg,
	}
}
func (h *Hub) BroadcastExcept(roomID string, msg []byte, ex *Client) {
	h.broadcast <- broadcastMessage{
		RoomID:  roomID,
		Payload: msg,
		Except:  ex,
	}
}
func (h *Hub) Run() {

	for {

		select {

		case client := <-h.register:

			if h.rooms[client.RoomID] == nil {
				h.rooms[client.RoomID] = map[*Client]struct{}{}
			}

			h.rooms[client.RoomID][client] = struct{}{}

		case client := <-h.unregister:

			if clients, ok := h.rooms[client.RoomID]; ok {
				if _, exists := clients[client]; exists {
					delete(clients, client)
					close(client.Send)

					if len(clients) == 0 {
						delete(h.rooms, client.RoomID)
					}
				}
			}

		case message := <-h.broadcast:

			for client := range h.rooms[message.RoomID] {

				client.Send <- message.Payload

			}

		}

	}

}
