package main

import (
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
)

func main() {
	m, err := migrate.New(
		"file:///migrations",
		"postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")+"?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("could not create migration: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("could not run up migrations: %v", err)
	}

	log.Println("migrations ran successfully")
}
