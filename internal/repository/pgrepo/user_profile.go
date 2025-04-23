package pgrepo

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

func (dbClient *PGRepo) CreateUserProfile(user_id uuid.UUID, profile models.UserProfile) error {
	db := dbClient.PostgresDBClient

	smt := `
		INSERT INTO user_profile 
		(user_id,party_move,guilty_song,
		shots_confess,chaotic_love,
		flirt_rating) 
		VALUES 
		($1,$2,$3,$4,$5,$6)
		`
	_, err := db.Exec(
		smt,
		user_id,
		profile.PartyMove,
		profile.GuiltySong, profile.ShotsConfess,
		profile.ChaoticLove,
		profile.FlirtRating,
	)

	return err
}
