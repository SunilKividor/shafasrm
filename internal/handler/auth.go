package handler

import (
	"net/http"

	"github.com/SunilKividor/shafasrm/internal/auth"
	"github.com/SunilKividor/shafasrm/internal/database/pgdb"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/SunilKividor/shafasrm/internal/repository/pgrepo"
	"github.com/SunilKividor/shafasrm/internal/util"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var body models.LoginRequestBody
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "invalid request body",
			"error": err.Error(),
		})
		return
	}

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	id, password, err := postgresRepo.GetIDPasswordQuery(body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "error getting Hashed Password",
			"error": err.Error(),
		})
		return
	}

	isVerified := util.ComparePassword(password, body.Password)
	if !isVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "password did not match",
			"error": "Error validating Password",
		})
		return
	}

	accessToken, refreshToken, err := auth.GenerateTokens(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Generating Tokens",
			"error": err.Error(),
		})
		return
	}

	err = postgresRepo.UpdateRefreshToken(refreshToken, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Updating Refresh Token",
			"error": err.Error(),
		})
		return
	}

	var res models.AuthResBody
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(
		http.StatusOK,
		res,
	)
}

func RegisterUser(c *gin.Context) {
	var registerReqBody models.RegisterRequestBody
	err := c.ShouldBind(&registerReqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "invalid request body",
			"error": err.Error(),
		})
		return
	}

	if !util.ValidateEmail(registerReqBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "only emails with @srmist.edu.in is allowed",
		})
		return
	}

	hashedPassword, err := util.HashPassword(registerReqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "password hashing",
			"error": err.Error(),
		})
		return
	}

	registerReqBody.Password = hashedPassword

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)
	userID, err := postgresRepo.RegisterUser(registerReqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error creating new User",
			"error": err.Error(),
		})
		return
	}

	err = postgresRepo.AddNewUserToRanking(userID)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"msg":   "error adding user to ranking",
			"error": err.Error(),
		})
	}

	accessToken, refreshToken, err := auth.GenerateTokens(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error creating access and refresh tokens",
			"error": err.Error(),
		})
		return
	}

	err = postgresRepo.AddRefreshToken(refreshToken, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error adding refresh token to db",
			"error": err.Error(),
		})
		return
	}

	var res models.AuthResBody
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	c.JSON(http.StatusOK, res)

}

func RefreshToken(c *gin.Context) {
	var body models.RefreshreqModel
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "invalid request body",
			"error": err.Error(),
		})
		return
	}
	refreshToken := body.RefreshToken

	id, err := auth.ExtractIdFromToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "error extracting id from token",
			"error": err.Error(),
		})
		return
	}

	pgDBClient := pgdb.GetDBClient()
	postgresRepo := pgrepo.NewPGRepo(pgDBClient)

	refresh_token, err := postgresRepo.GetRefreshToken(refreshToken, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Refresh token not found",
			"error": err.Error(),
		})
		return
	}

	ok := util.CompareRefreshToken(refreshToken, refresh_token)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Refresh tokens do not match",
		})
		return
	}

	accessToken, err := auth.RefreshAccessToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "error refreshing token",
			"error": err.Error(),
		})
		return
	}

	var res models.AuthResBody
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusOK, res)
}
