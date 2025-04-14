package router

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/authentication"
	"github.com/SunilKividor/shafasrm/internal/handler"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	//api v1
	api := router.Group("/api/v1")

	//health-check
	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": "server running",
			},
		)
	})

	auth := api.Group("/auth")
	{
		auth.POST("/register", handler.RegisterUser)
		auth.POST("/login", handler.LoginUser)
	}

	user := api.Group("/users")
	user.Use(authentication.AuthMiddleware())
	{
		user.POST("/profile", handler.CreateUserProfile) //create profile
		user.PUT("/profile", handler.CreateUserProfile)  //update profile //todo
		user.GET("/profile", handler.CreateUserProfile)  //Get profile //todo

		user.POST("/details", handler.AddUserDetails) //add user details
		user.PUT("/details", handler.AddUserDetails)  //update user details //todo

		user.POST("/photos/pre-signed-url", handler.GetPresignedUploadUrl) //Generate S3 URL
		user.POST("/photos", handler.GetPresignedUploadUrl)                //Store Photo URL //todo

		user.POST("/swipes", handler.Swipe)                          //Record a Swipe
		user.GET("/swipes/feed", handler.SwipeFeed)                  //Get Swipable Profiles
		user.POST("/matches", handler.CreateMatch)                   //Create a match
		user.GET("/matches", handler.GetMatches)                     //Get matches
		user.GET("/matches/:match_id/messages", handler.GetMessages) //chat history

		user.DELETE("", handler.DeleteUser)
	}

	ws := api.Group("/ws")
	ws.Use(authentication.AuthMiddleware())
	{
		ws.GET("/chat", handler.HanldeWSConnection)
	}
}
