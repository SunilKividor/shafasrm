package handler

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMessages(c *gin.Context) {
	match_id_str := c.Param("match_id")
	match_id := uuid.MustParse(match_id_str)

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	messages, err := postgresRepo.GetMessages(match_id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "error getting messages",
				"error": err.Error(),
			},
		)
	}

	c.JSON(http.StatusOK, messages)
}
