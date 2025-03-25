package ws

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	UserID  uuid.UUID
	MatchID string
	Send    chan Message
}

func newClient(conn *websocket.Conn, user_id uuid.UUID, match_id string) *Client {
	return &Client{
		Conn:    conn,
		UserID:  user_id,
		MatchID: match_id,
		Send:    make(chan Message, 10),
	}
}

func (c *Client) ReadConn(m *Manager) {
	defer func() {
		m.UnRegister(c)
		c.Conn.Close()
	}()
	var msg Message
	for {
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("read error: %v", err)
			}
			break
		}

		msg.SenderID = c.UserID

		if room, ok := m.Rooms[c.MatchID]; ok {
			room.Bcast <- msg
		}

		//insert msg in DB
	}
}

func (c *Client) WriteConn() {
	defer c.Conn.Close()
	log.Printf("writePump started for %v", c.UserID)
	for msg := range c.Send {
		log.Printf("Sending to %v: %v", c.UserID, msg.Content)
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("write error: %v", err)
			break
		}
	}
	log.Printf("writePump stopped for %v", c.UserID)
}
