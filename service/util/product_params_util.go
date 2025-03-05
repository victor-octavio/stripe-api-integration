package util

import (
	"github.com/stripe/stripe-go"
)

func NewProductParams(name, description string, active bool, provider string) (*stripe.ProductParams, error) {
	stripe.Key = provider
	params := &stripe.ProductParams{
		Name:        stripe.String(name),
		Description: stripe.String(description),
		Active:      stripe.Bool(active),
	}
	return params, nil
}

func ListProductParams(provider string) (*stripe.ProductListParams, error) {
	stripe.Key = provider
	params := &stripe.ProductListParams{}
	return params, nil
}

func ListProductParamsWithPrice(provider string) (*stripe.ProductListParams, error){
		stripe.Key = provider
		params := &stripe.ProductListParams{}
		params.Limit = stripe.Int64(10) // Limite de produtos por página
	
		return params, nil 
		
}
