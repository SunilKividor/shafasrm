package pgrepo

import (
	"github.com/google/uuid"
)

func (dbClient *PGRepo) AddRefreshToken(refreshToken string, user_id uuid.UUID) error {
	db := dbClient.PostgresDBClient

	smt := `
		INSERT INTO auth
		(user_id,refresh_token)
		VALUES
		($1,$2)
		`

	_, err := db.Exec(smt, user_id, refreshToken)
	return err
}

func (dbClient *PGRepo) UpdateRefreshToken(refreshToken string, user_id uuid.UUID) error {
	db := dbClient.PostgresDBClient

	smt := `
		UPDATE auth SET 
		refresh_token = $1 
		WHERE user_id = $2
		`

	_, err := db.Exec(smt, refreshToken, user_id)
	return err
}

func (dbClient *PGRepo) GetRefreshToken(refreshToken string, user_id uuid.UUID) (string, error) {
	var refresh_token string
	db := dbClient.PostgresDBClient

	smt := `
		SELECT refreshToken 
		FROM auth 
		WHERE user_id = $1 AND refresh_token = $2
		`

	err := db.QueryRow(smt, user_id).Scan(&refreshToken)
	return refresh_token, err

}
