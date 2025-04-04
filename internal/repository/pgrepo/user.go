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
		(name,password,email) 
		VALUES
		($1,$2,$3) 
		RETURNING id
		`
	err := db.QueryRow(smt, user.Name, user.Password, user.Email).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, err
}

func (dbClient *PGRepo) AddUserDetails(id uuid.UUID, user models.UserDetails) error {
	db := dbClient.PostgresDBClient

	smt := `
		INSERT INTO user_details
		(user_id,phone,gender,birthday,location,religion,department,stream,degree) 
		VALUES
		($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`
	_, err := db.Exec(smt, id, user.Phone, user.Gender, user.Birthday, user.Location, user.Religion, user.Department, user.Stream, user.Degree)
	return err
}

func (dbClient *PGRepo) GetIDPasswordQuery(email string) (uuid.UUID, string, error) {
	var password string
	var id uuid.UUID
	db := dbClient.PostgresDBClient
	smt := `SELECT id,password FROM users WHERE email = $1`
	err := db.QueryRow(smt, email).Scan(&id, &password)
	if err != nil {
		return id, "", err
	}
	return id, password, nil
}

func (dbClient *PGRepo) DeleteUser(user_id uuid.UUID) error {
	db := dbClient.PostgresDBClient
	smt := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(smt, user_id)
	return err
}
