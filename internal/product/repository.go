package product

import (
	"context"
)

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	//Delete(ctx context.Context, product *Product) error
}
