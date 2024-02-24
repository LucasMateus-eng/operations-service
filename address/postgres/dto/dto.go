package dto

import (
	"time"

	"github.com/uptrace/bun"
)

type AddressDTO struct {
	bun.BaseModel `bun:"table:adresses"`

	ID           int       `bun:"id,pk"`
	Locality     string    `bun:"locality,notnull"`
	Number       string    `bun:"number,notnull"`
	Complement   string    `bun:"complement"`
	Neighborhood string    `bun:"neighborhood,notnull"`
	City         string    `bun:"city,notnull"`
	State        string    `bun:"state,notnull"`
	CEP          string    `bun:"cep,notnull"`
	Country      string    `bun:"country,notnull"`
	UserID       int       `bun:"user_id,notnull,unique"`
	CreatedAt    time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt    time.Time `bun:"deleted_at,soft_delete,nullzero,notnull,default:'0001-01-01 00:00:00+00'"`
}
