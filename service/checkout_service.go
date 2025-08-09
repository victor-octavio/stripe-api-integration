package service

import "stripe-api-integration/model"

type CheckoutService interface {
	CreateCheckoutSession(input model.CheckoutSessionInput) (*model.CheckoutSessionOutput, error)
}
