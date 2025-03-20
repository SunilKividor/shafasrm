package pgrepo

import (
	"database/sql"
)

type PGRepo struct {
	PostgresDBClient *sql.DB
}

func NewPGRepo(dbClient *sql.DB) *PGRepo {
	return &PGRepo{
		PostgresDBClient: dbClient,
	}
}
