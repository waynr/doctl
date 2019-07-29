// Code generated by MockGen. DO NOT EDIT.
// Source: volumes.go

// Package mocks is a generated GoMock package.
package mocks

import (
	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVolumesService is a mock of VolumesService interface
type MockVolumesService struct {
	ctrl     *gomock.Controller
	recorder *MockVolumesServiceMockRecorder
}

// MockVolumesServiceMockRecorder is the mock recorder for MockVolumesService
type MockVolumesServiceMockRecorder struct {
	mock *MockVolumesService
}

// NewMockVolumesService creates a new mock instance
func NewMockVolumesService(ctrl *gomock.Controller) *MockVolumesService {
	mock := &MockVolumesService{ctrl: ctrl}
	mock.recorder = &MockVolumesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVolumesService) EXPECT() *MockVolumesServiceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockVolumesService) List() ([]do.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]do.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVolumesServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVolumesService)(nil).List))
}

// CreateVolume mocks base method
func (m *MockVolumesService) CreateVolume(arg0 *godo.VolumeCreateRequest) (*do.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0)
	ret0, _ := ret[0].(*do.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockVolumesServiceMockRecorder) CreateVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockVolumesService)(nil).CreateVolume), arg0)
}

// DeleteVolume mocks base method
func (m *MockVolumesService) DeleteVolume(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockVolumesServiceMockRecorder) DeleteVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockVolumesService)(nil).DeleteVolume), arg0)
}

// Get mocks base method
func (m *MockVolumesService) Get(arg0 string) (*do.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*do.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVolumesServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVolumesService)(nil).Get), arg0)
}

// CreateSnapshot mocks base method
func (m *MockVolumesService) CreateSnapshot(arg0 *godo.SnapshotCreateRequest) (*do.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0)
	ret0, _ := ret[0].(*do.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot
func (mr *MockVolumesServiceMockRecorder) CreateSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockVolumesService)(nil).CreateSnapshot), arg0)
}

// GetSnapshot mocks base method
func (m *MockVolumesService) GetSnapshot(arg0 string) (*do.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", arg0)
	ret0, _ := ret[0].(*do.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshot indicates an expected call of GetSnapshot
func (mr *MockVolumesServiceMockRecorder) GetSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockVolumesService)(nil).GetSnapshot), arg0)
}

// DeleteSnapshot mocks base method
func (m *MockVolumesService) DeleteSnapshot(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot
func (mr *MockVolumesServiceMockRecorder) DeleteSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockVolumesService)(nil).DeleteSnapshot), arg0)
}

// ListSnapshots mocks base method
func (m *MockVolumesService) ListSnapshots(arg0 string, arg1 *godo.ListOptions) ([]do.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshots", arg0, arg1)
	ret0, _ := ret[0].([]do.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshots indicates an expected call of ListSnapshots
func (mr *MockVolumesServiceMockRecorder) ListSnapshots(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshots", reflect.TypeOf((*MockVolumesService)(nil).ListSnapshots), arg0, arg1)
}
