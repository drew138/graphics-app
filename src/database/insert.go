package database

import (
	"log"
)

func insertImage(userId int, provider string) {
	db, err := createConnection()
	if err != nil {
		log.Fatalf("Error while connecting to daabase: %v", err)
	}
	db.Exec("")

}
