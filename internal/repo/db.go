package repo

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qustavo/dotsql"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

func CreateDb(dsn string, refresh bool) (*sql.DB, error) {

	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	db = conn

	if refresh {
		// Loads schema from file
		dot, _ := dotsql.LoadFromFile("./migrations/schema.sql")
		// Run queries
		_, err = dot.Exec(db, "drop-user-table")
		_, err = dot.Exec(db, "drop-post-table")
		_, err = dot.Exec(db, "create-user-table")
		_, err = dot.Exec(db, "create-post-table")

		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
