package pgrepo

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

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

func (dbClient *PGRepo) GetIDPasswordQuery(username string) (uuid.UUID, string, error) {
	var password string
	var id uuid.UUID
	db := dbClient.PostgresDBClient
	smt := `SELECT id,password FROM users WHERE username = $1`
	err := db.QueryRow(smt, username).Scan(&id, &password)
	if err != nil {
		return id, "", err
	}
	return id, password, nil
}
