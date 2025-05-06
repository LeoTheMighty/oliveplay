package utils

import (
	"os"
	"strings"
	"time"
)

// Environment variables
var (
	RedisNodes    = getEnvListOrDefault("REDIS_NODES", []string{"redis-node-1:6380", "redis-node-2:6381", "redis-node-3:6382"})
	RedisUsername = getEnvOrDefault("REDIS_USERNAME", "default")
	RedisPassword = getEnvOrDefault("REDIS_PASSWORD", "password")

	// API configuration
	APIPort    = getEnvOrDefault("API_PORT", "8080")
	APITimeout = getDurationEnvOrDefault("API_TIMEOUT", 30*time.Second)

	DatabaseUrl = getEnvOrDefault("DATABASE_URL", "NOT_SET")

	S3Bucket   = getEnvOrDefault("S3_BUCKET", "oliveplay")
	AWSRegion  = getEnvOrDefault("AWS_REGION", "us-west-2")
	AWSSession = getEnvOrDefault("AWS_SESSION", "default")
)

// Hardcoded constants
const (
	// Status codes
	StatusSuccess = "success"
	StatusError   = "error"

	// Cache keys
	CacheKeyPrefix = "oliveplay:"
	UserKeyPrefix  = "user:"

	// Default values
	DefaultPageSize   = 10
	MaxPageSize       = 100
	DefaultTimeoutSec = 30
)

// Helper function to get environment variable with default fallback
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Helper function to get list of strings from environment variable with default fallback
func getEnvListOrDefault(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}

// Helper function to get duration from environment variable with default fallback
func getDurationEnvOrDefault(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
