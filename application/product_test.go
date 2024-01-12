package application_test

import (
	"fsgabriel/go-hexagonal-architecture/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.Disable
	product.Price = 10

	err := product.Enable()
	assert.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	assert.Error(t, err, "the price must be greater than zero to enable the product")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.Enable
	product.Price = 0

	err := product.Disable()
	assert.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	assert.Error(t, err, "the price must be zero in order to have the product disabled")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.Disable
	product.Price = 10

	_, err := product.IsValid()
	assert.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	assert.Error(t, err, "the status must be enabled or disabled")

	product.Status = application.Enable
	_, err = product.IsValid()
	assert.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	assert.Error(t, err, "the price must be greater or equal zero")
}
