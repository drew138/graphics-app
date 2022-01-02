package database

import (
	"database/sql"
	"log"
	"os"
)

func createConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	log.Println("Successfully connected to database")
	return db, nil
}
