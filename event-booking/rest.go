package gol_eventbooking

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"go-learning.com/learning/event-booking/db"
	"go-learning.com/learning/event-booking/routes"
)

func DbMiddleWare(sqlDB *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn", sqlDB)
		c.Next()
	}
}

func Run() {
	sqlDB := db.InitDB()
	enginePtr := gin.Default()
	enginePtr.Use(DbMiddleWare(sqlDB))
	routes.RegisterRoutes(enginePtr)
	enginePtr.Run(":8080") // localhost:8080
}
