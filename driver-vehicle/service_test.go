package drivervehicle_test

import (
	"context"
	"errors"
	"testing"

	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/config"
	"github.com/LucasMateus-eng/operations-service/driver"
	drivervehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	driver_vehicle_mocks "github.com/LucasMateus-eng/operations-service/internal/mocks/driver-vehicle"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

var (
	errMocked             = errors.New("some error")
	mockedContext         = context.Background()
	expectedDriverVehicle = &drivervehicle.DriverVehicle{DriverID: 1, VehicleID: 1}
	expectedAddress       = &address.Address{ID: 1}
	expectedVehicle       = &vehicle.Vehicle{
		ID: 1,
	}
	expectedVehicles = &[]vehicle.Vehicle{
		*expectedVehicle,
	}
	expectedDriverWithEagerLoading = &driver.Driver{
		ID:       1,
		Address:  expectedAddress,
		Vehicles: *expectedVehicles,
	}
	expectedDriversWithEagerLoading = &[]driver.Driver{
		*expectedDriverWithEagerLoading,
	}
)

func TestService_GetByID(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx                 context.Context
		driverID, vehicleID int64
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *drivervehicle.DriverVehicle
		wantErr     bool
	}{
		{
			name: "Dado um ID válido de motorista e veículo quando o método GetByID é chamado então a relação motorista/veículo é retornada",
			args: args{
				ctx:       mockedContext,
				driverID:  1,
				vehicleID: 1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByID(p.ctx, p.driverID, p.vehicleID).Return(expectedDriverVehicle, nil)
			},
			want:    expectedDriverVehicle,
			wantErr: false,
		},
		{
			name: "Dado um ID inválido de motorista e veículo quando o método GetByID é chamado então um erro é retornado",
			args: args{
				ctx:       mockedContext,
				driverID:  0,
				vehicleID: 0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByID(p.ctx, p.driverID, p.vehicleID).Return(nil, errMocked)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   driver_vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := drivervehicle.NewService(sm.repo, sm.logger)

			actualDriver, err := s.GetByID(test.args.ctx, test.args.driverID, test.args.vehicleID)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriver)
		})
	}
}

func TestService_GetDriverListByVehicleID(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx           context.Context
		specification *drivervehicle.DriverVehicleSpecification
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *[]driver.Driver
		wantErr     bool
	}{
		{
			name: "Dado uma especificação válida quando o método GetDriverListByVehicleID é chamado então a lista de motoristas é retornada",
			args: args{
				ctx: mockedContext,
				specification: &drivervehicle.DriverVehicleSpecification{
					VehicleID: 1,
					Page:      1,
					PageSize:  10,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetDriverListByVehicleID(p.ctx, p.specification).Return(expectedDriversWithEagerLoading, nil)
			},
			want:    expectedDriversWithEagerLoading,
			wantErr: false,
		},
		{
			name: "Dado uma especificação inválida quando o método GetDriverListByVehicleID é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				specification: &drivervehicle.DriverVehicleSpecification{
					VehicleID: 1,
					Page:      0,
					PageSize:  0,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetDriverListByVehicleID(p.ctx, p.specification).Return(nil, errMocked)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   driver_vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := drivervehicle.NewService(sm.repo, sm.logger)

			actualDrivers, err := s.GetDriverListByVehicleID(test.args.ctx, test.args.specification)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDrivers)
		})
	}
}

func TestService_GetVehicleListByDriverID(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx           context.Context
		specification *drivervehicle.DriverVehicleSpecification
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *[]vehicle.Vehicle
		wantErr     bool
	}{
		{
			name: "Dado uma especificação válida quando o método GetVehicleListByDriverID é chamado então a lista de veículos é retornada",
			args: args{
				ctx: mockedContext,
				specification: &drivervehicle.DriverVehicleSpecification{
					DriverID: 1,
					Page:     1,
					PageSize: 10,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetVehicleListByDriverID(p.ctx, p.specification).Return(expectedVehicles, nil)
			},
			want:    expectedVehicles,
			wantErr: false,
		},
		{
			name: "Dado uma especificação inválida quando o método GetVehicleListByDriverID é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				specification: &drivervehicle.DriverVehicleSpecification{
					DriverID: 0,
					Page:     0,
					PageSize: 0,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetVehicleListByDriverID(p.ctx, p.specification).Return(nil, errMocked)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   driver_vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := drivervehicle.NewService(sm.repo, sm.logger)

			actualVehicles, err := s.GetVehicleListByDriverID(test.args.ctx, test.args.specification)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualVehicles)
		})
	}
}

func TestService_Create(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		dv  *drivervehicle.DriverVehicle
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *drivervehicle.DriverVehicle
		wantErr     bool
	}{
		{
			name: "Dado um DriverVehicle válido quando o método Create é chamado então a relação motorista/veículo é criada",
			args: args{
				ctx: mockedContext,
				dv:  &drivervehicle.DriverVehicle{DriverID: 1, VehicleID: 1},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.dv).Return(expectedDriverVehicle, nil)
			},
			want:    expectedDriverVehicle,
			wantErr: false,
		},
		{
			name: "Dado um DriverVehicle inválido quando o método Create é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				dv:  &drivervehicle.DriverVehicle{},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.dv).Return(nil, errMocked)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   driver_vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := drivervehicle.NewService(sm.repo, sm.logger)

			actualDriverVehicle, err := s.Create(test.args.ctx, test.args.dv)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriverVehicle)
		})
	}
}

func TestService_Delete(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx                 context.Context
		driverID, vehicleID int64
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		wantErr     bool
	}{
		{
			name: "Dado um ID válido de motorista e veículo quando o método Delete é chamado então a relação é removida com sucesso",
			args: args{
				ctx:       mockedContext,
				driverID:  1,
				vehicleID: 1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Delete(p.ctx, p.driverID, p.vehicleID).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Dado um ID inválido de motorista e veículo quando o método Delete é chamado então um erro é retornado",
			args: args{
				ctx:       mockedContext,
				driverID:  0,
				vehicleID: 0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Delete(p.ctx, p.driverID, p.vehicleID).Return(errMocked)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   driver_vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := drivervehicle.NewService(sm.repo, sm.logger)

			err := s.Delete(test.args.ctx, test.args.driverID, test.args.vehicleID)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}
