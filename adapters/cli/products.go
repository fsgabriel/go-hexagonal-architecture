package cli

import (
	"fmt"

	"github.com/fsgabriel/go-hexagonal-architecture/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result string
	switch action {
	case "create":
		p, err := service.Create(productName, price)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("product id %s with name %s has been created with price %f and status %s",
			p.GetID(), p.GetName(), p.GetPrice(), p.GetStatus())
	case "enable":
		p, err := service.Get(productId)
		if err != nil {
			return "", err
		}

		res, err := service.Enable(p)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("product %s has been enable.", res.GetName())
	case "disable":
		p, err := service.Get(productId)
		if err != nil {
			return "", err
		}

		res, err := service.Disable(p)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("product %s has been disable.", res.GetName())
	default:
		p, err := service.Get(productId)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("product id %s, product name %s, price %f, status %s",
			p.GetID(), p.GetName(), p.GetPrice(), p.GetStatus(),
		)
	}

	return result, nil
}
