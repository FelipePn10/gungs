package product

import (
	"context"

	"github.com/FelipePn10/gungs/internal/database"
	"github.com/FelipePn10/gungs/internal/database/sqlc"
)

type repositorySQLC struct {
	q *database.DB
}

func (r *repositorySQLC) Create(ctx context.Context, p *Product) error {
	id, err := r.q.Queries().CreateProduct(ctx, sqlc.CreateProductParams{
		UsersID:     p.UserID,
		Metadata:    p.MetaData,
		Description: p.Description,
		Price:       p.Price,
		Tags:        p.Tags,
		Location:    p.Location,
	})
	if err != nil {
		return err
	}
	p.ID = id.ID
	return nil
}

func (r *repositorySQLC) Update(ctx context.Context, p *Product) error {
	updatedProduct, err := r.q.Queries().UpdateProduct(ctx, sqlc.UpdateProductParams{
		ID:          p.ID,
		Metadata:    p.MetaData,
		Location:    p.Location,
		Description: p.Description,
		Price:       p.Price,
		Tags:        p.Tags,
	})
	if err != nil {
		return err
	}
	*p = mapSQLCProductToDomain(updatedProduct)
	return nil
}

// func (r *repositorySQLC) Delete(ctx context.Context, p *Product) error {

// }

func NewRepository(q *database.DB) ProductRepository {
	return &repositorySQLC{q: q}
}
