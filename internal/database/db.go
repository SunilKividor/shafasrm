package database

import (
	"os"
	"strconv"

	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
)

func InitPostgresql() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")
	port := os.Getenv("DB_PORT")
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
