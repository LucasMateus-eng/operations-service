package postgres

import (
	"context"
	"strings"

	"github.com/LucasMateus-eng/operations-service/vehicle"
	"github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/vehicle/postgres/mapping"
	"github.com/uptrace/bun"
)

type vehiclePostgresRepo struct {
	db *bun.DB
}

func New(db *bun.DB) *vehiclePostgresRepo {
	return &vehiclePostgresRepo{
		db: db,
	}
}

func (vr *vehiclePostgresRepo) GetByID(ctx context.Context, id int64) (*vehicle.Vehicle, error) {
	var vehicleDTO dto.VehicleDTO

	err := vr.db.NewSelect().Model(&vehicleDTO).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToVehicle(&vehicleDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (vr *vehiclePostgresRepo) GetByPlate(ctx context.Context, plate string) (*vehicle.Vehicle, error) {
	var vehicleDTO dto.VehicleDTO

	err := vr.db.NewSelect().Model(&vehicleDTO).Where("plate = ?", plate).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToVehicle(&vehicleDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (vr *vehiclePostgresRepo) GetByRenavam(ctx context.Context, renavam string) (*vehicle.Vehicle, error) {
	var vehicleDTO dto.VehicleDTO

	err := vr.db.NewSelect().Model(&vehicleDTO).Where("renavam = ?", renavam).Scan(ctx)
	if err != nil {
		return nil, err
	}

	mappedValue, err := mapping.MapDTOToVehicle(&vehicleDTO)
	if err != nil {
		return nil, err
	}

	return mappedValue, nil
}

func (vr *vehiclePostgresRepo) List(ctx context.Context, specification *vehicle.VehicleSpectification) (*[]vehicle.Vehicle, error) {
	var vehicleDTOs []dto.VehicleDTO

	query := vr.db.NewSelect().Model(&vehicleDTOs).Order("id ASC")

	if specification.Page > 0 && specification.PageSize > 0 {
		offset := (specification.Page - 1) * specification.PageSize
		query = query.Offset(offset).Limit(specification.PageSize)
	}

	if len(strings.TrimSpace(specification.Attributes.Brand)) == 0 {
		query = query.Where("brand = ?", specification.Attributes.Brand)
	}

	if len(strings.TrimSpace(specification.Attributes.Model)) == 0 {
		query = query.Where("model = ?", specification.Attributes.Model)
	}

	if len(strings.TrimSpace(specification.Attributes.YearOfManufacture.Format("2006/01/02"))) == 0 {
		query = query.Where("year_of_manufacture = ?", specification.Attributes.YearOfManufacture.Format("2006/01/02"))
	}

	if len(strings.TrimSpace(specification.Licensing.ExpiryDate.Format("2006-01-02 -07:00:00"))) == 0 {
		query = query.Where("licensing_expiry_date = ?", specification.Licensing.ExpiryDate.Format("2006-01-02 -07:00:00"))
	}

	if len(strings.TrimSpace(specification.Licensing.Status.String())) == 0 {
		query = query.Where("licensing_status = ?", specification.Licensing.Status.String())
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	var vehicles []vehicle.Vehicle
	for _, dto := range vehicleDTOs {
		mappedValue, err := mapping.MapDTOToVehicle(&dto)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, *mappedValue)
	}

	return &vehicles, nil
}

func (vr *vehiclePostgresRepo) Create(ctx context.Context, v *vehicle.Vehicle) (int64, error) {
	var vehicleID int64

	vehicleDTO := mapping.MapVehicleToDTO(v)

	query := vr.db.NewInsert().Model(vehicleDTO).On("CONFLICT (id) DO UPDATE").Returning("id")

	err := query.Scan(ctx, &vehicleID)
	if err != nil {
		return 0, err
	}

	return vehicleID, nil
}

func (vr *vehiclePostgresRepo) Update(ctx context.Context, v *vehicle.Vehicle) error {
	vehicleDTO := mapping.MapVehicleToDTO(v)

	_, err := vr.db.NewUpdate().Model(vehicleDTO).
		OmitZero().
		ExcludeColumn("plate", "renavam", "deleted_at").
		Where("id = ?", vehicleDTO.ID).
		Exec(ctx)

	return err
}

func (vr *vehiclePostgresRepo) Delete(ctx context.Context, id int64) error {
	_, err := vr.db.NewDelete().Model((*dto.VehicleDTO)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
