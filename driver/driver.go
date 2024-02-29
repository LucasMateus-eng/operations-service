package driver

import (
	"context"
	"time"

	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/vehicle"
)

type DriverAttributes struct {
	Name        string
	DateOfBirth time.Time
}

type DriverLegalInformation struct {
	RG            string
	CPF           string
	DriverLicense string
}

type Contact struct {
	CellPhone string
	Email     string
}

type Driver struct {
	ID               int64
	UserID           int64
	Attributes       DriverAttributes
	LegalInformation DriverLegalInformation
	Address          *address.Address
	Contact          Contact
	Vehicles         []vehicle.Vehicle
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type DriverSpecification struct {
	Page, PageSize int
}

type Reading interface {
	GetByID(ctx context.Context, id int64) (*Driver, error)
	GetByUserID(ctx context.Context, userId int64) (*Driver, error)
	GetByIDWithEagerLoading(ctx context.Context, id int64) (*Driver, error)
	GetByUserIDWithEagerLoading(ctx context.Context, userId int64) (*Driver, error)
	List(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
	ListWithEagerLoading(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
}

type Writing interface {
	Create(ctx context.Context, d *Driver) (int64, error)
	Update(ctx context.Context, d *Driver) error
	Delete(ctx context.Context, id int64) error
}

type Repository interface {
	Reading
	Writing
}

type UseCase interface {
	GetByID(ctx context.Context, id int64) (*Driver, error)
	GetByUserID(ctx context.Context, userId int64) (*Driver, error)
	GetByIDWithEagerLoading(ctx context.Context, id int64) (*Driver, error)
	GetByUserIDWithEagerLoading(ctx context.Context, userId int64) (*Driver, error)
	List(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
	ListWithEagerLoading(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
	Create(ctx context.Context, d *Driver) (int64, error)
	Update(ctx context.Context, d *Driver) error
	Delete(ctx context.Context, id int64) error
}
