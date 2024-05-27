package handler

import (
	"net/http"
	"techvit/gin/internal/application"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles HTTP requests for product operations.
type ProductHandler struct {
	productService *application.Se
}

// NewProductHandler creates a new ProductHandler.
func NewProductHandler(productService *application.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// RegisterRoutes registers the product-related routes to the Gin engine.
func (h *ProductHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/products/:id", h.GetProduct)
	router.POST("/products", h.CreateProduct)
}

// GetProduct handles GET requests for a product.
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct handles POST requests to create a new product.
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product application.ProductInput
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.productService.CreateProduct(product.ID, product.Name, product.Description, product.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
