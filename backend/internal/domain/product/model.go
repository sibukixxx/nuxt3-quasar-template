package product

import "fmt"

// Product represents the core of the product domain model.
type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

// NewProduct creates and validates a new product entity.
func NewProduct(id, name, description string, price float64) (*Product, error) {
	if price <= 0 {
		return nil, fmt.Errorf("price must be positive")
	}
	return &Product{id, name, description, price}, nil
}
