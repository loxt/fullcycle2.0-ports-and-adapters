package application_test

import (
	"github.com/loxt/fullcycle2.0-ports-and-adapters/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	if err != nil {
		require.Nil(t, err)
	}

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
