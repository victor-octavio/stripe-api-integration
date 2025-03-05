package model

type Product struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       int64  `json:"price" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
}

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Active      bool   `json:"active" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
}
