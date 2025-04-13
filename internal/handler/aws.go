package handler

import (
	"log"
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/auth"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/aws"
	"github.com/gin-gonic/gin"
)

func GetPresignedUrl(c *gin.Context) {
	id, err := auth.ExtractIdFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "could not get the id",
				"error": err.Error(),
			},
		)
		return
	}
	log.Println(id)
	var presignedUrlReq models.PreSignedUrlReq
	err = c.ShouldBind(&presignedUrlReq)
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

	log.Println(presignedUrlReq)
	url, err := aws.GetPreSignedUrl(id, presignedUrlReq)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"msg":   "unable to get presigned url",
				"error": err.Error(),
			},
		)
		return
	}

	var res models.PreSignedUrlRes
	res.Url = url
	c.JSON(http.StatusOK, res)
}
