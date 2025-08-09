package factory

import (
	"log"
	"os"
	"stripe-api-integration/controller"
	"stripe-api-integration/service"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v81"
)

type Factory struct {
	ProductController  *controller.ProductController
	CheckoutController *controller.CheckoutController
}

func NewFactory() *Factory {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env:", err)
	}

	apiKey := os.Getenv("STRIPE_KEY")
	stripe.Key = apiKey

	if stripe.Key == "" {
		log.Fatal("Stripe API key is missing.")
	}

	// Product
	productService := service.NewProductService(apiKey)
	productController := controller.NewProductController(productService)

	// Checkout
	checkoutService := service.NewCheckoutService(apiKey)
	checkoutController := controller.NewCheckoutController(checkoutService)

	return &Factory{
		ProductController:  productController,
		CheckoutController: checkoutController,
	}
}
