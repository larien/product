package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/larien/product/product/controller"
	"github.com/larien/product/product/controller/mock"
	"github.com/larien/product/product/entity"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Parallel()
	t.Run("product repository failed", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		p := mock.NewProduct()
		p.On("List").Return(nil, errors.New(faker.Sentence())).Once()

		c := controller.New(p, nil)
		product, err := c.List(nil, "")

		p.AssertExpectations(t)
		is.EqualError(err, "an error occurred when listing the products")
		is.Nil(product)
	})
	t.Run("product repository was empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		p := mock.NewProduct()
		p.On("List").Return(nil, nil).Once()

		c := controller.New(p, nil)
		product, err := c.List(nil, "")

		p.AssertExpectations(t)
		is.Nil(err)
		is.Len(product, 0)
	})
	t.Run("product repository was successful", func(t *testing.T) {
		t.Parallel()
		t.Run("but discount repository failed", func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			p := mock.NewProduct()
			product := entity.Product{}
			is.Nil(faker.FakeData(&product))
			expectedProducts := []entity.Product{product}
			expectedProducts[0].Discount = entity.Discount{}
			p.On("List").Return(expectedProducts, nil).Once()

			d := mock.NewDiscount()
			expectedError := errors.New(faker.Sentence())
			userID := faker.UUIDDigit()
			d.On("Get", nil, product.ID, userID).Return(int64(0), expectedError).Once()
			c := controller.New(p, d)
			products, err := c.List(nil, userID)

			p.AssertExpectations(t)
			is.Nil(err)
			is.Len(products, len(expectedProducts))
			obtainedProduct := products[0]
			is.Equal(product.Title, obtainedProduct.Title)
			is.Equal(product.Description, obtainedProduct.Description)
			is.Equal(product.PriceInCents, obtainedProduct.PriceInCents)
			is.Zero(obtainedProduct.Discount.Percentage)
			is.Zero(obtainedProduct.Discount.ValueInCents)
		})
		t.Run("and discount repository was successful", func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			p := mock.NewProduct()
			product := entity.Product{}
			is.Nil(faker.FakeData(&product))
			expectedProducts := []entity.Product{product}
			p.On("List").Return(expectedProducts, nil).Once()

			d := mock.NewDiscount()
			userID := faker.UUIDDigit()
			expectedPercentage := int64(10)
			ctx := context.Background()
			d.On("Get", ctx, product.ID, userID).Return(expectedPercentage, nil).Once()
			c := controller.New(p, d)
			products, err := c.List(ctx, userID)

			p.AssertExpectations(t)
			is.Nil(err)
			is.Len(products, len(expectedProducts))
			obtainedProduct := products[0]
			is.Equal(product.Title, obtainedProduct.Title)
			is.Equal(product.Description, obtainedProduct.Description)
			is.Equal(product.PriceInCents, obtainedProduct.PriceInCents)
			is.Equal(expectedPercentage, obtainedProduct.Discount.Percentage)
			is.Equal((obtainedProduct.PriceInCents*expectedPercentage)/100, obtainedProduct.Discount.ValueInCents)
		})
	})
}
