package handler

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/auth"
	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/gin-gonic/gin"
)

func AddSwipe(c *gin.Context) {
	var swipeReq models.SwipeReq

	err := c.ShouldBind(&swipeReq)
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

	var swipe models.Swipe
	swipe.SwiperID = id
	swipe.SwipedID = swipeReq.SwipedID
	swipe.Action = swipeReq.Action

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	err = postgresRepo.AddSwipeAction(swipe)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "error adding swipe",
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
