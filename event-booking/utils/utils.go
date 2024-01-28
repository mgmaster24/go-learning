package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdFromParam(context *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	return id, err
}
