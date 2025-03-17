package database

import (
	"os"
	"strconv"

	"github.com/SunilKividor/shafasrm/internal/database/postgres"
)

func InitPostgresql() error {
	host := os.Getenv("HOST")
	user := os.Getenv("User")
	password := os.Getenv("Password")
	dbName := os.Getenv("DBName")
	sslMode := os.Getenv("SSLMode")
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		portInt = 5432
	}
	cfg := postgres.NewPostgresConfig(
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
