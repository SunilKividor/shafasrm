package router

import (
	"github.com/SunilKividor/shafasrm/internal/auth"
	"github.com/SunilKividor/shafasrm/internal/handler"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	//routers
	authRouter := router.Group("auth")
	userRouter := router.Group("user")

	//adding middleware to groups
	userRouter.Use(auth.AuthMiddleware())

	//auth routes
	authRouter.POST("/register", handler.RegisterUser)
	authRouter.POST("/login", handler.LoginUser)

	//user routes
	userRouter.POST("/profile/create", handler.CreateUserProfile)
	userRouter.POST("/matches/create", handler.CreateMatch)
	userRouter.GET("/matches", handler.GetMatches)
	userRouter.POST("/swipe", handler.AddSwipe)
	userRouter.DELETE("/", handler.DeleteUser)
}
