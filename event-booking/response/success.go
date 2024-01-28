package response

import "github.com/gin-gonic/gin"

func Success(context *gin.Context, statusCode int, obj any) {
	context.JSON(statusCode, obj)
}

func SuccesWithMsg(context *gin.Context, statusCode int, msg string, objKey string, obj any) {
	context.JSON(statusCode, gin.H{"message": msg, objKey: obj})
}

func SuccessWithMsgNoObj(context *gin.Context, statusCode int, msg string) {
	context.JSON(statusCode, gin.H{"message": msg})
}
