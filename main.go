package main

import (
	"log"
	"stripe-api-integration/docs"
	"stripe-api-integration/router"

	"github.com/gin-gonic/gin"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Stripe API Integration
// @version 1.0
// @description API para produtos, preços e checkout com Stripe.
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Stripe API Integration"
	docs.SwaggerInfo.Description = "API para produtos, preços e checkout com Stripe."
	docs.SwaggerInfo.BasePath = "/"

	router.InitRoutes(r.Group(""))

	if err := r.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
