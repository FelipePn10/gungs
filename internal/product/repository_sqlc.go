package product

import (
	"context"

	"github.com/FelipePn10/gungs/internal/database/sqlc"
)

type repositorySQLC struct {
	q *sqlc.Queries
}

func (r *repositorySQLC) Create(ctx context.Context, p *Product) error {
	id, err := r.q.CreateProduct(ctx, sqlc.CreateProductParams{
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
	updated, err := r.q.UpdateProduct(ctx, sqlc.UpdateProductParams{
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

	*p = mapSQLCProductToDomain(updated)
	return nil
}

// func (r *repositorySQLC) Delete(ctx context.Context, p *Product) error {

// }

func NewRepository(q *sqlc.Queries) ProductRepository {
	return &repositorySQLC{q: q}
}
