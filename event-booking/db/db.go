package db

import (
	"database/sql"

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
			user_id INTEGER
		)
	`

	_, err := db.Exec(createEventsTable)
	if err != nil {
		panic("Could NOT create the events table!")
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

func GetValById(db *sql.DB, query string, id int64, vals ...any) error {
	row := db.QueryRow("SELECT * FROM events WHERE id = ?", id)
	return row.Scan(vals...)
}
