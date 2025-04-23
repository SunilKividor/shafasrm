package models

import "github.com/google/uuid"

// this stores the options
type UserProfile struct {
	UserID       uuid.UUID `json:"user_id"`
	PartyMove    int       `json:"party_move"`
	GuiltySong   int       `json:"guilty_song"`
	ShotsConfess int       `json:"shots_confess"`
	ChaoticLove  int       `json:"chaotic_love"`
	FlirtRating  int       `json:"flirt_rating"`
}

type UserDetails struct {
	Phone      string `json:"phone"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Location   string `json:"location"`
	Religion   string `json:"religion"`
	Department string `json:"department"`
	Stream     string `json:"stream"`
	Degree     string `json:"degree"`
}
