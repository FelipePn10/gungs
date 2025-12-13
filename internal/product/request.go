package product

import (
	"encoding/json"
)

type CreateProductRequest struct {
	UserID      int64           `json:"user_id" binding:"required"`
	MetaData    json.RawMessage `json:"metadata" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Price       float64         `json:"price" binding:"required"`
	Tags        []string        `json:"tags"`
	Location    string          `json:"location" binding:"required"`
}

type UpdateProductRequest struct {
	MetaData    json.RawMessage `json:"metadata" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Price       float64         `json:"price" binding:"required"`
	Tags        []string        `json:"tags"`
	Location    string          `json:"location" binding:"required"`
}
