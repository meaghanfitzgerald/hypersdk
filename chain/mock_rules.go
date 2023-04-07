// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/hypersdk/chain (interfaces: Rules)

// Package chain is a generated GoMock package.
package chain

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	gomock "github.com/golang/mock/gomock"
)

// MockRules is a mock of Rules interface.
type MockRules struct {
	ctrl     *gomock.Controller
	recorder *MockRulesMockRecorder
}

// MockRulesMockRecorder is the mock recorder for MockRules.
type MockRulesMockRecorder struct {
	mock *MockRules
}

// NewMockRules creates a new mock instance.
func NewMockRules(ctrl *gomock.Controller) *MockRules {
	mock := &MockRules{ctrl: ctrl}
	mock.recorder = &MockRulesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRules) EXPECT() *MockRulesMockRecorder {
	return m.recorder
}

// FetchCustom mocks base method.
func (m *MockRules) FetchCustom(arg0 string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCustom", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FetchCustom indicates an expected call of FetchCustom.
func (mr *MockRulesMockRecorder) FetchCustom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCustom", reflect.TypeOf((*MockRules)(nil).FetchCustom), arg0)
}

// GetBaseUnits mocks base method.
func (m *MockRules) GetBaseUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetBaseUnits indicates an expected call of GetBaseUnits.
func (mr *MockRulesMockRecorder) GetBaseUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseUnits", reflect.TypeOf((*MockRules)(nil).GetBaseUnits))
}

// GetBlockCostChangeDenominator mocks base method.
func (m *MockRules) GetBlockCostChangeDenominator() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockCostChangeDenominator")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetBlockCostChangeDenominator indicates an expected call of GetBlockCostChangeDenominator.
func (mr *MockRulesMockRecorder) GetBlockCostChangeDenominator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockCostChangeDenominator", reflect.TypeOf((*MockRules)(nil).GetBlockCostChangeDenominator))
}

// GetMaxBlockTxs mocks base method.
func (m *MockRules) GetMaxBlockTxs() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxBlockTxs")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxBlockTxs indicates an expected call of GetMaxBlockTxs.
func (mr *MockRulesMockRecorder) GetMaxBlockTxs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxBlockTxs", reflect.TypeOf((*MockRules)(nil).GetMaxBlockTxs))
}

// GetMaxBlockUnits mocks base method.
func (m *MockRules) GetMaxBlockUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxBlockUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetMaxBlockUnits indicates an expected call of GetMaxBlockUnits.
func (mr *MockRulesMockRecorder) GetMaxBlockUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxBlockUnits", reflect.TypeOf((*MockRules)(nil).GetMaxBlockUnits))
}

// GetMaxChunks mocks base method.
func (m *MockRules) GetMaxChunks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxChunks")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxChunks indicates an expected call of GetMaxChunks.
func (mr *MockRulesMockRecorder) GetMaxChunks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxChunks", reflect.TypeOf((*MockRules)(nil).GetMaxChunks))
}

// GetMinBlockCost mocks base method.
func (m *MockRules) GetMinBlockCost() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinBlockCost")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetMinBlockCost indicates an expected call of GetMinBlockCost.
func (mr *MockRulesMockRecorder) GetMinBlockCost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinBlockCost", reflect.TypeOf((*MockRules)(nil).GetMinBlockCost))
}

// GetMinUnitPrice mocks base method.
func (m *MockRules) GetMinUnitPrice() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinUnitPrice")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetMinUnitPrice indicates an expected call of GetMinUnitPrice.
func (mr *MockRulesMockRecorder) GetMinUnitPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinUnitPrice", reflect.TypeOf((*MockRules)(nil).GetMinUnitPrice))
}

// GetUnitPriceChangeDenominator mocks base method.
func (m *MockRules) GetUnitPriceChangeDenominator() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitPriceChangeDenominator")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetUnitPriceChangeDenominator indicates an expected call of GetUnitPriceChangeDenominator.
func (mr *MockRulesMockRecorder) GetUnitPriceChangeDenominator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitPriceChangeDenominator", reflect.TypeOf((*MockRules)(nil).GetUnitPriceChangeDenominator))
}

// GetValidityWindow mocks base method.
func (m *MockRules) GetValidityWindow() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidityWindow")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetValidityWindow indicates an expected call of GetValidityWindow.
func (mr *MockRulesMockRecorder) GetValidityWindow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidityWindow", reflect.TypeOf((*MockRules)(nil).GetValidityWindow))
}

// GetWarpBaseFee mocks base method.
func (m *MockRules) GetWarpBaseFee() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWarpBaseFee")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetWarpBaseFee indicates an expected call of GetWarpBaseFee.
func (mr *MockRulesMockRecorder) GetWarpBaseFee() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWarpBaseFee", reflect.TypeOf((*MockRules)(nil).GetWarpBaseFee))
}

// GetWarpConfig mocks base method.
func (m *MockRules) GetWarpConfig(arg0 ids.ID) (bool, uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWarpConfig", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	return ret0, ret1, ret2
}

// GetWarpConfig indicates an expected call of GetWarpConfig.
func (mr *MockRulesMockRecorder) GetWarpConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWarpConfig", reflect.TypeOf((*MockRules)(nil).GetWarpConfig), arg0)
}

// GetWarpFeePerSigner mocks base method.
func (m *MockRules) GetWarpFeePerSigner() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWarpFeePerSigner")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetWarpFeePerSigner indicates an expected call of GetWarpFeePerSigner.
func (mr *MockRulesMockRecorder) GetWarpFeePerSigner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWarpFeePerSigner", reflect.TypeOf((*MockRules)(nil).GetWarpFeePerSigner))
}

// GetWindowTargetBlocks mocks base method.
func (m *MockRules) GetWindowTargetBlocks() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWindowTargetBlocks")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetWindowTargetBlocks indicates an expected call of GetWindowTargetBlocks.
func (mr *MockRulesMockRecorder) GetWindowTargetBlocks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWindowTargetBlocks", reflect.TypeOf((*MockRules)(nil).GetWindowTargetBlocks))
}

// GetWindowTargetUnits mocks base method.
func (m *MockRules) GetWindowTargetUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWindowTargetUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetWindowTargetUnits indicates an expected call of GetWindowTargetUnits.
func (mr *MockRulesMockRecorder) GetWindowTargetUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWindowTargetUnits", reflect.TypeOf((*MockRules)(nil).GetWindowTargetUnits))
}
