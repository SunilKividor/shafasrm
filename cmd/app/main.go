package main

import (
	"log"
	"os"

	"github.com/SunilKividor/shafasrm/internal/database"
	"github.com/SunilKividor/shafasrm/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Panicln("Error loading .env file")
		}
	}
}

func main() {
	err := database.InitPostgresql()
	if err != nil {
		log.Fatalf("Error starting the database: %s", err.Error())
	}

	port := os.Getenv("PORT")
	log.Println(port)
	server := server.NewServer(port)
	err = server.RunServer()
	if err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
