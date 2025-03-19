package server

import (
	"log"

	"github.com/SunilKividor/shafasrm/internal/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (server *Server) RunServer() error {
	port := server.Port
	if port == "" {
		port = "8080"
		log.Println("Defaulting to port:8080")
	}
	log.Println(port)
	r := gin.Default()

	router.Router(r)

	err := r.Run(":" + port)
	return err
}
