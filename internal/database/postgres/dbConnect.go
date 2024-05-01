package postgres

import (
	"database/sql"
	"log"
)

func DBConnect() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password= dbname=fin sslmode=disable")
	if err!= nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	err = db.Ping()
	if err!= nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	return db
}