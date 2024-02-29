package vehicle_test

import (
	"context"
	"errors"
	"testing"

	"github.com/LucasMateus-eng/operations-service/config"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	vehicle_mocks "github.com/LucasMateus-eng/operations-service/internal/mocks/vehicle"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

var (
	errMocked       = errors.New("some error")
	mockedContext   = context.Background()
	expectedVehicle = &vehicle.Vehicle{
		ID: 1,
	}
	expectedVehicles = &[]vehicle.Vehicle{
		*expectedVehicle,
	}
)

func TestService_GetByID(t *testing.T) {
	type serviceMocks struct {
		repo   *vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *vehicle.Vehicle
		wantErr     bool
	}{
		{
			name: "Dado um ID válido quando o método GetByID é chamado então o veículo é retornado",
			args: args{
				ctx: mockedContext,
				id:  1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByID(p.ctx, p.id).Return(expectedVehicle, nil)
			},
			want:    expectedVehicle,
			wantErr: false,
		},
		{
			name: "Dado um ID inválido quando o método GetByID é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				id:  0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByID(p.ctx, p.id).Return(nil, errMocked)
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
				repo:   vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := vehicle.NewService(sm.repo, sm.logger)

			actualVehicle, err := s.GetByID(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualVehicle)
		})
	}
}

func TestService_GetByPlate(t *testing.T) {
	type serviceMocks struct {
		repo   *vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx   context.Context
		plate string
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *vehicle.Vehicle
		wantErr     bool
	}{
		{
			name: "Dado uma placa válida quando o método GetByPlate é chamado então o veículo é retornado",
			args: args{
				ctx:   mockedContext,
				plate: "ABC-1234",
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByPlate(p.ctx, p.plate).Return(expectedVehicle, nil)
			},
			want:    expectedVehicle,
			wantErr: false,
		},
		{
			name: "Dado uma placa inválida quando o método GetByPlate é chamado então um erro é retornado",
			args: args{
				ctx:   mockedContext,
				plate: "INVALID-PLATE",
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByPlate(p.ctx, p.plate).Return(nil, errMocked)
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
				repo:   vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := vehicle.NewService(sm.repo, sm.logger)

			actualVehicle, err := s.GetByPlate(test.args.ctx, test.args.plate)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualVehicle)
		})
	}
}

func TestService_GetByRenavam(t *testing.T) {
	type serviceMocks struct {
		repo   *vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx     context.Context
		renavam string
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *vehicle.Vehicle
		wantErr     bool
	}{
		{
			name: "Dado um Renavam válido quando o método GetByRenavam é chamado então o veículo é retornado",
			args: args{
				ctx:     mockedContext,
				renavam: "ABC1234",
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByRenavam(p.ctx, p.renavam).Return(expectedVehicle, nil)
			},
			want:    expectedVehicle,
			wantErr: false,
		},
		{
			name: "Dado um Renavam inválido quando o método GetByRenavam é chamado então um erro é retornado",
			args: args{
				ctx:     mockedContext,
				renavam: "INVALID-RENAVAM",
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByRenavam(p.ctx, p.renavam).Return(nil, errMocked)
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
				repo:   vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := vehicle.NewService(sm.repo, sm.logger)

			actualVehicle, err := s.GetByRenavam(test.args.ctx, test.args.renavam)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualVehicle)
		})
	}
}

func TestService_List(t *testing.T) {
	type serviceMocks struct {
		repo   *vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx           context.Context
		specification *vehicle.VehicleSpectification
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *[]vehicle.Vehicle
		wantErr     bool
	}{
		{
			name: "Dado uma especificação válida quando o método List é chamado então a lista de veículos é retornada",
			args: args{
				ctx: mockedContext,
				specification: &vehicle.VehicleSpectification{
					Page:     1,
					PageSize: 10,
				},
			},
			prepareMock: func(p args, m serviceMocks) {

				m.repo.EXPECT().List(p.ctx, p.specification).Return(expectedVehicles, nil)
			},
			want:    expectedVehicles,
			wantErr: false,
		},
		{
			name: "Dado uma especificação inválida quando o método List é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				specification: &vehicle.VehicleSpectification{
					Page:     0,
					PageSize: 0,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().List(p.ctx, p.specification).Return(nil, errMocked)
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
				repo:   vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := vehicle.NewService(sm.repo, sm.logger)

			actualVehicles, err := s.List(test.args.ctx, test.args.specification)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualVehicles)
		})
	}
}

func TestService_Create(t *testing.T) {
	type serviceMocks struct {
		repo   *vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		v   *vehicle.Vehicle
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        int64
		wantErr     bool
	}{
		{
			name: "Dado um veículo válido quando o método Create é chamado então o ID do veículo é retornado",
			args: args{
				ctx: mockedContext,
				v:   &vehicle.Vehicle{},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.v).Return(int64(1), nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Dado um veículo inválido quando o método Create é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				v:   &vehicle.Vehicle{},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.v).Return(int64(0), errMocked)
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := vehicle.NewService(sm.repo, sm.logger)

			actualVehicleID, err := s.Create(test.args.ctx, test.args.v)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualVehicleID)
		})
	}
}

func TestService_Update(t *testing.T) {
	type serviceMocks struct {
		repo   *vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		v   *vehicle.Vehicle
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		wantErr     bool
	}{
		{
			name: "Dado um veículo válido quando o método Update é chamado então o veículo é atualizado",
			args: args{
				ctx: mockedContext,
				v:   &vehicle.Vehicle{ID: 1},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.v).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Dado um veículo inválido quando o método Update é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				v:   &vehicle.Vehicle{},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.v).Return(errMocked)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := vehicle.NewService(sm.repo, sm.logger)

			err := s.Update(test.args.ctx, test.args.v)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}

func TestService_Delete(t *testing.T) {
	type serviceMocks struct {
		repo   *vehicle_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		wantErr     bool
	}{
		{
			name: "Dado um ID válido quando o método Delete é chamado então o veículo é excluído",
			args: args{
				ctx: mockedContext,
				id:  1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Delete(p.ctx, p.id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Dado um ID inválido quando o método Delete é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				id:  0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Delete(p.ctx, p.id).Return(errMocked)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   vehicle_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := vehicle.NewService(sm.repo, sm.logger)

			err := s.Delete(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}
