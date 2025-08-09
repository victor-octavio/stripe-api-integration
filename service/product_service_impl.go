package service

import (
	"github.com/stripe/stripe-go/product"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/price"
	"stripe-api-integration/service/util"
)

type productServiceImpl struct {
	apiKey string
}

func NewProductService(apiKey string) ProductService {
	stripe.Key = apiKey
	return &productServiceImpl{apiKey: apiKey}
}
func (s *productServiceImpl) CreateProduct(name string, description string, active bool, price int64, currency string) (map[string]interface{}, error) {
	params, err := util.NewProductParams(name, description, active, s.apiKey)
	if err != nil {
		return nil, err
	}

	p, err := product.New(params)
	if err != nil {
		return nil, err
	}

	priceParams, err := util.NewPriceParams(p.ID, price, currency, s.apiKey)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":          p.ID,
		"name":        p.Name,
		"description": p.Description,
		"active":      p.Active,
		"price_id":    priceParams.ID,
		"price":       priceParams.UnitAmount,
		"currency":    priceParams.Currency,
	}, nil
}

func (s *productServiceImpl) ListProducts() ([]map[string]interface{}, error) {
	params, err := util.ListProductParams(s.apiKey)

	if err != nil {
		return nil, err
	}

	result := product.List(params)
	var products []map[string]interface{}

	for result.Next() {
		p := result.Product()
		products = append(products, map[string]interface{}{
			"id":          p.ID,
			"name":        p.Name,
			"description": p.Description,
			"active":      p.Active,
		})
	}

	if err := result.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productServiceImpl) ListProductsWithPrices() ([]map[string]interface{}, error) {
	params, err := util.ListProductParamsWithPrice(s.apiKey)

	if err != nil {
		return nil, err
	}

	productIterator := product.List(params)

	var productsWithPrices []map[string]interface{}

	for productIterator.Next() {
		p := productIterator.Product()

		priceListParams, err := util.GetProducWithPriceParams(*stripe.String(p.ID), s.apiKey)

		if err != nil {
			return nil, err
		}

		priceIterator := price.List(priceListParams)

		var prices []map[string]interface{}

		for priceIterator.Next() {
			pr := priceIterator.Price()
			prices = append(prices, map[string]interface{}{
				"id":       pr.ID,
				"amount":   pr.UnitAmount,
				"currency": pr.Currency,
			})
		}

		productsWithPrices = append(productsWithPrices, map[string]interface{}{
			"id":          p.ID,
			"name":        p.Name,
			"description": p.Description,
			"active":      p.Active,
			"prices":      prices,
		})
	}

	if err := productIterator.Err(); err != nil {
		return nil, err
	}

	return productsWithPrices, nil
}
