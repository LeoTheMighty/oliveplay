package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	// "github.com/hibiken/asynq"

	"github.com/leothemighty/oliveplay/utils/services/group"
)

// RegisterPingRoutes attaches routes for the Ping worker.
func RegisterGroupRoutes(r chi.Router) {
	// // create a new asynq client using your environment-based cluster configuration.
	// asynqclient := asynq.newclient(asynq.redisclusterclientopt{
	// 	addrs:    utils.redisnodes,
	// 	password: utils.redispassword,
	// })

	r.Post("/group", func(w http.ResponseWriter, req *http.Request) {
		response, err := group.CreateGroup(context.Background(), req)
		if err != nil {
			fmt.Println("Error creating group:", err)
			http.Error(w, "Unable to create group", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	r.Get("/groups", func(w http.ResponseWriter, req *http.Request) {
		response, err := group.GetCloseGroups(context.Background(), req)
		if err != nil {
			fmt.Println("Error getting close groups:", err)
			http.Error(w, "Unable to get close groups", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})


	r.Get("/group/{id}", func(w http.ResponseWriter, req *http.Request) {
		// response, err := group.GetGroup(context.Background(), req)
		// Enqueue the ping task
		// if err != nil {
		// 	log.Printf("Failed to enqueue ping task: %v", err)
		// 	http.Error(w, "Unable to enqueue the ping task", http.StatusInternalServerError)
		// 	return
		// }

		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("Ping task enqueued"))
	})
}
