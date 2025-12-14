package product

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repo     ProductRepository
	validate *validator.Validate
}

func (s *Service) Create(ctx context.Context, req CreateProductRequest) (*Product, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	product := &Product{
		UserID:      req.UserID,
		MetaData:    req.MetaData,
		Description: req.Description,
		Price:       req.Price,
		Tags:        req.Tags,
		Location:    req.Location,
		CreatedAt:   time.Now(),
	}
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func NewService(repo ProductRepository) *Service {
	return &Service{
		repo:     repo,
		validate: validator.New(),
	}
}
