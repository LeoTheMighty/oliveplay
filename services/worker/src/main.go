package main

import (
	"context"
	"log"

	"github.com/leothemighty/oliveplay/utils/utils/tasks"
)

func main() {
	// Initialize Redis cluster
	tasks.StartTaskServer(context.Background())
	log.Println("Task service started. Listening for tasks...")
}
