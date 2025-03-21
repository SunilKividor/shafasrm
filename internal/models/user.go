package models

import "github.com/google/uuid"

type UserID struct {
	ID uuid.UUID `json:"id"`
}

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

type Match struct {
	UserID uuid.UUID `json:"user_id"`
}

type Matches struct {
	UserIDs []uuid.UUID `json:"user_ids"`
}
