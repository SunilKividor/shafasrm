package handler

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerBody models.RegisterBody
	err := c.Bind(&registerBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Body passed",
		})
		return
	}

	c.JSON(200, gin.H{
		"username": registerBody.Username,
	})
}
