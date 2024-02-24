package vehicle

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type LicensingStatus int64

const (
	UNDEFINED LicensingStatus = iota
	REGULAR
	LATE
	BLOCKED
	SEIZED
	STOLEN
)

var (
	licensingStatusMap = map[string]LicensingStatus{
		"UNDEFINED": UNDEFINED,
		"REGULAR":   REGULAR,
		"LATE":      LATE,
		"BLOCKED":   BLOCKED,
		"SEIZED":    SEIZED,
		"STOLEN":    STOLEN,
	}

	licensingStatusList = []LicensingStatus{
		UNDEFINED,
		REGULAR,
		LATE,
		BLOCKED,
		SEIZED,
		STOLEN,
	}
)

func GetLicensingStatus(name string) (*LicensingStatus, error) {
	for key, value := range licensingStatusMap {
		if strings.EqualFold(key, name) {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("the given licensing status [%s] is non-existent in the map of valid values", name)
}

func (ls LicensingStatus) Change(new LicensingStatus) (*LicensingStatus, error) {
	if new == UNDEFINED {
		return nil, errors.New("the new licensing status cannot be equal to undefined")
	}

	if new == ls {
		return nil, errors.New("the new licensing status cannot be the same as the old one")
	}

	if !slices.Contains(licensingStatusList, new) {
		return nil, fmt.Errorf("the new licensing status [%v] is not present in the available list of valid values", new)
	}

	return &new, nil
}

func (ls LicensingStatus) String() string {
	switch ls {
	case REGULAR:
		return "REGULAR"
	case LATE:
		return "LATE"
	case BLOCKED:
		return "BLOCKED"
	case SEIZED:
		return "SEIZED"
	case STOLEN:
		return "STOLEN"
	}

	return "UNDEFINED"
}

type Licensing struct {
	ExpiryDate time.Time
	Status     LicensingStatus
}

type VehicleLegalInformation struct {
	Plate     string
	Renavam   string
	Licensing Licensing
}

type VehicleAttributes struct {
	Brand             string
	Model             string
	YearOfManufacture time.Time
}

type Vehicle struct {
	ID               int
	Attributes       VehicleAttributes
	LegalInformation VehicleLegalInformation
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type VehicleList []Vehicle

type VehicleSpectification struct {
	Attributes     VehicleAttributes
	Licensing      Licensing
	Page, PageSize int
}

type Reading interface {
	GetById(ctx context.Context, id int) (*Vehicle, error)
	GetByPlate(ctx context.Context, plate string) (*Vehicle, error)
	GetByRenavam(ctx context.Context, renavam string) (*Vehicle, error)
	List(ctx context.Context, specification VehicleSpectification) (*[]Vehicle, error)
}

type Writing interface {
	Create(ctx context.Context, v *Vehicle) (int, error)
	Update(ctx context.Context, v *Vehicle) error
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Reading
	Writing
}

type UseCase interface {
	GetByVehicleId(ctx context.Context, id int) (*Vehicle, error)
	GetByPlate(ctx context.Context, plate string) (*Vehicle, error)
	GetByRenavam(ctx context.Context, renavam string) (*Vehicle, error)
	List(ctx context.Context, specification VehicleSpectification) (*[]Vehicle, error)
	Create(ctx context.Context, v *Vehicle) (int, error)
	Update(ctx context.Context, v *Vehicle) error
	Delete(ctx context.Context, id int) error
}
