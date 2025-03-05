package router

import (
	"back_end_mirumuh/factory"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	f := factory.NewFactory()

	// Products
	r.GET("/products", f.ProductController.GetAllProducts)
	r.POST("/product", f.ProductController.CreateProduct)
	r.GET("/available-products", f.ProductController.GetAllProductsWithPrice)

	//Checkout
	r.POST("/checkout", f.CheckoutController.CreateCheckoutSession)
}
