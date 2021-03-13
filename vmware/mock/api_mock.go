// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"
	api "vmware/pkg/api"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockClient) Get(url string) (*api.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url)
	ret0, _ := ret[0].(*api.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), url)
}

// GetTimeout mocks base method.
func (m *MockClient) GetTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetTimeout indicates an expected call of GetTimeout.
func (mr *MockClientMockRecorder) GetTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeout", reflect.TypeOf((*MockClient)(nil).GetTimeout))
}

// GetURL mocks base method.
func (m *MockClient) GetURL(path string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", path)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetURL indicates an expected call of GetURL.
func (mr *MockClientMockRecorder) GetURL(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockClient)(nil).GetURL), path)
}
