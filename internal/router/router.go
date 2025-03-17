package router

import (
	"github.com/SunilKividor/shafasrm/internal/handler"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	authRouter := router.Group("users")
	//auth
	authRouter.POST("/register", handler.Register)
}
