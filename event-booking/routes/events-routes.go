package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-learning.com/learning/event-booking/models"
	"go-learning.com/learning/event-booking/response"
)

func getEvents(c *gin.Context) {
	db, err := getDBConn(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not establish DB connection", err)
	}
	events, err := models.GetEvents(db)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not get events from the DB.", err)
		return
	}
	response.Success(c, http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	db, err := getDBConn(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not establish DB connection", err)
	}

	id := getIdParam(c)
	event, err := models.GetEvent(db, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not get event from the DB.", err)
		return
	}

	response.Success(c, http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	db, err := getDBConn(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not establish DB connection", err)
	}
	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Could not parse request body.", err)
		return
	}

	event.UserId = uuid.New()
	err = event.Save(db)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not save event to DB.", err)
		return
	}

	response.SuccesWithMsg(c, http.StatusCreated, "Event created", "event", event)
}

func updatEvent(c *gin.Context) {
	db, err := getDBConn(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not establish DB connection", err)
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Could not parse request body.", err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Could parse id param from url.", err)
		return
	}

	err = event.Update(db, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not update the event", err)
		return
	}

	response.SuccesWithMsg(c, http.StatusCreated, "Event updated", "event", event)
}

func deleteEvent(c *gin.Context) {
	db, err := getDBConn(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not establish DB connection", err)
	}

	id := getIdParam(c)
	err = models.Delete(db, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not update the event", err)
		return
	}

	response.SuccessWithMsgNoObj(c, http.StatusAccepted, "Event deleted")
}
