package ws

import (
	"log"

	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	UserID  uuid.UUID
	MatchID uuid.UUID
	Send    chan models.ChatMessage
}

func newClient(conn *websocket.Conn, user_id uuid.UUID, match_id uuid.UUID) *Client {
	return &Client{
		Conn:    conn,
		UserID:  user_id,
		MatchID: match_id,
		Send:    make(chan models.ChatMessage, 10),
	}
}

func (c *Client) ReadConn(m *Manager) {
	defer func() {
		m.UnRegister(c)
		c.Conn.Close()
	}()
	var msg models.ChatMessage
	for {
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("read error: %v", err)
			}
			break
		}

		//insert msg in DB
		pgDBClient := pgdb.GetDBClient()
		postgresRepo := pgrepo.NewPGRepo(pgDBClient)

		msg.MatchID = c.MatchID
		msg.SenderID = c.UserID

		err = postgresRepo.AddMessage(msg)
		if err != nil {
			log.Panicln("Error adding message to the database")
		}

		if room, ok := m.Rooms[c.MatchID]; ok {
			room.Bcast <- msg
		}
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
