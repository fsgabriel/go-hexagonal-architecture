package application_test

import (
	"fsgabriel/go-hexagonal-architecture/application"
	mock_application "fsgabriel/go-hexagonal-architecture/application/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Repository: persistence,
	}

	result, err := service.Get("3123")
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Repository: persistence,
	}

	result, err := service.Create("product1", 12.5)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Repository: persistence,
	}

	result, err := service.Enable(product)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil).AnyTimes()
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Repository: persistence,
	}

	result, err := service.Disable(product)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}
