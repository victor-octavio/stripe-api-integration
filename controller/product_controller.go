package controller

import (
	"back_end_mirumuh/model"
	"back_end_mirumuh/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{service: s}
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.service.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetAllProductsWithPrice(c *gin.Context) {
	products, err := pc.service.ListProductsWithPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}



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
