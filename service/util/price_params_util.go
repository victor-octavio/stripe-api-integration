package util

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/price"
)

func NewPriceParams(product string, unitAmount int64, currency string, provider string) (*stripe.Price, error) {
	stripe.Key = provider
	params := &stripe.PriceParams{
		Product:    stripe.String(product),
		UnitAmount: stripe.Int64(unitAmount),
		Currency:   stripe.String(currency),
	}
	return price.New(params)
}

func GetProducWithPriceParams(product string, provider string) (*stripe.PriceListParams, error) {
	stripe.Key = provider
	params := &stripe.PriceListParams{
		Product:    stripe.String(product),
	}
	return params, nil
}
