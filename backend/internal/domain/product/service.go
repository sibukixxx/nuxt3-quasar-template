package product

import "techvit/gin/internal/infrastructure/repository"

// ServiceProduct handles operations related to products.
type ServiceProduct struct {
	repo Repository
}

// NewProductService creates a new instance of ProductService.
func NewProductService(repo *repository.ProductRepository) *ServiceProduct {
	return &ServiceProduct{repo: repo}
}

// AddProduct encapsulates the logic to add a new product to the repository.
func (s *ServiceProduct) AddProduct(p *Product) error {
	return s.repo.Add(p)
}
