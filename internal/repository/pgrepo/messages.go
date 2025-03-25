package pgrepo

import (
	"log"

	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

func (dbClient *PGRepo) AddMessage(message models.ChatMessage) error {
	db := dbClient.PostgresDBClient

	smt := `INSERT INTO messages (match_id,sender_id,content) VALUES ($1,$2,$3)`

	_, err := db.Exec(smt, message.MatchID, message.SenderID, message.Content)
	return err
}

func (dbClient *PGRepo) GetMessages(match_id uuid.UUID) ([]models.Message, error) {
	db := dbClient.PostgresDBClient
	var messages []models.Message
	smt :=
		`
		SELECT m.sender_id,m.content,m.sent_at 
		FROM messages m 
		WHERE m.match_id = $1
		ORDER BY m.sent_at ASC
	`

	rows, err := db.Query(smt, match_id)
	if err != nil {
		return messages, err
	}

	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.SenderID, &msg.Content, &msg.SentAt)
		if err != nil {
			log.Printf("row scan error: %v", err)
			continue
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
