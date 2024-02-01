package models

import (
	"database/sql"
	"time"

	"go-learning.com/learning/event-booking/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

type Registration struct {
	Id      int64
	EventId int64
	UserId  int64
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
	err := db.GetRowById(
		sqldb,
		"events",
		id,
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId)

	return event, err
}

func (event *Event) Update(sqldb *sql.DB) error {
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
		event.Id)

	return err
}

func (event *Event) Delete(sqldb *sql.DB) error {
	_, err := db.PrepareAndExecVars(
		sqldb,
		`DELETE FROM events WHERE id=?`,
		event.Id)
	return err
}

func (event *Event) Register(sqldb *sql.DB, userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?,?)"
	_, err := db.PrepareAndExecVars(sqldb, query, event.Id, userId)
	return err
}

func (event *Event) Unregiter(sqldb *sql.DB) error {
	_, err := db.PrepareAndExecVars(
		sqldb,
		`DELETE FROM registrations WHERE event_id=? AND user_id = ?`,
		event.Id,
		event.UserId)
	return err
}

func GetRegistrations(sqldb *sql.DB) ([]Registration, error) {
	rows, err := db.GetRows(sqldb, "SELECT * FROM registrations")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registrations = []Registration{}
	for rows.Next() {
		var registration Registration
		err = rows.Scan(
			&registration.Id,
			&registration.EventId,
			&registration.UserId)

		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}

	return registrations, nil
}
