package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"stripe-api-integration/factory"
)

func InitRoutes(r *gin.RouterGroup) {
	f := factory.NewFactory()

	// Products
	r.GET("/products", f.ProductController.GetAllProducts)
	r.POST("/product", f.ProductController.CreateProduct)
	r.GET("/available-products", f.ProductController.GetAllProductsWithPrice)

	//Checkout
	r.POST("/checkout", f.CheckoutController.CreateCheckoutSession)

	//Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
