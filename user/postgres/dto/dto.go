package dto

import (
	"time"

	address_dto "github.com/LucasMateus-eng/operations-service/address/postgres/dto"
	"github.com/uptrace/bun"
)

type UserDTO struct {
	bun.BaseModel `bun:"table:users"`

	ID             int64                   `bun:"id,pk,autoincrement"`
	Username       string                  `bun:"username,unique,notnull"`
	HashedPassword string                  `bun:"hashed_password,notnull"`
	Role           string                  `bun:"role,notnull"`
	CreatedAt      time.Time               `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time               `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt      time.Time               `bun:"deleted_at,soft_delete,nullzero,notnull,default:'0001-01-01 00:00:00+00'"`
	AddressDTO     *address_dto.AddressDTO `bun:"rel:has-one,join:id=user_id"`
}
