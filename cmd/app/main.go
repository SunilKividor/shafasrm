package main

import (
	"log"
	"os"

	"github.com/SunilKividor/shafasrm/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln("Error loading .env")
	}
}

func main() {
	port := os.Getenv("PORT")
	server := server.NewServer(port)
	err := server.RunServer()
	if err != nil {
		log.Fatalf("Error running server: %s", err)
	}
}
