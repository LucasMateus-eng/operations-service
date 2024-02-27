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
	ID               int
	UserID           int
	Attributes       DriverAttributes
	LegalInformation DriverLegalInformation
	Address          *address.Address
	Contact          Contact
	Vehicles         vehicle.VehicleList
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type DriverList []Driver

type DriverSpecification struct {
	Page, PageSize int
}

type Reading interface {
	GetById(ctx context.Context, id int) (*Driver, error)
	GetByUserId(ctx context.Context, userId int) (*Driver, error)
	GetByIdWithEagerLoading(ctx context.Context, id int) (*Driver, error)
	GetByUserIdWithEagerLoading(ctx context.Context, userId int) (*Driver, error)
	List(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
	ListWithEagerLoading(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
}

type Writing interface {
	Create(ctx context.Context, d *Driver) (int, error)
	Update(ctx context.Context, d *Driver) error
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Reading
	Writing
}

type UseCase interface {
	GetByDriverId(ctx context.Context, id int) (*Driver, error)
	GetByUserId(ctx context.Context, userId int) (*Driver, error)
	GetByIdWithEagerLoading(ctx context.Context, id int) (*Driver, error)
	GetByUserIdWithEagerLoading(ctx context.Context, userId int) (*Driver, error)
	List(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
	ListWithEagerLoading(ctx context.Context, specification *DriverSpecification) (*[]Driver, error)
	Create(ctx context.Context, d *Driver) (int, error)
	Update(ctx context.Context, d *Driver) error
	Delete(ctx context.Context, id int) error
}
