package cmd

import (
	"log"
	"techvit/gin/internal/application"
	"techvit/gin/internal/domain/product"
	"techvit/gin/internal/infrastructure/repository"
)

func main() {
	productRepo := repository.NewProductRepository()
	productService := product.NewProductService(productRepo)
	appService := application.NewProductService(productService)

	// Example usage
	err := appService.CreateProduct("1", "Gadget", "An awesome gadget", 99.99)
	if err != nil {
		log.Fatal(err)
	}
}
