// Code generated by MockGen. DO NOT EDIT.
// Source: certificates.go

// Package mocks is a generated GoMock package.
package mocks

import (
	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCertificatesService is a mock of CertificatesService interface
type MockCertificatesService struct {
	ctrl     *gomock.Controller
	recorder *MockCertificatesServiceMockRecorder
}

// MockCertificatesServiceMockRecorder is the mock recorder for MockCertificatesService
type MockCertificatesServiceMockRecorder struct {
	mock *MockCertificatesService
}

// NewMockCertificatesService creates a new mock instance
func NewMockCertificatesService(ctrl *gomock.Controller) *MockCertificatesService {
	mock := &MockCertificatesService{ctrl: ctrl}
	mock.recorder = &MockCertificatesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificatesService) EXPECT() *MockCertificatesServiceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockCertificatesService) Get(cID string) (*do.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", cID)
	ret0, _ := ret[0].(*do.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCertificatesServiceMockRecorder) Get(cID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCertificatesService)(nil).Get), cID)
}

// Create mocks base method
func (m *MockCertificatesService) Create(cr *godo.CertificateRequest) (*do.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cr)
	ret0, _ := ret[0].(*do.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCertificatesServiceMockRecorder) Create(cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCertificatesService)(nil).Create), cr)
}

// List mocks base method
func (m *MockCertificatesService) List() (do.Certificates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(do.Certificates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCertificatesServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCertificatesService)(nil).List))
}

// Delete mocks base method
func (m *MockCertificatesService) Delete(cID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", cID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCertificatesServiceMockRecorder) Delete(cID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCertificatesService)(nil).Delete), cID)
}
