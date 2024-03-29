package main

import (
	"database/sql"
	"graphics-app-backend/src/api"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// Run migrations / panic if can not migrate
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://src/migrations",
		"postgres",
		driver)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		panic(err)
	}

	// Set up server
	r := gin.Default()
	r.Use(gin.Recovery())

	api.RegisterRoutes(r)
	r.Run()
}
