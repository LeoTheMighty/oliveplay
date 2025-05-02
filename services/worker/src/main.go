package main

import (
	"context"
	"log"

	"cate/cate_utils/cate_utils/tasks"
)

func main() {
	// Initialize Redis cluster
	tasks.StartTaskServer(context.Background())
	log.Println("Task service started. Listening for tasks...")
}
