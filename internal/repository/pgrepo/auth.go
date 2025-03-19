package pgrepo

import (
	"database/sql"

	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

type PGRepo struct {
	PostgresDBClient *sql.DB
}

func NewPGRepo(dbClient *sql.DB) *PGRepo {
	return &PGRepo{
		PostgresDBClient: dbClient,
	}
}

func (dbClient *PGRepo) RegisterUser(user models.RegisterRequestBody) (uuid.UUID, error) {
	var id uuid.UUID
	db := dbClient.PostgresDBClient

	smt := `
		INSERT INTO users 
		(name,username,password,email,phone,gender,birthday,location,religion,department,stream,degree) 
		VALUES
		($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) 
		RETURNING id
		`
	err := db.QueryRow(smt, user.Name, user.Username, user.Password, user.Email, user.Phone, user.Gender, user.Birthday, user.Location, user.Religion, user.Department, user.Stream, user.Degree).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, err
}

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
