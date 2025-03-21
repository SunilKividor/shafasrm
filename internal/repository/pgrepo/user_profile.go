package pgrepo

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

func (dbClient *PGRepo) CreateUserProfile(user_id uuid.UUID, profile models.UserProfile) error {
	db := dbClient.PostgresDBClient

	smt := `
		INSERT INTO user_profile 
		(user_id,party_move,wild_place,zombie_days,guilty_song,
		ghost_reason,shots_confess,late_food,chaotic_love,
		breakup_power,weak_spot,flirt_rating,into_signal,
		dumb_line,trouble_sign,campus_rumor) 
		VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
		`
	_, err := db.Exec(
		smt,
		user_id,
		profile.PartyMove, profile.WildPlace, profile.ZombieDays,
		profile.GuiltySong, profile.GhostReason, profile.ShotsConfess,
		profile.LateFood, profile.ChaoticLove, profile.BreakupPower,
		profile.WeakSpot, profile.FlirtRating, profile.IntoSignal,
		profile.DumbLine, profile.TroubleSign, profile.CampusRumor,
	)

	return err
}
