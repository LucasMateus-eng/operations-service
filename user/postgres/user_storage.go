package postgres

import (
	"context"

	"github.com/LucasMateus-eng/operations-service/user"
	"github.com/LucasMateus-eng/operations-service/user/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/user/postgres/mapping"

	"github.com/uptrace/bun"
)

type userPostgresRepo struct {
	db *bun.DB
}

func New(db *bun.DB) *userPostgresRepo {
	return &userPostgresRepo{
		db: db,
	}
}

func (ur *userPostgresRepo) GetByID(ctx context.Context, id int64) (*user.User, error) {
	var userDTO dto.UserDTO

	err := ur.db.NewSelect().Model(&userDTO).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToUser(&userDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (ur *userPostgresRepo) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var userDTO dto.UserDTO

	err := ur.db.NewSelect().Model(&userDTO).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToUser(&userDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (ur *userPostgresRepo) GetByRole(ctx context.Context, role user.Role) (*user.User, error) {
	var userDTO dto.UserDTO

	err := ur.db.NewSelect().Model(&userDTO).Where("role = ?", role.String()).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToUser(&userDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (ur *userPostgresRepo) Create(ctx context.Context, u *user.User) (int64, error) {
	var userID int64

	userDTO := mapping.MapUserToDTO(u)

	query := ur.db.NewInsert().Model(userDTO).On("CONFLICT (id) DO UPDATE").Returning("id")

	err := query.Scan(ctx, &userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ur *userPostgresRepo) Update(ctx context.Context, u *user.User) error {
	userDTO := mapping.MapUserToDTO(u)

	_, err := ur.db.NewUpdate().Model(userDTO).
		OmitZero().
		ExcludeColumn("plate", "renavam", "deleted_at").
		Where("id = ?", userDTO.ID).
		Exec(ctx)

	return err
}

func (ur *userPostgresRepo) Delete(ctx context.Context, id int64) error {
	_, err := ur.db.NewDelete().Model((*dto.UserDTO)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
