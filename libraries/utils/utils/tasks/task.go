package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"

	"oliveplay/utils/utils"
)

// TaskStatus represents the current state of a task
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusComplete  TaskStatus = "complete"
	TaskStatusFailed    TaskStatus = "failed"
)

// Task represents an asynchronous task that can be queued and executed
type Task struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Status    TaskStatus             `json:"status"`
	Payload   map[string]interface{} `json:"payload"`
	Error     string                 `json:"error,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

// TaskHandler is a function that processes a task
type TaskHandler func(context.Context, *Task) error

// TaskQueue manages task operations with Redis
type TaskQueue struct {
	client     *redis.Client
	handlers   map[string]TaskHandler
	queueName  string
}

func generateTaskID() string {
	return uuid.New().String()
}

// NewTaskQueue creates a new TaskQueue instance
func NewTaskQueue(client *redis.Client, queueName string) *TaskQueue {
	return &TaskQueue{
		client:    client,
		handlers:  make(map[string]TaskHandler),
		queueName: queueName,
	}
}

// RegisterHandler registers a handler function for a specific task type
func (tq *TaskQueue) RegisterHandler(taskType string, handler TaskHandler) {
	tq.handlers[taskType] = handler
}

// EnqueueTask adds a new task to the queue
func (tq *TaskQueue) EnqueueTask(ctx context.Context, taskType string, payload map[string]interface{}) (*Task, error) {
	task := &Task{
		ID:        generateTaskID(),
		Type:      taskType,
		Status:    TaskStatusPending,
		Payload:   payload,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	taskJSON, err := json.Marshal(task)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal task: %w", err)
	}

	// Add task to Redis list
	err = tq.client.LPush(ctx, tq.queueName, taskJSON).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to enqueue task: %w", err)
	}

	return task, nil
}

// ProcessTasks starts processing tasks from the queue
func (tq *TaskQueue) ProcessTasks(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Pop task from queue (blocking)
			result, err := tq.client.BRPop(ctx, 0, tq.queueName).Result()
			if err != nil {
				if err == redis.Nil {
					continue
				}
				return fmt.Errorf("failed to dequeue task: %w", err)
			}

			var task Task
			if err := json.Unmarshal([]byte(result[1]), &task); err != nil {
				continue
			}

			// Update task status
			task.Status = TaskStatusRunning
			task.UpdatedAt = time.Now()

			// Execute handler
			handler, exists := tq.handlers[task.Type]
			if !exists {
				task.Status = TaskStatusFailed
				task.Error = "no handler registered for task type"
				continue
			}

			if err := handler(ctx, &task); err != nil {
				task.Status = TaskStatusFailed
				task.Error = err.Error()
			} else {
				task.Status = TaskStatusComplete
			}

			task.UpdatedAt = time.Now()
			// Store task result in Redis (you might want to implement this)
		}
	}
}

// NewAsynqServer sets up the Asynq server for a Redis cluster.
func NewAsynqServer() (*asynq.Server, error) {
	// Use the environment-based cluster configuration from cate_utils.
	log.Println("RedisNodes:", utils.RedisNodes)
	redisOpt := asynq.RedisClusterClientOpt{
		Addrs:        utils.RedisNodes,
		Password:     utils.RedisPassword, // Uncomment if password is required
		Username:     utils.RedisUsername, // Add if username is required
		DialTimeout:  5 * time.Second,          // Add reasonable timeouts
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		// PoolSize:     10,      	
	}

	// Create a new Asynq server with some default concurrency settings.
	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10, // adjust as needed
		RetryDelayFunc: func(n int, err error, t *asynq.Task) time.Duration {
			return time.Second * 5 // Wait 5 seconds between retries
		},
		// Logger: log.New(os.Stdout, "asynq: ", log.Ldate|log.Ltime|log.Lmicroseconds),
	})

	return server, nil
}

// StartTaskServer launches the Asynq server and serves any registered tasks.
func StartTaskServer(ctx context.Context) error {
	server, err := NewAsynqServer()
	if err != nil {
		return err
	}

	// A ServeMux allows us to register multiple task handlers from different files.
	mux := asynq.NewServeMux()

	// ----------------------------------------------------------------
	// Register handlers from your worker files below.
	// Example: mux.Handle(TypePing, asynq.HandlerFunc(HandlePingTask))
	// ----------------------------------------------------------------

	// This will register the ping worker's task type and handler.
	mux.Handle(TypePing, asynq.HandlerFunc(HandlePingTask))

	// Run the server in a blocking manner until context is done.
	errCh := make(chan error, 1)
	go func() {
		if serr := server.Run(mux); serr != nil {
			errCh <- serr
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Shutting down the Asynq server due to context cancellation...")
		server.Shutdown()
		return ctx.Err()
	case runErr := <-errCh:
		log.Printf("Asynq server returned an error: %v\n", runErr)
		return runErr
	}
}
