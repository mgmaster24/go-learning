package routes

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-learning.com/learning/event-booking/response"
)

func RegisterRoutes(enginePtr *gin.Engine) {
	enginePtr.GET("/events", getEvents)
	enginePtr.GET("/events/:id", getEvent)
	enginePtr.POST("/events", createEvent)
	enginePtr.PUT("events/:id", updatEvent)
	enginePtr.DELETE("events/:id", deleteEvent)
}

func getIdParam(context *gin.Context) int64 {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		response.Error(context, http.StatusBadRequest, "Could parse id param from url.", err)
	}

	return id
}

func getDBConn(c *gin.Context) (*sql.DB, error) {
	db, ok := c.MustGet("dbConn").(*sql.DB)
	if !ok {
		return nil, errors.New("No database connection in middleware")
	}

	return db, nil
}
