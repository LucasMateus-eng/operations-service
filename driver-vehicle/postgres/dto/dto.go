package dto

import (
	"time"

	driver_dto "github.com/LucasMateus-eng/operations-service/driver/postgres/dto"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	"github.com/uptrace/bun"
)

type DriverVehicleDTO struct {
	bun.BaseModel `bun:"table:drivers_vehicles"`

	DriverID  int64                  `bun:"driver_id,pk"`
	Driver    driver_dto.DriverDTO   `bun:"rel:belongs-to,join:driver_id=id"`
	VehicleID int64                  `bun:"vehicle_id,pk"`
	Vehicle   vehicle_dto.VehicleDTO `bun:"rel:belongs-to,join:vehicle_id=id"`
	CreatedAt time.Time              `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time              `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time              `bun:"deleted_at,soft_delete,nullzero,notnull,default:'0001-01-01 00:00:00+00'"`
}
