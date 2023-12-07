// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// GravityKeeper is an autogenerated mock type for the GravityKeeper type
type GravityKeeper struct {
	mock.Mock
}

// GetLastObservedEventNonce provides a mock function with given fields: ctx
func (_m *GravityKeeper) GetLastObservedEventNonce(ctx types.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastObservedEventNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGravityKeeper creates a new instance of GravityKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGravityKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *GravityKeeper {
	mock := &GravityKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}