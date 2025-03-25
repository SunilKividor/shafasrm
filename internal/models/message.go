package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatMessage struct {
	MatchID  uuid.UUID `json:"match_id"`
	SenderID uuid.UUID `json:"sender_id"`
	Content  string    `json:"content"`
}

type Message struct {
	SenderID uuid.UUID `json:"sender_id" db:"sender_id"`
	Content  string    `json:"content" db:"content"`
	SentAt   time.Time `json:"sent_at" db:"sent_at"`
}
