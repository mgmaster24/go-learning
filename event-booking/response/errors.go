package response

import "github.com/gin-gonic/gin"

func Error(context *gin.Context, status int, msg string, err error) {
	context.JSON(status, gin.H{"message": msg, "error": err.Error()})
}
