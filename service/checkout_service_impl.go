package service

import (
	"errors"
	"stripe-api-integration/model"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

type checkoutServiceImpl struct {
	apiKey string
}

func NewCheckoutService(apiKey string) CheckoutService {
	return &checkoutServiceImpl{apiKey: apiKey}
}

func (s *checkoutServiceImpl) CreateCheckoutSession(input model.CheckoutSessionInput) (*model.CheckoutSessionOutput, error) {
	stripe.Key = s.apiKey

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}), // TODO: Verificar outras formas de pagamento
		CustomerEmail: stripe.String(input.CustomerEmail),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(input.PriceID),
				Quantity: stripe.Int64(input.Quantity),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(input.SuccessURL),
		CancelURL:  stripe.String(input.CancelURL),
	}

	sess, err := session.New(params)
	if err != nil {
		return nil, errors.New("failed to create checkout session: " + err.Error())
	}

	return &model.CheckoutSessionOutput{
		SessionID: sess.ID,
		URL:       sess.URL,
	}, nil
}
