package postgres

import (
	"context"

	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/address/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/address/postgres/mapping"
	"github.com/uptrace/bun"
)

type addressPostgresRepo struct {
	db *bun.DB
}

func New(db *bun.DB) *addressPostgresRepo {
	return &addressPostgresRepo{
		db: db,
	}
}

func (ar *addressPostgresRepo) GetByID(ctx context.Context, id int64) (*address.Address, error) {
	var addressDTO dto.AddressDTO

	err := ar.db.NewSelect().Model(&addressDTO).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToAddress(&addressDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (ar *addressPostgresRepo) GetByUserID(ctx context.Context, userID int64) (*address.Address, error) {
	var addressDTO dto.AddressDTO

	err := ar.db.NewSelect().Model(&addressDTO).Where("user_id = ?", userID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToAddress(&addressDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (ar *addressPostgresRepo) Create(ctx context.Context, a *address.Address) (int64, error) {
	var addressID int64

	addressDTO := mapping.MapAddressToDTO(a)

	query := ar.db.NewInsert().Model(addressDTO).On("CONFLICT (id) DO UPDATE").Returning("id")

	err := query.Scan(ctx, &addressID)
	if err != nil {
		return 0, err
	}

	return addressID, nil
}

func (ar *addressPostgresRepo) Update(ctx context.Context, a *address.Address) error {
	addressDTO := mapping.MapAddressToDTO(a)

	_, err := ar.db.NewUpdate().Model(addressDTO).
		OmitZero().
		ExcludeColumn("deleted_at").
		Where("id = ?", addressDTO.ID).
		Exec(ctx)

	return err
}

func (ar *addressPostgresRepo) Delete(ctx context.Context, id int64) error {
	_, err := ar.db.NewDelete().Model((*dto.AddressDTO)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
