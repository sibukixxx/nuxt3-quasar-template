package application

import (
	"techvit/gin/internal/domain/product"
)

// ServiceProduct provides application services related to products.
type ServiceProduct struct {
	serviceProduct *product.ServiceProduct
}

// NewProductService creates a new ProductService.
func NewProductService(serviceProduct *product.ServiceProduct) *ServiceProduct {
	return &ServiceProduct{serviceProduct: serviceProduct}
}

// CreateProduct handles the creation of a product.
func (s *ServiceProduct) CreateProduct(id, name, description string, price float64) error {
	prod, err := product.NewProduct(id, name, description, price)
	if err != nil {
		return err
	}
	return s.serviceProduct.AddProduct(prod)
}
