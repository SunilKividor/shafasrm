package database

import (
	"os"
	"strconv"

	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
)

func InitPostgresql() error {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSLMODE")
	port := os.Getenv("POSTGRES_PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		portInt = 5432
	}
	cfg := pgdb.NewPostgresConfig(
		host,
		portInt,
		user,
		password,
		dbName,
		sslMode,
	)
	err = cfg.RunPostgresql()
	return err
}
