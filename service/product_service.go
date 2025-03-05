package service

type ProductService interface {
	ListProducts() ([]map[string]interface{}, error)
	CreateProduct(name string, description string, active bool, price int64, currency string) (map[string]interface{}, error)
	ListProductsWithPrices() ([]map[string]interface{}, error)
}
