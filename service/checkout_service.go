package service

import "back_end_mirumuh/model"

type CheckoutService interface {
	CreateCheckoutSession(input model.CheckoutSessionInput) (*model.CheckoutSessionOutput, error)
}
