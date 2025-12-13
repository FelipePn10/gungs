package product

import (
	"encoding/json"
	"time"
)

type CreateProductResponse struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateProductResponse struct {
	MetaData    json.RawMessage `json:"metadata" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Price       float64         `json:"price" binding:"required"`
	Tags        []string        `json:"tags"`
	Location    string          `json:"location" binding:"required"`
}

func ToProductCreateResponse(p *Product) CreateProductResponse {
	return CreateProductResponse{
		ID:          p.ID,
		Description: p.Description,
		Price:       p.Price,
		Location:    p.Location,
		CreatedAt:   p.CreatedAt,
	}
}

func ToProductUpdateResponse(p *Product) UpdateProductResponse {
	return UpdateProductResponse{
		MetaData:    p.MetaData,
		Description: p.Description,
		Price:       p.Price,
		Tags:        p.Tags,
		Location:    p.Location,
	}
}
