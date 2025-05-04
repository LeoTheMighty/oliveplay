package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/leothemighty/oliveplay/utils/utils"
)

type Database struct {
	pool    *pgxpool.Pool
	Queries *Queries
}

func OpenDatabase() *Database {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, utils.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		pool:    pool,
		Queries: New(pool),
	}
}

func CloseDatabase(database *Database) {
	database.pool.Close()
}
