package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-learning.com/learning/event-booking/models"
	"go-learning.com/learning/event-booking/response"
)

func (sqlDB *SqlDB) getEvents(c *gin.Context) {
	events, err := models.GetEvents(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not get events from the DB.", err)
		return
	}
	response.Success(c, http.StatusOK, events)
}

func (sqlDB *SqlDB) getEvent(c *gin.Context) {
	id, ok := getIdParam(c)
	if !ok {
		return
	}

	event, err := models.GetEvent(sqlDB.DB, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not get event from the DB.", err)
		return
	}

	response.Success(c, http.StatusOK, event)
}

func (sqlDB *SqlDB) createEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	var event models.Event
	if !shouldBindJSON(c, &event) {
		return
	}

	event.UserId = userId

	err := event.Save(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not save event to DB.", err)
		return
	}

	response.SuccesWithMsg(c, http.StatusCreated, "Event created", "event", event)
}

func (sqlDB *SqlDB) updateEvent(c *gin.Context) {
	var requestEvent models.Event
	if !shouldBindJSON(c, &requestEvent) {
		return
	}

	id, ok := getIdParam(c)
	if !ok {
		return
	}

	retrievedEvent, ok := canPerfromAction(sqlDB, id, c, "update")
	if !ok {
		return
	}

	requestEvent.Id = retrievedEvent.Id
	requestEvent.UserId = retrievedEvent.UserId
	err := requestEvent.Update(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not update the event", err)
		return
	}

	response.SuccesWithMsg(c, http.StatusCreated, "Event updated", "event", requestEvent)
}

func (sqlDB *SqlDB) deleteEvent(c *gin.Context) {
	id, ok := getIdParam(c)
	if !ok {
		return
	}

	retrievedEvent, ok := canPerfromAction(sqlDB, id, c, "delete")
	if !ok {
		return
	}

	err := retrievedEvent.Delete(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not update the event", err)
		return
	}

	response.SuccessWithMsgNoObj(c, http.StatusAccepted, "Event deleted")
}

func canPerfromAction(sqlDB *SqlDB, id int64, c *gin.Context, action string) (*models.Event, bool) {
	retrievedEvent, err := models.GetEvent(sqlDB.DB, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not get event from the DB.", err)
		return nil, false
	}

	userId := c.GetInt64("userId")
	if retrievedEvent.UserId != userId {
		response.Error(
			c,
			http.StatusUnauthorized,
			fmt.Sprintf("Not authroized to perform %s", action),
			fmt.Errorf("unathorized to perform %s on resource", action))
		return nil, false
	}

	return &retrievedEvent, true
}
