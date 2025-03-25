package handler

import (
	"log"

	"github.com/SunilKividor/shafasrm/internal/repository/ws"
	"github.com/gin-gonic/gin"
)

var manager = ws.NewManager()

func HanldeWSConnection(c *gin.Context) {
	log.Println("connection handle")
	manager.ServeWS(c)
}
