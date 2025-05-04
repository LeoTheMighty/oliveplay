package types

import "time"

type GroupRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       *string `json:"image,omitempty"`
	Longitude   string  `json:"longitude"`
	Latitude    string  `json:"latitude"`
}

type GroupResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       *string   `json:"image,omitempty"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
}
