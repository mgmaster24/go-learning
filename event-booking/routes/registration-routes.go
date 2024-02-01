package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-learning.com/learning/event-booking/models"
	"go-learning.com/learning/event-booking/response"
)

func (sqlDB *SqlDB) register(c *gin.Context) {
	eventId, ok := getIdParam(c)
	if !ok {
		return
	}

	evt, err := models.GetEvent(sqlDB.DB, eventId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not get event from the DB.", err)
		return
	}

	userId := c.GetInt64("userId")
	err = evt.Register(sqlDB.DB, userId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not register user for event.", err)
		return
	}

	response.SuccessWithMsgNoObj(c, http.StatusCreated, "Registered for event")
}

func (sqlDB *SqlDB) unregister(c *gin.Context) {
	eventId, ok := getIdParam(c)
	if !ok {
		return
	}

	evt, ok := canPerfromAction(sqlDB, eventId, c, "unregister")
	if !ok {
		return
	}

	err := evt.Unregiter(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not unregister user from event.", err)
		return
	}

	response.SuccessWithMsgNoObj(c, http.StatusOK, "Unregistered for event")
}

func (sqlDB *SqlDB) getRegistrations(c *gin.Context) {
	regs, err := models.GetRegistrations(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not get events from the DB.", err)
		return
	}
	response.Success(c, http.StatusOK, regs)
}
