package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hibiken/asynq"

	"oliveplay/utils/utils"
	"oliveplay/utils/utils/tasks"
)

// RegisterPingRoutes attaches routes for the Ping worker.
func RegisterPingRoutes(r chi.Router) {
	// Create a new Asynq client using your environment-based cluster configuration.
	asynqClient := asynq.NewClient(asynq.RedisClusterClientOpt{
		Addrs:    utils.RedisNodes,
		Password: utils.RedisPassword,
	})

	r.Post("/ping", func(w http.ResponseWriter, req *http.Request) {
		var p tasks.Ping

		log.Println("Request Body:", req.Body)

		// Decode the JSON request body into a Ping struct
		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Enqueue the ping task
		_, err := tasks.EnqueuePingTask(asynqClient, p)
		if err != nil {
			log.Printf("Failed to enqueue ping task: %v", err)
			http.Error(w, "Unable to enqueue the ping task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("Ping task enqueued"))
	})
}
