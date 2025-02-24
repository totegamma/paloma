// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"

	valsettypes "github.com/palomachain/paloma/x/valset/types"
)

// ValsetKeeper is an autogenerated mock type for the ValsetKeeper type
type ValsetKeeper struct {
	mock.Mock
}

// FindSnapshotByID provides a mock function with given fields: ctx, id
func (_m *ValsetKeeper) FindSnapshotByID(ctx types.Context, id uint64) (*valsettypes.Snapshot, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindSnapshotByID")
	}

	var r0 *valsettypes.Snapshot
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (*valsettypes.Snapshot, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) *valsettypes.Snapshot); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*valsettypes.Snapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllChainInfos provides a mock function with given fields: ctx
func (_m *ValsetKeeper) GetAllChainInfos(ctx types.Context) ([]*valsettypes.ValidatorExternalAccounts, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllChainInfos")
	}

	var r0 []*valsettypes.ValidatorExternalAccounts
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) ([]*valsettypes.ValidatorExternalAccounts, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) []*valsettypes.ValidatorExternalAccounts); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*valsettypes.ValidatorExternalAccounts)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentSnapshot provides a mock function with given fields: ctx
func (_m *ValsetKeeper) GetCurrentSnapshot(ctx types.Context) (*valsettypes.Snapshot, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentSnapshot")
	}

	var r0 *valsettypes.Snapshot
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context) (*valsettypes.Snapshot, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) *valsettypes.Snapshot); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*valsettypes.Snapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestSnapshotOnChain provides a mock function with given fields: ctx, chainReferenceID
func (_m *ValsetKeeper) GetLatestSnapshotOnChain(ctx types.Context, chainReferenceID string) (*valsettypes.Snapshot, error) {
	ret := _m.Called(ctx, chainReferenceID)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestSnapshotOnChain")
	}

	var r0 *valsettypes.Snapshot
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, string) (*valsettypes.Snapshot, error)); ok {
		return rf(ctx, chainReferenceID)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) *valsettypes.Snapshot); ok {
		r0 = rf(ctx, chainReferenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*valsettypes.Snapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) error); ok {
		r1 = rf(ctx, chainReferenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidatorChainInfos provides a mock function with given fields: ctx, valAddr
func (_m *ValsetKeeper) GetValidatorChainInfos(ctx types.Context, valAddr types.ValAddress) ([]*valsettypes.ExternalChainInfo, error) {
	ret := _m.Called(ctx, valAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetValidatorChainInfos")
	}

	var r0 []*valsettypes.ExternalChainInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) ([]*valsettypes.ExternalChainInfo, error)); ok {
		return rf(ctx, valAddr)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) []*valsettypes.ExternalChainInfo); ok {
		r0 = rf(ctx, valAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*valsettypes.ExternalChainInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.ValAddress) error); ok {
		r1 = rf(ctx, valAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsJailed provides a mock function with given fields: ctx, val
func (_m *ValsetKeeper) IsJailed(ctx types.Context, val types.ValAddress) bool {
	ret := _m.Called(ctx, val)

	if len(ret) == 0 {
		panic("no return value specified for IsJailed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) bool); ok {
		r0 = rf(ctx, val)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Jail provides a mock function with given fields: ctx, valAddr, reason
func (_m *ValsetKeeper) Jail(ctx types.Context, valAddr types.ValAddress, reason string) error {
	ret := _m.Called(ctx, valAddr, reason)

	if len(ret) == 0 {
		panic("no return value specified for Jail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress, string) error); ok {
		r0 = rf(ctx, valAddr, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KeepValidatorAlive provides a mock function with given fields: ctx, valAddr, pigeonVersion
func (_m *ValsetKeeper) KeepValidatorAlive(ctx types.Context, valAddr types.ValAddress, pigeonVersion string) error {
	ret := _m.Called(ctx, valAddr, pigeonVersion)

	if len(ret) == 0 {
		panic("no return value specified for KeepValidatorAlive")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress, string) error); ok {
		r0 = rf(ctx, valAddr, pigeonVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSnapshotOnChain provides a mock function with given fields: ctx, snapshotID, chainReferenceID
func (_m *ValsetKeeper) SetSnapshotOnChain(ctx types.Context, snapshotID uint64, chainReferenceID string) error {
	ret := _m.Called(ctx, snapshotID, chainReferenceID)

	if len(ret) == 0 {
		panic("no return value specified for SetSnapshotOnChain")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, string) error); ok {
		r0 = rf(ctx, snapshotID, chainReferenceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetValidatorBalance provides a mock function with given fields: ctx, valAddr, chainType, chainReferenceID, externalAddress, balance
func (_m *ValsetKeeper) SetValidatorBalance(ctx types.Context, valAddr types.ValAddress, chainType string, chainReferenceID string, externalAddress string, balance *big.Int) error {
	ret := _m.Called(ctx, valAddr, chainType, chainReferenceID, externalAddress, balance)

	if len(ret) == 0 {
		panic("no return value specified for SetValidatorBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress, string, string, string, *big.Int) error); ok {
		r0 = rf(ctx, valAddr, chainType, chainReferenceID, externalAddress, balance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewValsetKeeper creates a new instance of ValsetKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValsetKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *ValsetKeeper {
	mock := &ValsetKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
