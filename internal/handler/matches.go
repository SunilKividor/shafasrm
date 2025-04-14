package handler

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/authentication"
	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/gin-gonic/gin"
)

func CreateMatch(c *gin.Context) {
	var match models.Match

	err := c.ShouldBind(&match)
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

	id, err := authentication.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"msg":   "error getting user id",
				"error": err.Error(),
			},
		)
		return
	}

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	err = postgresRepo.CreateNewMatch(id, match)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "error creating a match",
				"error": err.Error(),
			},
		)
		return
	}
}

func GetMatches(c *gin.Context) {

	id, err := authentication.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"msg":   "error getting user id",
				"error": err.Error(),
			},
		)
		return
	}

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	matches, err := postgresRepo.GetMatches(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "error getting matches",
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, matches)
}
