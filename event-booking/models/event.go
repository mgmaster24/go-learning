package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"go-learning.com/learning/event-booking/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      uuid.UUID
}

func (event *Event) Save(sqldb *sql.DB) error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)
	`

	res, err := db.PrepareAndExecVars(
		sqldb,
		query,
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	event.Id = id
	return err
}

func GetEvents(sqldb *sql.DB) ([]Event, error) {
	rows, err := db.GetRows(sqldb, "SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events = []Event{}
	for rows.Next() {
		var event Event
		err = rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEvent(sqldb *sql.DB, id int64) (Event, error) {
	var event Event
	err := db.GetValById(
		sqldb,
		"SELECT * FROM events WHERE id = ?",
		id,
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId)

	return event, err
}

func (event *Event) Update(sqldb *sql.DB, id int64) error {
	_, err := db.PrepareAndExecVars(
		sqldb,
		`
		UPDATE events
		SET name=?, description=?, location=?, dateTime=?, user_id=?
		WHERE id=?`,
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
		id)

	return err
}

func Delete(sqldb *sql.DB, id int64) error {
	_, err := db.PrepareAndExecVars(
		sqldb,
		`DELETE FROM events WHERE id=?`,
		id)
	return err
}
