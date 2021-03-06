// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DiscountMock is an autogenerated mock type for the Discount type
type DiscountMock struct {
	mock.Mock
}

// NewDiscount creates a new instance of Discount repository mock
func NewDiscount() *DiscountMock {
	return &DiscountMock{}
}

// Get provides a mock function with given fields: ctx, productID, userID
func (_m *DiscountMock) Get(ctx context.Context, productID string, userID string) (int64, error) {
	ret := _m.Called(ctx, productID, userID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, productID, userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, productID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
