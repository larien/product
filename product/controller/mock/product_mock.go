// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	entity "github.com/larien/product/product/entity"
	mock "github.com/stretchr/testify/mock"
)

// ProductMock is an autogenerated mock type for the Product type
type ProductMock struct {
	mock.Mock
}

// NewProduct creates a new instance of Discount repository mock
func NewProduct() *ProductMock {
	return &ProductMock{}
}

// List provides a mock function
func (_m *ProductMock) List() ([]entity.Product, error) {
	ret := _m.Called()

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func() []entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
