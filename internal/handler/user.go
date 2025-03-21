package handler

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "empty string",
			"error": "error",
		})
		return
	}

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	user_id, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "could not parse id",
			"error": err.Error(),
		})
		return
	}

	err = postgresRepo.DeleteUser(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "could not delete the user",
			"error": err.Error(),
		})
		return
	}
}
