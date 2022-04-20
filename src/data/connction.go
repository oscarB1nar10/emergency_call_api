
package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetConnection() *sql.DB {
	// Connection String to the local db
	connStr := "postgres://postgres:B1nar10.01@localhost/USER_TRACKING?sslmode=disable"
	// Return a connection to the db or an error
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}