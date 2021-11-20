package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vitortolentino/go-ports-and-adapters/application"
)

func TestProduct_Enable(t * testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)
}

func TestProduct_EnableErrorWithPriceZero(t * testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t * testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)
}

func TestProduct_DisableErrorWithPriceGreaterThanZero(t * testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()

	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}