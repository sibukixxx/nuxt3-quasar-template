package product

// Repository provides an interface for product data storage.
type Repository interface {
	Add(product *Product) error
	GetByID(productID string) (*Product, error)
	Update(product *Product) error
	Delete(productID string) error
}
