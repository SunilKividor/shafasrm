package router

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/auth"
	"github.com/SunilKividor/shafasrm/internal/handler"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	//health-check
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": "server running",
			},
		)
	})

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
	userRouter.POST("/details/create", handler.AddUserDetails)
	userRouter.POST("/matches/create", handler.CreateMatch)
	userRouter.GET("/matches", handler.GetMatches)
	userRouter.GET("/matches/:match_id/messages", handler.GetMessages)
	userRouter.POST("/swipe", handler.Swipe)
	userRouter.GET("/swipe/feed", handler.SwipeFeed)
	userRouter.DELETE("/", handler.DeleteUser)

	//aws presigned user routes
	userRouter.POST("/aws/generate-presigned-url", handler.GetPresignedUrl)

	router.GET("/ws", handler.HanldeWSConnection)
}
