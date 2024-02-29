package address_test

import (
	"context"
	"errors"
	"testing"

	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/config"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	address_mocks "github.com/LucasMateus-eng/operations-service/internal/mocks/address"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

var (
	errMocked       = errors.New("some error")
	mockedContext   = context.Background()
	expectedAddress = &address.Address{
		ID: 1,
	}
)

func TestService_GetByID(t *testing.T) {
	type serviceMocks struct {
		repo   *address_mocks.MockRepository
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
		want        *address.Address
		wantErr     bool
	}{
		{
			name: "Dado um ID válido quando o método GetByID é chamado então o endereço é retornado",
			args: args{
				ctx: mockedContext,
				id:  1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByID(p.ctx, p.id).Return(expectedAddress, nil)
			},
			want:    expectedAddress,
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
				repo:   address_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := address.NewService(sm.repo, sm.logger)

			actualUser, err := s.GetByID(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualUser)
		})
	}
}

func TestService_GetByUserID(t *testing.T) {
	type serviceMocks struct {
		repo   *address_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx    context.Context
		userID int64
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *address.Address
		wantErr     bool
	}{
		{
			name: "Dado um UserID válido quando o método GetByUserID é chamado então o endereço é retornado",
			args: args{
				ctx:    mockedContext,
				userID: 1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUserID(p.ctx, p.userID).Return(expectedAddress, nil)
			},
			want:    expectedAddress,
			wantErr: false,
		},
		{
			name: "Dado um UserID inválido quando o método GetByUserID é chamado então um erro é retornado",
			args: args{
				ctx:    mockedContext,
				userID: 0,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUserID(p.ctx, p.userID).Return(nil, errMocked)
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
				repo:   address_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := address.NewService(sm.repo, sm.logger)

			actualAddress, err := s.GetByUserID(test.args.ctx, test.args.userID)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualAddress)
		})
	}
}

func TestService_Create(t *testing.T) {
	type serviceMocks struct {
		repo   *address_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		a   *address.Address
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        int64
		wantErr     bool
	}{
		{
			name: "Dado um endereço válido quando o método Create é chamado então o endereço é criado",
			args: args{
				ctx: mockedContext,
				a:   &address.Address{ID: 1, UserID: 1, Locality: "Localidade Teste"},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.a).Return(int64(1), nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Dado um endereço inválido quando o método Create é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				a:   &address.Address{},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.a).Return(int64(0), errMocked)
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
				repo:   address_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := address.NewService(sm.repo, sm.logger)

			actualID, err := s.Create(test.args.ctx, test.args.a)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualID)
		})
	}
}

func TestService_Update(t *testing.T) {
	type serviceMocks struct {
		repo   *address_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		a   *address.Address
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		wantErr     bool
	}{
		{
			name: "Dado um endereço válido quando o método Update é chamado então o endereço é atualizado",
			args: args{
				ctx: mockedContext,
				a:   &address.Address{ID: 1, Locality: "Nova Localidade"},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.a).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Dado um endereço inválido quando o método Update é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				a:   &address.Address{ID: 0},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.a).Return(errMocked)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   address_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := address.NewService(sm.repo, sm.logger)

			err := s.Update(test.args.ctx, test.args.a)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}

func TestService_Delete(t *testing.T) {
	type serviceMocks struct {
		repo   *address_mocks.MockRepository
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
			name: "Dado um ID válido quando o método Delete é chamado então o endereço é deletado",
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
				repo:   address_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := address.NewService(sm.repo, sm.logger)

			err := s.Delete(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}
