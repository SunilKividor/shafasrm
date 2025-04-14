package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/SunilKividor/shafasrm/internal/authentication"
	"github.com/SunilKividor/shafasrm/internal/configs"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/aws"
	"github.com/gin-gonic/gin"
)

func GetPresignedUploadUrl(c *gin.Context) {
	id, err := authentication.ExtractIdFromContext(c)
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

	// var presignedUrlReq models.PreSignedUrlReq
	// err = c.ShouldBind(&presignedUrlReq)
	// if err != nil {
	// 	c.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"msg":   "invalid request body",
	// 			"error": err.Error(),
	// 		},
	// 	)
	// 	return
	// }

	contentType := c.Query("contentType")
	if contentType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "content type not found",
			"error": "invalid req",
		})
		return
	}

	awsDefaultConfig, err := configs.DefaultConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "aws config error",
			"error": err.Error(),
		})
		return
	}

	bucketName := os.Getenv("S3BUCKETNAME")
	ctx := context.TODO()
	preSignUrlService, err := aws.NewPresignS3Service(awsDefaultConfig, bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "PreSignUrl sercvice error",
			"error": err.Error(),
		})
		return
	}

	url, key, err := preSignUrlService.GenerateUploadUrl(ctx, id, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error generating upload url",
			"error": err.Error(),
		})
		return
	}

	var res models.PreSignedUrlRes
	res.Url = url
	res.Key = key
	c.JSON(http.StatusOK, res)
}
