// Code generated by MockGen. DO NOT EDIT.
// Source: iface/filter_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	iface "github.com/ytbiu/tool/filter/iface"
	reflect "reflect"
)

// MockFilter is a mock of Filter interface
type MockFilter struct {
	ctrl     *gomock.Controller
	recorder *MockFilterMockRecorder
}

// MockFilterMockRecorder is the mock recorder for MockFilter
type MockFilterMockRecorder struct {
	mock *MockFilter
}

// NewMockFilter creates a new mock instance
func NewMockFilter(ctrl *gomock.Controller) *MockFilter {
	mock := &MockFilter{ctrl: ctrl}
	mock.recorder = &MockFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilter) EXPECT() *MockFilterMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockFilter) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockFilterMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFilter)(nil).Name))
}

// Check mocks base method
func (m *MockFilter) Check(arg0 interface{}) bool {
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockFilterMockRecorder) Check(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockFilter)(nil).Check), arg0)
}

// Next mocks base method
func (m *MockFilter) Next() iface.Filter {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(iface.Filter)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockFilterMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockFilter)(nil).Next))
}

// SetNext mocks base method
func (m *MockFilter) SetNext(arg0 iface.Filter) {
	m.ctrl.Call(m, "SetNext", arg0)
}

// SetNext indicates an expected call of SetNext
func (mr *MockFilterMockRecorder) SetNext(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNext", reflect.TypeOf((*MockFilter)(nil).SetNext), arg0)
}

// MockNode is a mock of Node interface
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockNode) Next() iface.Filter {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(iface.Filter)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockNodeMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockNode)(nil).Next))
}

// SetNext mocks base method
func (m *MockNode) SetNext(arg0 iface.Filter) {
	m.ctrl.Call(m, "SetNext", arg0)
}

// SetNext indicates an expected call of SetNext
func (mr *MockNodeMockRecorder) SetNext(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNext", reflect.TypeOf((*MockNode)(nil).SetNext), arg0)
}
