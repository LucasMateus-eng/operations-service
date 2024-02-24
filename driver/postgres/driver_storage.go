package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LucasMateus-eng/operations-service/driver"
	"github.com/LucasMateus-eng/operations-service/driver/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/driver/postgres/mapping"
	"github.com/uptrace/bun"
)

type driverPostgresRepo struct {
	db *bun.DB
}

func New(db *bun.DB) *driverPostgresRepo {
	return &driverPostgresRepo{
		db: db,
	}
}

func (dr *driverPostgresRepo) GetById(ctx context.Context, id int) (*driver.Driver, error) {
	var driverDTO dto.DriverDTO

	err := dr.db.NewSelect().Model(&driverDTO).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToDriver(&driverDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (dr *driverPostgresRepo) GetByUserId(ctx context.Context, userId int) (*driver.Driver, error) {
	var driverDTO dto.DriverDTO

	err := dr.db.NewSelect().Model(&driverDTO).Where("user_id = ?", userId).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToDriver(&driverDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (dr *driverPostgresRepo) GetByIdWithEagerLoading(ctx context.Context, id int) (*driver.Driver, error) {
	var driverDTO dto.DriverDTO

	err := dr.db.NewSelect().Model(&driverDTO).
		Relation("Vehicles").
		Relation("User").
		Relation("User.Address").
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToDriver(&driverDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (dr *driverPostgresRepo) GetByUserIdWithEagerLoading(ctx context.Context, userId int) (*driver.Driver, error) {
	var driverDTO dto.DriverDTO

	err := dr.db.NewSelect().Model(&driverDTO).
		Relation("Vehicles").
		Relation("User").
		Relation("User.Address").
		Where("user_id = ?", userId).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToDriver(&driverDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (dr *driverPostgresRepo) List(ctx context.Context, specification driver.DriverSpectification) (*[]driver.Driver, error) {
	var driverDTOs []dto.DriverDTO

	query := dr.db.NewSelect().Model(&driverDTOs).Order("id ASC")

	if specification.Page > 0 && specification.PageSize > 0 {
		offset := (specification.Page - 1) * specification.PageSize
		query = query.Offset(offset).Limit(specification.PageSize)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	var drivers []driver.Driver
	for _, dto := range driverDTOs {
		mappedValue, err := mapping.MapDTOToDriver(&dto)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, *mappedValue)
	}

	return &drivers, nil
}

func (dr *driverPostgresRepo) ListWithEagerLoading(ctx context.Context, specification driver.DriverSpectification) (*[]driver.Driver, error) {
	var driverDTOs []dto.DriverDTO

	query := dr.db.NewSelect().
		Model(&driverDTOs).
		Relation("Vehicles").
		Relation("User").
		Relation("User.Address").
		Order("id ASC")

	if specification.Page > 0 && specification.PageSize > 0 {
		offset := (specification.Page - 1) * specification.PageSize
		query = query.Offset(offset).Limit(specification.PageSize)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	var drivers []driver.Driver
	for _, dto := range driverDTOs {
		mappedValue, err := mapping.MapDTOToDriver(&dto)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, *mappedValue)
	}

	return &drivers, nil
}

func (dr *driverPostgresRepo) Create(ctx context.Context, d *driver.Driver) (int, error) {
	var driverID int

	driverDTO := mapping.MapDriverToDTO(d)

	query := dr.db.NewInsert().Model(driverDTO).On("CONFLICT (id) DO UPDATE").Returning("id")

	err := query.Scan(ctx, &driverID)
	if err != nil {
		return 0, err
	}

	return driverID, nil
}

func (dr *driverPostgresRepo) Update(ctx context.Context, d *driver.Driver) error {
	driverDTO := mapping.MapDriverToDTO(d)

	_, err := dr.db.NewUpdate().Model(driverDTO).
		OmitZero().
		ExcludeColumn("rg", "cpf", "driver_license", "deleted_at").
		Where("id = ?", driverDTO.ID).
		Exec(ctx)

	return err
}

func (dr *driverPostgresRepo) Delete(ctx context.Context, id int) error {
	_, err := dr.db.NewDelete().Model((*dto.DriverDTO)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
