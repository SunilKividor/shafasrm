package handler

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/auth"
	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/gin-gonic/gin"
)

func CreateUserProfile(c *gin.Context) {
	var profile models.UserProfile

	err := c.ShouldBind(&profile)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"msg":   "invalid request body",
				"error": err.Error(),
			},
		)
		return
	}

	id, err := auth.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "error getting user id from token",
				"error": err.Error(),
			},
		)
		return
	}

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	err = postgresRepo.CreateUserProfile(id, profile)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "error creating user profle",
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
