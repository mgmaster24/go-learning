package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-learning.com/learning/event-booking/response"
)

type SqlDB struct {
	DB *sql.DB
}

func RegisterRoutes(enginePtr *gin.Engine, sqlDB *SqlDB) {
	enginePtr.GET("/events", sqlDB.getEvents)
	enginePtr.GET("/registrations", sqlDB.getRegistrations)
	enginePtr.GET("/events/:id", sqlDB.getEvent)

	authenticatedGroup := enginePtr.Group("/")
	authenticatedGroup.Use(authenticate)
	authenticatedGroup.POST("/events", sqlDB.createEvent)
	authenticatedGroup.PUT("/events/:id", sqlDB.updateEvent)
	authenticatedGroup.DELETE("/events/:id", sqlDB.deleteEvent)
	authenticatedGroup.POST("events/:id/register", sqlDB.register)
	authenticatedGroup.DELETE("events/:id/register", sqlDB.unregister)

	enginePtr.POST("/signup", sqlDB.createUser)
	enginePtr.POST("/login", sqlDB.login)
}

func getIdParam(context *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		response.Error(context, http.StatusBadRequest, "Could parse id param from url.", err)
		return -1, false
	}

	return id, true
}

func shouldBindJSON(context *gin.Context, obj any) bool {
	err := context.ShouldBindJSON(&obj)
	if err != nil {
		response.Error(context, http.StatusBadRequest, "Could not parse request body.", err)
		return false
	}

	return true
}
