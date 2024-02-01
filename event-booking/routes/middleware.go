package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-learning.com/learning/event-booking/utils"
)

func authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	parsedToken, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	userId, err := utils.GetUserIdFromToken(parsedToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
