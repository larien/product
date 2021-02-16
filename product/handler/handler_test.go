package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/larien/product/product/drivers/request"
	"github.com/larien/product/product/entity"
	"github.com/larien/product/product/handler"
	handlerMock "github.com/larien/product/product/handler/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	baseURL = "/v1/product"
	null    = "null\n"
)

func TestListProduct(t *testing.T) {
	t.Parallel()
	t.Run("an error occurred", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		req, err := http.NewRequest(http.MethodGet, baseURL, nil)
		is.Nil(err)
		rec := httptest.NewRecorder()

		controller := handlerMock.New()
		expectedError := errors.New(faker.Sentence())
		controller.On("List", mock.Anything, mock.Anything).Return(nil, expectedError).Once()

		handler := handler.New(controller)
		handler.ServeHTTP(rec, req)

		controller.AssertExpectations(t)
		is.Equal(http.StatusInternalServerError, rec.Code)
		var body request.ErrorResponse
		is.Nil(json.Unmarshal(rec.Body.Bytes(), &body))
		is.Equal(body.Error, expectedError.Error())
	})
	t.Run("listing was successful", func(t *testing.T) {
		t.Parallel()
		t.Run("but no product was found", func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, baseURL, nil)
			is.Nil(err)
			rec := httptest.NewRecorder()

			controller := handlerMock.New()
			controller.On("List", mock.Anything, mock.Anything).Return(nil, nil).Once()

			handler := handler.New(controller)
			handler.ServeHTTP(rec, req)

			controller.AssertExpectations(t)
			is.Equal(http.StatusNoContent, rec.Code)
			is.Equal(null, rec.Body.String())
		})
		t.Run("and returned three products", func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			req, err := http.NewRequest(http.MethodGet, baseURL, nil)
			is.Nil(err)
			rec := httptest.NewRecorder()

			controller := handlerMock.New()
			product1 := entity.Product{}
			is.Nil(faker.FakeData(&product1))
			product2 := entity.Product{}
			is.Nil(faker.FakeData(&product2))
			product3 := entity.Product{}
			is.Nil(faker.FakeData(&product3))
			expectedProducts := []*entity.Product{&product1, &product2, &product3}
			is.Nil(faker.FakeData(&expectedProducts))
			controller.On("List", mock.Anything, mock.Anything).Return(expectedProducts, nil).Once()

			handler := handler.New(controller)
			handler.ServeHTTP(rec, req)

			controller.AssertExpectations(t)
			is.Equal(http.StatusOK, rec.Code)
			var products []entity.Product
			is.Nil(json.Unmarshal(rec.Body.Bytes(), &products))
			is.Len(products, len(expectedProducts))
		})
	})
}

func TestHealthcheck(t *testing.T) {
	t.Parallel()
	t.Run("service is up", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)
		req, err := http.NewRequest(http.MethodGet, "/status", nil)
		is.Nil(err)
		rec := httptest.NewRecorder()

		handler := handler.New(nil)
		handler.ServeHTTP(rec, req)
		is.Equal(http.StatusOK, rec.Code)
	})
}
