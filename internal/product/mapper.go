package product

import (
	"time"

	"github.com/FelipePn10/gungs/internal/database/sqlc"
)

func mapSQLCProductToDomain(p sqlc.Product) Product {
	return Product{
		ID:          p.ID,
		UserID:      p.UsersID,
		MetaData:    p.Metadata,
		Description: p.Description,
		Price:       p.Price,
		Tags:        p.Tags,
		Location:    p.Location,
		CreatedAt:   time.Time{},
	}
}
