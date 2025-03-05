package model

type CheckoutSessionInput struct {
	CustomerEmail string `json:"customer_email" binding:"required"`
	PriceID    string `json:"price_id" binding:"required"`
	Quantity   int64  `json:"quantity" binding:"required,min=1"`
	SuccessURL string `json:"success_url" binding:"required,url"`
	CancelURL  string `json:"cancel_url" binding:"required,url"`
}

type CheckoutSessionOutput struct {
	SessionID string `json:"session_id"`
	URL       string `json:"url"`
}
