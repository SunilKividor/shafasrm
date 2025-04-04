package models

import "github.com/google/uuid"

// this stores the options
type UserProfile struct {
	UserID       uuid.UUID `json:"user_id"`
	PartyMove    int       `json:"party_move"`
	WildPlace    int       `json:"wild_place"`
	ZombieDays   int       `json:"zombie_days"`
	GuiltySong   string    `json:"guilty_song"`
	GhostReason  int       `json:"ghost_reason"`
	ShotsConfess int       `json:"shots_confess"`
	LateFood     int       `json:"late_food"`
	ChaoticLove  string    `json:"chaotic_love"`
	BreakupPower int       `json:"breakup_power"`
	WeakSpot     string    `json:"weak_spot"`
	FlirtRating  int       `json:"flirt_rating"`
	IntoSignal   int       `json:"into_signal"`
	DumbLine     string    `json:"dumb_line"`
	TroubleSign  int       `json:"trouble_sign"`
	CampusRumor  string    `json:"campus_rumor"`
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
