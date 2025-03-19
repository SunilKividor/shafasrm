package pgdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	postgresDBClient *sql.DB
)

func GetDBClient() *sql.DB {
	return postgresDBClient
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConfig(host string, port int, user string, password string, dbName string, sslMode string) *PostgresConfig {
	return &PostgresConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}
}

func (cfg *PostgresConfig) RunPostgresql() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	postgresDBClient = db
	return nil
}
