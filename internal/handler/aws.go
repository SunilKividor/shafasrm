package handler

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/SunilKividor/shafasrm/internal/authentication"
	"github.com/SunilKividor/shafasrm/internal/configs"
	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/aws"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
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

	contentType := c.Query("content_type")
	if contentType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":          "content type not found",
			"error":        "invalid req",
			"content-type": contentType,
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

func GetAllUserPhotos(c *gin.Context) {
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

	//get the photo keys from db
	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)
	photoKeys, err := postgresRepo.GetPhotos(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error getting photo keys",
			"error": err.Error(),
		})
		return
	}
	var keys []string
	for _, key := range photoKeys {
		keys = append(keys, key.Key)
	}
	urls, err := multiPreSignedGetUrls(keys)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error getting pre-signed photo urls",
			"error": err.Error(),
		})
		return
	}
	for i, url := range urls {
		photoKeys[i].Key = url
	}

	c.JSON(http.StatusOK, photoKeys)
}

func multiPreSignedGetUrls(photoKeys []string) ([]string, error) {
	awsDefaultConfig, err := configs.DefaultConfig()
	if err != nil {
		return nil, err
	}

	bucketName := os.Getenv("S3BUCKETNAME")
	ctx := context.TODO()
	preSignUrlService, err := aws.NewPresignS3Service(awsDefaultConfig, bucketName)
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, key := range photoKeys {
		url, err := preSignUrlService.GenerateDownloadUrl(ctx, key)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		photoUrl := url

		urls = append(urls, photoUrl)
	}

	return urls, nil
}

func StorePhotoKey(c *gin.Context) {
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
	var body models.PhotoObject
	err = c.ShouldBind(&body)
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

	//check if the key exists in s3
	exists, err := preSignUrlService.VerifyObjectExists(ctx, body.Key)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error verifying key",
			"error": err.Error(),
		})
		return
	}

	//continue uploading to db
	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	err = postgresRepo.StoreNewPhotoKey(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error Storing photo key",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
