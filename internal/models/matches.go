package models

import "github.com/google/uuid"

type Match struct {
	UserID uuid.UUID `json:"user_id"`
}

type Matches struct {
	UserIDs []uuid.UUID `json:"user_ids"`
}
