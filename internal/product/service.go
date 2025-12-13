package product

type Service struct {
	repo ProductRepository
}

// func (s *Service) Create(ctx context.Context, req CreateProductRequest) (*Product, error) {

// }

func NewService(repo ProductRepository) *Service {
	return &Service{repo: repo}
}
