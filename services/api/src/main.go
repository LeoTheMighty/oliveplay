package main

import (
	// 	"context"
	"log"
	"net/http"

	// 	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// 	"github.com/jackc/pgx/v5/pgxpool"

	"oliveplay/utils/utils"

	"oliveplay.co/api/src/routers"
)

func main() {
	// Initialize DB connection
	// dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatalf("Unable to connect to database: %v\n", err)
	// }
	// defer dbpool.Close()

	// // Example usage of sqlc generated code
	// queries := db.New(dbpool)

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	log.Println("CateUtils: ", utils.Utils("API"))

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(utils.Utils("API")))
		// w.Write([]byte("Welcome to the API"))
	})

	routers.RegisterPingRoutes(r)

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
