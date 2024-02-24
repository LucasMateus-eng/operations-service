package dto

import (
	"time"

	"github.com/uptrace/bun"
)

type VehicleDTO struct {
	bun.BaseModel `bun:"table:vehicles"`

	ID                  int       `bun:"id,pk"`
	Brand               string    `bun:"brand,notnull"`
	Model               string    `bun:"model,notnull"`
	YearOfManufacture   time.Time `bun:"year_of_manufacture,notnull"`
	Plate               string    `bun:"plate,notnull,unique"`
	Renavam             string    `bun:"renavam,notnull,unique"`
	LicensingExpiryDate time.Time `bun:"licensing_expiry_date,notnull"`
	LicensingStatus     string    `bun:"licensing_status,notnull"`
	CreatedAt           time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt           time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt           time.Time `bun:"deleted_at,soft_delete,nullzero,notnull,default:'0001-01-01 00:00:00+00'"`
}
