package models

import (
	"database/sql"
	"errors"

	"go-learning.com/learning/event-booking/db"
	"go-learning.com/learning/event-booking/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save(sqldb *sql.DB) error {
	query := "INSERT INTO users(email, password)	VALUES (?,?)"
	hashedPW, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := db.PrepareAndExecVars(
		sqldb,
		query,
		u.Email,
		hashedPW)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.Id = id
	return err
}

func (u *User) ValidateUser(sqldb *sql.DB) error {
	var retrievedPW string
	var id int64
	err := db.GetRowByVal(
		sqldb,
		"id, password",
		"users",
		"email",
		u.Email,
		&id,
		&retrievedPW)
	if err != nil {
		return err
	}

	u.Id = id
	if utils.CheckPWHash(u.Password, retrievedPW) {
		return nil
	}

	return errors.New("Password validation failed.")
}

func GetUser(sqldb *sql.DB, id int64) (User, error) {
	var user User
	err := db.GetRowById(
		sqldb,
		"users",
		id,
		&user.Id,
		&user.Email,
		&user.Password)

	return user, err
}
