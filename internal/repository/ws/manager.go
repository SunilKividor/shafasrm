package ws

import (
	"log"
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/auth"
	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	Clients map[uuid.UUID]*Client
	Rooms   map[uuid.UUID]*Room
	PGRepo  *pgrepo.PGRepo
}

type Room struct {
	MatchID uuid.UUID
	Clients map[*Client]bool
	Bcast   chan models.ChatMessage
}

func NewManager() *Manager {
	pgDBClient := pgdb.GetDBClient()
	pgRepo := pgrepo.NewPGRepo(pgDBClient)
	return &Manager{
		Clients: make(map[uuid.UUID]*Client),
		Rooms:   make(map[uuid.UUID]*Room),
		PGRepo:  pgRepo,
	}
}

func (manager *Manager) ServeWS(c *gin.Context) {
	user_id, err := auth.ExtractIdFromContext(c)
	if err != nil {
		log.Println("error getting id from context")
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "error getting user id from token",
				"error": err.Error(),
			},
		)
		return
	}

	match_id_str := c.Query("match_id")
	match_id := uuid.MustParse(match_id_str)

	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, user_id, match_id)

	manager.RegisterClient(client)
	go client.ReadConn(manager)
	go client.WriteConn()
}

func (m *Manager) RegisterClient(client *Client) {
	m.Clients[client.UserID] = client

	room, exists := m.Rooms[client.MatchID]
	if !exists {
		room = &Room{
			MatchID: client.MatchID,
			Clients: make(map[*Client]bool),
			Bcast:   make(chan models.ChatMessage),
		}
		m.Rooms[client.MatchID] = room
		go room.run()
	}

	room.Clients[client] = true
}

func (m *Manager) UnRegister(client *Client) {
	delete(m.Clients, client.UserID)

	if room, ok := m.Rooms[client.MatchID]; ok {
		delete(room.Clients, client)
		if len(room.Clients) == 0 {
			close(room.Bcast)
			delete(m.Rooms, client.MatchID)
		}
	}

	close(client.Send)
}
