package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub        *Hub
	conn       *websocket.Conn
	controller WSController
	// Buffered channel of outbound messages.
	send             chan []byte
	actionMarshaller ActionMarshaller
}

type ClientResponse struct {
	client *Client
	data   []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error(err.Error())
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		c.handleMessage(message)
	}
}

func (c *Client) handleMessage(message []byte) {
	var action Action
	json.Unmarshal(message, &action)
	r := c.controller.GetV2(&action)
	responseBytes, _ := json.Marshal(r)
	clientResponse := ClientResponse{c, responseBytes}
	c.hub.response <- clientResponse

	// r := c.controller.RunCommand(&action)
	// responseBytes, err := c.actionMarshaller.Marshal(r)

	// if err != nil {
	// 	slog.Error(err.Error())
	// 	return
	// }
	//
	//	if cmd.Command == CommandTypeUpdate {
	//		c.hub.broadcast <- responseBytes
	//	} else {
	//
	//		response := ClientResponse{c, responseBytes}
	//		c.hub.response <- response
	//	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, c WSController, w http.ResponseWriter, r *http.Request, am ActionMarshaller) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Info(fmt.Sprint(err))
		return
	}
	client := &Client{hub: hub, conn: conn, controller: c, send: make(chan []byte, 256), actionMarshaller: am}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
