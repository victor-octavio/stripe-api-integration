package controller

import (
	"net/http"
	"stripe-api-integration/model"
	"stripe-api-integration/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{service: s}
}

// GetAllProducts godoc
// @Summary Lista todos os produtos
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.service.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetAllProductsWithPrice godoc
// @Summary Lista produtos com seus preços
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {object} map[string]string
// @Router /products/with-prices [get]
func (pc *ProductController) GetAllProductsWithPrice(c *gin.Context) {
	products, err := pc.service.ListProductsWithPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary Cria um produto (e preço inicial)
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.CreateProductInput true "Dados do produto"
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input model.CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := pc.service.CreateProduct(input.Name, input.Description, input.Active, input.Price, input.Currency)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product."})
		return
	}

	c.JSON(http.StatusOK, product)
}
