// Code generated by mockery v2.36.1. DO NOT EDIT.

package storagemocks

import (
	context "context"
	mailing "mail_service/internal"

	mock "github.com/stretchr/testify/mock"
)

// ClienteRepository is an autogenerated mock type for the ClienteRepository type
type ClienteRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx
func (_m *ClienteRepository) Get(ctx context.Context) ([]struct {
	Client mailing.Cliente
	Plan   mailing.Plan
}, error) {
	ret := _m.Called(ctx)

	var r0 []struct {
		Client mailing.Cliente
		Plan   mailing.Plan
	}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]struct {
		Client mailing.Cliente
		Plan   mailing.Plan
	}, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []struct {
		Client mailing.Cliente
		Plan   mailing.Plan
	}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]struct {
				Client mailing.Cliente
				Plan   mailing.Plan
			})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, cliente
func (_m *ClienteRepository) Save(ctx context.Context, cliente mailing.Cliente) error {
	ret := _m.Called(ctx, cliente)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mailing.Cliente) error); ok {
		r0 = rf(ctx, cliente)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClienteRepository creates a new instance of ClienteRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClienteRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClienteRepository {
	mock := &ClienteRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}