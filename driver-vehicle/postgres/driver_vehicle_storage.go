package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LucasMateus-eng/operations-service/driver"
	driver_vehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	"github.com/LucasMateus-eng/operations-service/driver-vehicle/postgres/dto"
	"github.com/LucasMateus-eng/operations-service/driver-vehicle/postgres/mapping"
	driver_dto "github.com/LucasMateus-eng/operations-service/driver/postgres/dto"
	mapping_driver "github.com/LucasMateus-eng/operations-service/driver/postgres/mapping"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	mapping_vehicle "github.com/LucasMateus-eng/operations-service/vehicle/postgres/mapping"
	"github.com/uptrace/bun"
)

type driverVehiclePostgresRepo struct {
	db *bun.DB
}

func New(db *bun.DB) *driverVehiclePostgresRepo {
	return &driverVehiclePostgresRepo{
		db: db,
	}
}

func (dr *driverVehiclePostgresRepo) GetByID(ctx context.Context, driverID, vehicleID int) (*driver_vehicle.DriverVehicle, error) {
	var driverVehicleDTO dto.DriverVehicleDTO

	err := dr.db.NewSelect().
		Model(&driverVehicleDTO).
		Where("driver_id = ? AND vehicle_id = ?", driverID, vehicleID).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}

	mappedValue := mapping.MapDTOToDriverVehicle(&driverVehicleDTO)

	return mappedValue, nil
}

func (dr *driverVehiclePostgresRepo) GetDriverListByVehicleID(ctx context.Context, specification driver_vehicle.DriverVehicleSpectification) (*[]driver.Driver, error) {
	var driverVehicleDTOs []dto.DriverVehicleDTO

	query := dr.db.NewSelect().
		Model(&driverVehicleDTOs).
		Relation("Driver").
		Where("vehicle_id = ?", specification.VehicleID).
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
	for _, dto := range driverVehicleDTOs {
		mappedValue, err := mapping_driver.MapDTOToDriver(&dto.Driver)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, *mappedValue)
	}

	return &drivers, nil
}

func (dv *driverVehiclePostgresRepo) GetVehicleListByDriverID(ctx context.Context, specification driver_vehicle.DriverVehicleSpectification) (*[]vehicle.Vehicle, error) {
	var driverVehicleDTOs []dto.DriverVehicleDTO

	query := dv.db.NewSelect().
		Model(&driverVehicleDTOs).
		Relation("Vehicle").
		Where("driver_id = ?", specification.DriverID).
		Order("id ASC")

	if specification.Page > 0 && specification.PageSize > 0 {
		offset := (specification.Page - 1) * specification.PageSize
		query = query.Offset(offset).Limit(specification.PageSize)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	var vehicles []vehicle.Vehicle
	for _, dto := range driverVehicleDTOs {
		mappedValue, err := mapping_vehicle.MapDTOToVehicle(&dto.Vehicle)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, *mappedValue)
	}

	return &vehicles, nil
}

func (dr *driverVehiclePostgresRepo) Create(ctx context.Context, dv *driver_vehicle.DriverVehicle) (*driver_vehicle.DriverVehicle, error) {
	tx, err := dr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	driverExists, err := tx.NewSelect().Model((*driver_dto.DriverDTO)(nil)).Where("id = ?", dv.DriverID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	vehicleExists, err := tx.NewSelect().Model((*vehicle_dto.VehicleDTO)(nil)).Where("id = ?", dv.VehicleID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if driverExists && vehicleExists {
		driverVehicleDTO := mapping.MapDriverVehicleToDTO(dv)

		_, err := tx.NewInsert().Model(driverVehicleDTO).Exec(ctx)
		if err != nil {
			return nil, err
		}

		if err := tx.Commit(); err != nil {
			return nil, err
		}

		mappedValue := mapping.MapDTOToDriverVehicle(driverVehicleDTO)

		return mappedValue, nil
	}

	return nil, errors.New("driver or vehicle does not exist")
}

func (dr *driverVehiclePostgresRepo) Delete(ctx context.Context, driverID, vehicleID int) error {
	_, err := dr.db.NewDelete().Model((*dto.DriverVehicleDTO)(nil)).
		Where("driver_id = ? AND vehicle_id = ?", driverID, vehicleID).
		Exec(ctx)
	return err
}
