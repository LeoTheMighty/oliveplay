package main

import (
	"fmt"
	"log"

	// "os"

	"github.com/golang-migrate/migrate/v4"
	// For pgx driver

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// We'll import cate_utils so we can reference the migrations dir if needed
	"github.com/leothemighty/oliveplay/utils/utils"
)

func main() {
	dbURL := utils.DatabaseUrl
	fmt.Println("Database URL: ", dbURL)
	if dbURL == "" {
		log.Fatalf("DATABASE_URL not set")
	}

	// Example local usage: use the "file://" prefix to point to the actual folder
	// containing your migrations from cate_utils if it's bundled together or copied
	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("failed to create migrator: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("Migrations ran successfully!")
}
