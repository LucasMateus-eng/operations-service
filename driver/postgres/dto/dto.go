package dto

import (
	"time"

	user_dto "github.com/LucasMateus-eng/operations-service/user/postgres/dto"
	vehicle_dto "github.com/LucasMateus-eng/operations-service/vehicle/postgres/dto"
	"github.com/uptrace/bun"
)

type DriverDTO struct {
	bun.BaseModel `bun:"table:drivers"`

	ID            int64                    `bun:"id,pk,autoincrement"`
	Name          string                   `bun:"name,notnull"`
	RG            string                   `bun:"rg,notnull,unique"`
	CPF           string                   `bun:"cpf,notnull,unique"`
	DriverLicense string                   `bun:"driver_license,notnull,unique"`
	DateOfBirth   time.Time                `bun:"date_of_birth,notnull"`
	CellPhone     string                   `bun:"cell_phone,notnull"`
	Email         string                   `bun:"email,notnull"`
	UserID        int64                    `bun:"user_id,notnull,unique"`
	User          user_dto.UserDTO         `bun:"rel:belongs-to,join:user_id=id"`
	Vehicles      []vehicle_dto.VehicleDTO `bun:"m2m:drivers_vehicles,join:Driver=Vehicle"`
	CreatedAt     time.Time                `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time                `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time                `bun:"deleted_at,soft_delete,nullzero,notnull,default:'0001-01-01 00:00:00+00'"`
}
