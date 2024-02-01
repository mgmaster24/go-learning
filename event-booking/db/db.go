package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could NOT connect to database!")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createEventsTable(db)
	createUsersTable(db)
	createRegistrationsTable(db)

	return db
}

func createEventsTable(db *sql.DB) {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, err := db.Exec(createEventsTable)
	if err != nil {
		panic("Could NOT create the events table!")
	}
}

func createUsersTable(db *sql.DB) {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		panic("Could NOT create the users table!")
	}
}

func createRegistrationsTable(db *sql.DB) {
	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err := db.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could NOT create the registrations table!")
	}
}

func PrepareAndExecVars(db *sql.DB, query string, vars ...any) (sql.Result, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return stmt.Exec(vars...)
}

func GetRows(db *sql.DB, query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func GetRowById(db *sql.DB, table string, id int64, vals ...any) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table)
	row := db.QueryRow(query, id)
	return row.Scan(vals...)
}

func GetRowByVal(db *sql.DB, selectVal string, table string, key string, val any, vals ...any) error {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", selectVal, table, key)
	row := db.QueryRow(query, val)
	return row.Scan(vals...)
}
