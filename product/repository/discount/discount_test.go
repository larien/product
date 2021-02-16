package discount_test

import (
	"context"
	"errors"
	"testing"

	faker "github.com/bxcodec/faker/v3"
	"github.com/larien/product/product/repository/discount"
	serverMock "github.com/larien/product/product/repository/discount/mock"
	protobuf "github.com/larien/product/protos"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()
	t.Run("user ID is empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		d := discount.New(nil)
		percentage, err := d.Get(nil, "", "")

		is.Equal(percentage, int64(0))
		is.Nil(err)
	})
	t.Run("failed to obtain discount", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		s := serverMock.New()

		req := &protobuf.DiscountRequest{
			ProductID: faker.UUIDDigit(),
			UserID:    faker.UUIDDigit(),
			RequestID: "",
		}
		expectedError := errors.New(faker.Sentence())
		s.On("Discount", context.Background(), req).Return(nil, expectedError).Once()

		d := discount.New(s)
		percentage, err := d.Get(context.Background(), req.ProductID, req.UserID)

		s.AssertExpectations(t)
		is.Equal(percentage, int64(0))
		is.Contains(err.Error(), expectedError.Error())
	})
	t.Run("obtained discount with success", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		s := serverMock.New()

		req := &protobuf.DiscountRequest{
			ProductID: faker.UUIDDigit(),
			UserID:    faker.UUIDDigit(),
			RequestID: "",
		}
		res := &protobuf.DiscountResponse{
			Percentage: 10,
		}
		s.On("Discount", context.Background(), req).Return(res, nil).Once()

		d := discount.New(s)
		percentage, err := d.Get(context.Background(), req.ProductID, req.UserID)

		s.AssertExpectations(t)
		is.Equal(percentage, res.Percentage)
		is.Nil(err)
	})
}
