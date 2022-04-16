package cli

import (
	"fmt"
	"github.com/loxt/fullcycle2.0-ports-and-adapters/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""
	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s and price %f has been created", product.GetID(), product.GetName(), product.GetPrice())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		result = fmt.Sprintf("Product %s has been enabled", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s and price %f has been found - Status: %s", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
