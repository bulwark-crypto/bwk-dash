package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NewSQL will setup a new database connection.
func NewSQL(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", dsn)
	return
}
