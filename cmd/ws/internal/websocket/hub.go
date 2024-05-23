package websocket

type Hub struct {
	clients    map[*Client]bool
	response   chan ClientResponse
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		// All these are unbuffered channels, it might be interesting to
		// consider buffered channels to be able to resist traffic bursts.
		broadcast:  make(chan []byte),
		response:   make(chan ClientResponse),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case clientResponse := <-h.response:
			select {
			case clientResponse.client.send <- clientResponse.data:
			default:
				close(clientResponse.client.send)
				delete(h.clients, clientResponse.client)
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
