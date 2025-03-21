package pgrepo

import "github.com/google/uuid"

const (
	ranked   = "ranked"
	unranked = "unranked"
)

func (dbClient *PGRepo) AddNewUserToRanking(user_id uuid.UUID) error {
	smt := `INSERT INTO ranking(user_id,status) VALUES ($1,$2)`

	db := dbClient.PostgresDBClient

	_, err := db.Exec(smt, user_id, unranked)
	return err
}
