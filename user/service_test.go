package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/LucasMateus-eng/operations-service/config"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	user_mocks "github.com/LucasMateus-eng/operations-service/internal/mocks/user"
	"github.com/LucasMateus-eng/operations-service/user"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

var (
	errMocked     = errors.New("some error")
	mockedContext = context.Background()
	expectedUser  = &user.User{
		ID: 1,
	}
)

func TestService_GetByID(t *testing.T) {
	type serviceMocks struct {
		repo   *user_mocks.MockRepository
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
		want        *user.User
		wantErr     bool
	}{
		{
			name: "Dado um ID válido quando o método GetByID é chamado então o usuário é retornado",
			args: args{
				ctx: mockedContext,
				id:  1,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByID(p.ctx, p.id).Return(expectedUser, nil)
			},
			want:    expectedUser,
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
				repo:   user_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := user.NewService(sm.repo, sm.logger)

			actualUser, err := s.GetByID(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualUser)
		})
	}
}

func TestService_GetByUsername(t *testing.T) {
	type serviceMocks struct {
		repo   *user_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx      context.Context
		username string
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *user.User
		wantErr     bool
	}{
		{
			name: "Dado um username válido quando o método GetByUsername é chamado então o usuário é retornado",
			args: args{
				ctx:      mockedContext,
				username: "user123",
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUsername(p.ctx, p.username).Return(expectedUser, nil)
			},
			want:    expectedUser,
			wantErr: false,
		},
		{
			name: "Dado um username inválido quando o método GetByUsername é chamado então um erro é retornado",
			args: args{
				ctx:      mockedContext,
				username: "userInexistente",
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByUsername(p.ctx, p.username).Return(nil, errMocked)
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
				repo:   user_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := user.NewService(sm.repo, sm.logger)

			actualUser, err := s.GetByUsername(test.args.ctx, test.args.username)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualUser)
		})
	}
}

func TestService_GetByRole(t *testing.T) {
	type serviceMocks struct {
		repo   *user_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx  context.Context
		role user.Role
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        *user.User
		wantErr     bool
	}{
		{
			name: "Dado um role válido quando o método GetByRole é chamado então o usuário é retornado",
			args: args{
				ctx:  mockedContext,
				role: user.ADMINISTRATOR,
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByRole(p.ctx, p.role).Return(expectedUser, nil)
			},
			want:    expectedUser,
			wantErr: false,
		},
		{
			name: "Dado um role inválido quando o método GetByRole é chamado então um erro é retornado",
			args: args{
				ctx:  mockedContext,
				role: user.Role(99),
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().GetByRole(p.ctx, p.role).Return(nil, errMocked)
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
				repo:   user_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := user.NewService(sm.repo, sm.logger)

			actualUser, err := s.GetByRole(test.args.ctx, test.args.role)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualUser)
		})
	}
}

func TestService_Create(t *testing.T) {
	type serviceMocks struct {
		repo   *user_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		u   *user.User
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		want        int64
		wantErr     bool
	}{
		{
			name: "Dado um usuário válido quando o método Create é chamado então o ID do usuário é retornado",
			args: args{
				ctx: mockedContext,
				u: &user.User{
					ID: 1,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.u).Return(int64(1), nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Dado um usuário inválido quando o método Create é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				u: &user.User{
					ID: 0,
				},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Create(p.ctx, p.u).Return(int64(0), errMocked)
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
				repo:   user_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := user.NewService(sm.repo, sm.logger)

			actualUserID, err := s.Create(test.args.ctx, test.args.u)

			assert.Equal(tt, test.wantErr, err != nil)
			assert.Equal(tt, test.want, actualUserID)
		})
	}
}

func TestService_Update(t *testing.T) {
	type serviceMocks struct {
		repo   *user_mocks.MockRepository
		logger *logging.Logging
	}

	type args struct {
		ctx context.Context
		u   *user.User
	}

	tests := []struct {
		name        string
		args        args
		prepareMock func(p args, m serviceMocks)
		wantErr     bool
	}{
		{
			name: "Dado um usuário válido quando o método Update é chamado então o usuário é atualizado",
			args: args{
				ctx: mockedContext,
				u:   &user.User{ID: 1, Username: "newUsername"},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.u).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Dado um usuário inválido quando o método Update é chamado então um erro é retornado",
			args: args{
				ctx: mockedContext,
				u:   &user.User{ID: 0},
			},
			prepareMock: func(p args, m serviceMocks) {
				m.repo.EXPECT().Update(p.ctx, p.u).Return(errMocked)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			sm := serviceMocks{
				repo:   user_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := user.NewService(sm.repo, sm.logger)

			err := s.Update(test.args.ctx, test.args.u)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}

func TestService_Delete(t *testing.T) {
	type serviceMocks struct {
		repo   *user_mocks.MockRepository
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
			name: "Dado um ID válido quando o método Delete é chamado então o usuário é excluído",
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
				repo:   user_mocks.NewMockRepository(ctrl),
				logger: logging.InitializerLogging(&config.Config{}),
			}

			if test.prepareMock != nil {
				test.prepareMock(test.args, sm)
			}

			s := user.NewService(sm.repo, sm.logger)

			err := s.Delete(test.args.ctx, test.args.id)

			assert.Equal(tt, test.wantErr, err != nil)
		})
	}
}
