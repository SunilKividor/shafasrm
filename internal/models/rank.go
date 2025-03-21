package models

import "github.com/google/uuid"

type UserRank struct {
	Rank   int       `json:"rank"`
	UserID uuid.UUID `json:"user_id"`
	Points int       `json:"points"`
}
