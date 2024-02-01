package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-learning.com/learning/event-booking/models"
	"go-learning.com/learning/event-booking/response"
	"go-learning.com/learning/event-booking/utils"
)

func (sqlDB *SqlDB) createUser(c *gin.Context) {
	var user models.User
	if !shouldBindJSON(c, &user) {
		return
	}

	err := user.Save(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not save User to DB.", err)
		return
	}

	response.SuccesWithMsg(c, http.StatusCreated, "User created", "user", user.Email)
}

func (sqlDB *SqlDB) login(c *gin.Context) {
	var user models.User
	if !shouldBindJSON(c, &user) {
		return
	}

	err := user.ValidateUser(sqlDB.DB)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Unable to login", err)
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.Id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Token generation failed", err)
	}

	response.SuccesWithMsg(c, http.StatusOK, "Successfully logged in!", "token", token)
}
