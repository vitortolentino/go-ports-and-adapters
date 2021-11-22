package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
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

func Test_IsValid(t *testing.T) {
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "Hello",
		Price: 10,
		Status: application.DISABLED,
	}	

	_, err := product.IsValid()

	require.Nil(t, err)
}

func Test_IsValid_ReturnErrWithWrongStatus(t *testing.T) {
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "Hello",
		Price: 10,
		Status: "INVALID",
	}	

	_, err := product.IsValid()

	require.Equal(t, "the status must be enabled or disabled", err.Error())
}

func Test_IsValid_ReturnErrWithNegativeNumber(t *testing.T) {
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "Hello",
		Price: -10,
		Status: application.ENABLED,
	}	

	_, err := product.IsValid()
	require.Equal(t, "the price must be greather or equal zero", err.Error())
}