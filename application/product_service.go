package application

import "fmt"

type ProductService struct {
	Repository ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %w", err)
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}

	newProduct, err := s.Repository.Save(product)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}

	result, err := s.Repository.Save(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return nil, err
	}

	result, err := s.Repository.Save(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}
