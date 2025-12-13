package product

import (
	"encoding/json"
	"time"
)

type Product struct {
	ID          int64           `json:"id"`
	UserID      int64           `json:"user_id"`
	MetaData    json.RawMessage `json:"metadata"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Tags        []string        `json:"tags"`
	Location    string          `json:"location"`
	CreatedAt   time.Time       `json:"created_at"`
}
