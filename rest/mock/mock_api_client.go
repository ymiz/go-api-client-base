// Code generated by MockGen. DO NOT EDIT.
// Source: api_client.go

// Package mock_rest is a generated GoMock package.
package mock_rest

import (
	gomock "github.com/golang/mock/gomock"
	rest "go-api-client-base/rest"
	reflect "reflect"
)

// MockApiClient is a mock of ApiClient interface
type MockApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientMockRecorder
}

// MockApiClientMockRecorder is the mock recorder for MockApiClient
type MockApiClientMockRecorder struct {
	mock *MockApiClient
}

// NewMockApiClient creates a new mock instance
func NewMockApiClient(ctrl *gomock.Controller) *MockApiClient {
	mock := &MockApiClient{ctrl: ctrl}
	mock.recorder = &MockApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiClient) EXPECT() *MockApiClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockApiClient) Get(arg0 string) (*rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockApiClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApiClient)(nil).Get), arg0)
}
