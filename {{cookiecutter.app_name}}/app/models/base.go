package models

import "time"

// Base is a model is used common attributes of models
type Base struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"updated_at"`
	UpdatedAt time.Time `json:"created_at"`
}

// Pagination is common models to contain pagination data
type Pagination struct {
	Page  *int64 `json:"page"`
	Limit *int64 `json:"limit"`
}

// Order is common models to contain sorting data
type Order struct {
	OrderBy   *string
	SortOrder *string
}
