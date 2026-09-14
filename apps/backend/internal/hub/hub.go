package hub

type Hub struct {
	name       string
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub(name string) *Hub {

	h := &Hub{
		name:       name,
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	return h
}

func (h *Hub) Register(c *Client)   { h.register <- c }
func (h *Hub) Unregister(c *Client) { h.unregister <- c }
func (h *Hub) Broadcast(msg []byte) { h.broadcast <- msg }

func (h *Hub) Run() {

	for {
		select {

		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				select {

				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}

		}

	}

}
