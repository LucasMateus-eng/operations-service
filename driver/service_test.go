package driver_test

import (
	"context"
	"errors"
	"testing"

	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/config"
	"github.com/LucasMateus-eng/operations-service/driver"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	driver_mocks "github.com/LucasMateus-eng/operations-service/internal/mocks/driver"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

var (
	errMocked      = errors.New("some error")
	mockedContext  = context.Background()
	expectedDriver = &driver.Driver{
		ID: 1,
	}
	expectedDrivers = &[]driver.Driver{
		*expectedDriver,
	}
	expectedAddress = &address.Address{ID: 1}
	expectedVehicle = &vehicle.Vehicle{
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
		repo   *driver_mocks.MockRepository
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
		want        *driver.Driver
		wantErr     bool
	}{
		{
			name: "Dado um ID válido quando o método GetByID é chamado então o motorista é retornado",
			args: args{
				ctx: mockedContext,
				id:  1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByID(p.ctx, p.id).Return(expectedDriver, nil)
			},
			want:    expectedDriver,
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			actualDriver, err := s.GetByID(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriver)
		})
	}
}

func TestService_GetByUserID(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx    context.Context
		userId int64
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *driver.Driver
		wantErr     bool
	}{
		{
			name: "Dado um ID de usuário válido quando o método GetByUserID é chamado então o motorista é retornado",
			args: args{
				ctx:    mockedContext,
				userId: 1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUserID(p.ctx, p.userId).Return(expectedDriver, nil)
			},
			want:    expectedDriver,
			wantErr: false,
		},
		{
			name: "Dado um ID de usuário inválido quando o método GetByUserID é chamado então um erro é retornado",
			args: args{
				ctx:    mockedContext,
				userId: 0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUserID(p.ctx, p.userId).Return(nil, errMocked)
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			actualDriver, err := s.GetByUserID(test.args.ctx, test.args.userId)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriver)
		})
	}
}

func TestService_GetByIDWithEagerLoading(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
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
		want        *driver.Driver
		wantErr     bool
	}{
		{
			name: "Dado um ID válido quando o método GetByIDWithEagerLoading é chamado então o motorista com endereço e veículos é retornado",
			args: args{
				ctx: mockedContext,
				id:  1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByIDWithEagerLoading(p.ctx, p.id).Return(expectedDriverWithEagerLoading, nil)
			},
			want:    expectedDriverWithEagerLoading,
			wantErr: false,
		},
		{
			name: "Dado um ID inválido quando o método GetByIDWithEagerLoading é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				id:  0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByIDWithEagerLoading(p.ctx, p.id).Return(nil, errMocked)
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			actualDriver, err := s.GetByIDWithEagerLoading(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriver)
		})
	}
}

func TestService_GetByUserIDWithEagerLoading(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx    context.Context
		userId int64
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *driver.Driver
		wantErr     bool
	}{
		{
			name: "Dado um ID de usuário válido quando o método GetByUserIDWithEagerLoading é chamado então o motorista com endereço e veículos é retornado",
			args: args{
				ctx:    mockedContext,
				userId: 1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUserIDWithEagerLoading(p.ctx, p.userId).Return(expectedDriverWithEagerLoading, nil)
			},
			want:    expectedDriverWithEagerLoading,
			wantErr: false,
		},
		{
			name: "Dado um ID de usuário inválido quando o método GetByUserIDWithEagerLoading é chamado então um erro é retornado",
			args: args{
				ctx:    mockedContext,
				userId: 0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUserIDWithEagerLoading(p.ctx, p.userId).Return(nil, errMocked)
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			actualDriver, err := s.GetByUserIDWithEagerLoading(test.args.ctx, test.args.userId)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriver)
		})
	}
}

func TestService_List(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx           context.Context
		specification *driver.DriverSpecification
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *[]driver.Driver
		wantErr     bool
	}{
		{
			name: "Dado uma especificação válida quando o método List é chamado então a lista de motoristas é retornada",
			args: args{
				ctx: mockedContext,
				specification: &driver.DriverSpecification{
					Page:     1,
					PageSize: 10,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().List(p.ctx, p.specification).Return(expectedDrivers, nil)
			},
			want:    expectedDrivers,
			wantErr: false,
		},
		{
			name: "Dado uma especificação inválida quando o método List é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				specification: &driver.DriverSpecification{
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			actualDrivers, err := s.List(test.args.ctx, test.args.specification)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDrivers)
		})
	}
}

func TestService_ListWithEagerLoading(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx           context.Context
		specification *driver.DriverSpecification
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *[]driver.Driver
		wantErr     bool
	}{
		{
			name: "Dado uma especificação válida quando o método ListWithEagerLoading é chamado então a lista de motoristas com endereço e veículos é retornado",
			args: args{
				ctx: mockedContext,
				specification: &driver.DriverSpecification{
					Page:     1,
					PageSize: 10,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().ListWithEagerLoading(p.ctx, p.specification).Return(expectedDriversWithEagerLoading, nil)
			},
			want:    expectedDriversWithEagerLoading,
			wantErr: false,
		},
		{
			name: "Dado uma especificação inválida quando o método ListWithEagerLoading é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				specification: &driver.DriverSpecification{
					Page:     0,
					PageSize: 0,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().ListWithEagerLoading(p.ctx, p.specification).Return(nil, errMocked)
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			actualDrivers, err := s.ListWithEagerLoading(test.args.ctx, test.args.specification)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDrivers)
		})
	}
}

func TestService_Create(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		d   *driver.Driver
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        int64
		wantErr     bool
	}{
		{
			name: "Dado um driver válido quando o método Create é chamado então o ID do driver é retornado",
			args: args{
				ctx: mockedContext,
				d: &driver.Driver{
					UserID: 1,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.d).Return(int64(1), nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Dado um driver inválido quando o método Create é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				d: &driver.Driver{
					UserID: 0,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.d).Return(int64(0), errMocked)
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			actualDriverID, err := s.Create(test.args.ctx, test.args.d)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualDriverID)
		})
	}
}

func TestService_Update(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		d   *driver.Driver
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		wantErr     bool
	}{
		{
			name: "Dado um driver válido quando o método Update é chamado então o driver é atualizado",
			args: args{
				ctx: mockedContext,
				d:   &driver.Driver{ID: 1, Attributes: driver.DriverAttributes{Name: "Novo nome"}},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.d).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Dado um driver inválido quando o método Update é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				d:   &driver.Driver{},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.d).Return(errMocked)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			err := s.Update(test.args.ctx, test.args.d)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}

func TestService_Delete(t *testing.T) {
	type serviceMocks struct {
		repo   *driver_mocks.MockRepository
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
			name: "Dado um ID válido quando o método Delete é chamado então o motorista é removido",
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
				repo:   driver_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := driver.NewService(sm.repo, sm.logger)

			err := s.Delete(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}
