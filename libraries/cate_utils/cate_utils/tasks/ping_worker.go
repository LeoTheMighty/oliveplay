package tasks

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

// TypePing identifies our ping task type.
const TypePing = "ping"
type Ping struct {
	Message string `json:"message"`
} 

// NewPingTask creates an Asynq task for a Ping payload.
func NewPingTask(p Ping) (*asynq.Task, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypePing, data), nil
}

// EnqueuePingTask is a helper function to enqueue a ping task with a given Asynq client.
func EnqueuePingTask(client *asynq.Client, p Ping) (*asynq.TaskInfo, error) {
	task, err := NewPingTask(p)
	if err != nil {
		return nil, err
	}
	return client.Enqueue(task)
}

// HandlePingTask processes the Ping task and logs its contents.
func HandlePingTask(ctx context.Context, t *asynq.Task) error {
	var pingData Ping
	if err := json.Unmarshal(t.Payload(), &pingData); err != nil {
		log.Printf("Failed to unmarshal Ping payload: %v", err)
		return err
	}

	// Here is where you would implement logic to handle "ping".
	// For now, we just log it:
	log.Printf("Received Ping: %+v", pingData)
	return nil
} 