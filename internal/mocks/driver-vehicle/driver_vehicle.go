// Code generated by MockGen. DO NOT EDIT.
// Source: driver-vehicle/driver_vehicle.go
//
// Generated by this command:
//
//	mockgen -source=driver-vehicle/driver_vehicle.go -destination=internal/mocks/driver-vehicle/driver_vehicle.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	driver "github.com/LucasMateus-eng/operations-service/driver"
	drivervehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	vehicle "github.com/LucasMateus-eng/operations-service/vehicle"
	gomock "go.uber.org/mock/gomock"
)

// MockReading is a mock of Reading interface.
type MockReading struct {
	ctrl     *gomock.Controller
	recorder *MockReadingMockRecorder
}

// MockReadingMockRecorder is the mock recorder for MockReading.
type MockReadingMockRecorder struct {
	mock *MockReading
}

// NewMockReading creates a new mock instance.
func NewMockReading(ctrl *gomock.Controller) *MockReading {
	mock := &MockReading{ctrl: ctrl}
	mock.recorder = &MockReadingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReading) EXPECT() *MockReadingMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockReading) GetByID(ctx context.Context, driverID, vehicleID int64) (*drivervehicle.DriverVehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, driverID, vehicleID)
	ret0, _ := ret[0].(*drivervehicle.DriverVehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockReadingMockRecorder) GetByID(ctx, driverID, vehicleID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockReading)(nil).GetByID), ctx, driverID, vehicleID)
}

// GetDriverListByVehicleID mocks base method.
func (m *MockReading) GetDriverListByVehicleID(ctx context.Context, specification *drivervehicle.DriverVehicleSpecification) (*[]driver.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverListByVehicleID", ctx, specification)
	ret0, _ := ret[0].(*[]driver.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverListByVehicleID indicates an expected call of GetDriverListByVehicleID.
func (mr *MockReadingMockRecorder) GetDriverListByVehicleID(ctx, specification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverListByVehicleID", reflect.TypeOf((*MockReading)(nil).GetDriverListByVehicleID), ctx, specification)
}

// GetVehicleListByDriverID mocks base method.
func (m *MockReading) GetVehicleListByDriverID(ctx context.Context, specification *drivervehicle.DriverVehicleSpecification) (*[]vehicle.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVehicleListByDriverID", ctx, specification)
	ret0, _ := ret[0].(*[]vehicle.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVehicleListByDriverID indicates an expected call of GetVehicleListByDriverID.
func (mr *MockReadingMockRecorder) GetVehicleListByDriverID(ctx, specification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVehicleListByDriverID", reflect.TypeOf((*MockReading)(nil).GetVehicleListByDriverID), ctx, specification)
}

// MockWriting is a mock of Writing interface.
type MockWriting struct {
	ctrl     *gomock.Controller
	recorder *MockWritingMockRecorder
}

// MockWritingMockRecorder is the mock recorder for MockWriting.
type MockWritingMockRecorder struct {
	mock *MockWriting
}

// NewMockWriting creates a new mock instance.
func NewMockWriting(ctrl *gomock.Controller) *MockWriting {
	mock := &MockWriting{ctrl: ctrl}
	mock.recorder = &MockWritingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriting) EXPECT() *MockWritingMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWriting) Create(ctx context.Context, dv *drivervehicle.DriverVehicle) (*drivervehicle.DriverVehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dv)
	ret0, _ := ret[0].(*drivervehicle.DriverVehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWritingMockRecorder) Create(ctx, dv any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriting)(nil).Create), ctx, dv)
}

// Delete mocks base method.
func (m *MockWriting) Delete(ctx context.Context, driverID, vehicleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, driverID, vehicleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWritingMockRecorder) Delete(ctx, driverID, vehicleID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriting)(nil).Delete), ctx, driverID, vehicleID)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, dv *drivervehicle.DriverVehicle) (*drivervehicle.DriverVehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dv)
	ret0, _ := ret[0].(*drivervehicle.DriverVehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, dv any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, dv)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, driverID, vehicleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, driverID, vehicleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, driverID, vehicleID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, driverID, vehicleID)
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(ctx context.Context, driverID, vehicleID int64) (*drivervehicle.DriverVehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, driverID, vehicleID)
	ret0, _ := ret[0].(*drivervehicle.DriverVehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(ctx, driverID, vehicleID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), ctx, driverID, vehicleID)
}

// GetDriverListByVehicleID mocks base method.
func (m *MockRepository) GetDriverListByVehicleID(ctx context.Context, specification *drivervehicle.DriverVehicleSpecification) (*[]driver.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverListByVehicleID", ctx, specification)
	ret0, _ := ret[0].(*[]driver.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverListByVehicleID indicates an expected call of GetDriverListByVehicleID.
func (mr *MockRepositoryMockRecorder) GetDriverListByVehicleID(ctx, specification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverListByVehicleID", reflect.TypeOf((*MockRepository)(nil).GetDriverListByVehicleID), ctx, specification)
}

// GetVehicleListByDriverID mocks base method.
func (m *MockRepository) GetVehicleListByDriverID(ctx context.Context, specification *drivervehicle.DriverVehicleSpecification) (*[]vehicle.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVehicleListByDriverID", ctx, specification)
	ret0, _ := ret[0].(*[]vehicle.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVehicleListByDriverID indicates an expected call of GetVehicleListByDriverID.
func (mr *MockRepositoryMockRecorder) GetVehicleListByDriverID(ctx, specification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVehicleListByDriverID", reflect.TypeOf((*MockRepository)(nil).GetVehicleListByDriverID), ctx, specification)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCase) Create(ctx context.Context, dv *drivervehicle.DriverVehicle) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dv)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(ctx, dv any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), ctx, dv)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(ctx context.Context, driverID, vehicleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, driverID, vehicleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(ctx, driverID, vehicleID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), ctx, driverID, vehicleID)
}

// GetByID mocks base method.
func (m *MockUseCase) GetByID(ctx context.Context, driverID, vehicleID int64) (*drivervehicle.DriverVehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, driverID, vehicleID)
	ret0, _ := ret[0].(*drivervehicle.DriverVehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUseCaseMockRecorder) GetByID(ctx, driverID, vehicleID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUseCase)(nil).GetByID), ctx, driverID, vehicleID)
}

// GetDriverListByVehicleID mocks base method.
func (m *MockUseCase) GetDriverListByVehicleID(ctx context.Context, specification *drivervehicle.DriverVehicleSpecification) (*[]driver.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverListByVehicleID", ctx, specification)
	ret0, _ := ret[0].(*[]driver.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverListByVehicleID indicates an expected call of GetDriverListByVehicleID.
func (mr *MockUseCaseMockRecorder) GetDriverListByVehicleID(ctx, specification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverListByVehicleID", reflect.TypeOf((*MockUseCase)(nil).GetDriverListByVehicleID), ctx, specification)
}

// GetVehicleListByDriverID mocks base method.
func (m *MockUseCase) GetVehicleListByDriverID(ctx context.Context, specification *drivervehicle.DriverVehicleSpecification) (*[]vehicle.Vehicle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVehicleListByDriverID", ctx, specification)
	ret0, _ := ret[0].(*[]vehicle.Vehicle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVehicleListByDriverID indicates an expected call of GetVehicleListByDriverID.
func (mr *MockUseCaseMockRecorder) GetVehicleListByDriverID(ctx, specification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVehicleListByDriverID", reflect.TypeOf((*MockUseCase)(nil).GetVehicleListByDriverID), ctx, specification)
}
