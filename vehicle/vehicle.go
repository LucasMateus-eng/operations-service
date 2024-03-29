package vehicle

import (
	"context"
	"encoding/json"
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
		"REGULAR": REGULAR,
		"LATE":    LATE,
		"BLOCKED": BLOCKED,
		"SEIZED":  SEIZED,
		"STOLEN":  STOLEN,
	}

	licensingStatusList = []LicensingStatus{
		REGULAR,
		LATE,
		BLOCKED,
		SEIZED,
		STOLEN,
	}
)

func GetLicensingStatus(name string) (LicensingStatus, error) {
	for key, value := range licensingStatusMap {
		if strings.EqualFold(key, name) {
			return value, nil
		}
	}

	return UNDEFINED, fmt.Errorf("the given licensing status [%s] is non-existent in the map of valid values", name)
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

func (ls LicensingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(ls.String())
}

func (ls *LicensingStatus) UnmarshalJSON(data []byte) (err error) {
	var licensingStatus string
	if err := json.Unmarshal(data, &licensingStatus); err != nil {
		return err
	}
	if *ls, err = GetLicensingStatus(licensingStatus); err != nil {
		return err
	}

	return nil
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
	ID               int64
	Attributes       VehicleAttributes
	LegalInformation VehicleLegalInformation
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type VehicleSpectification struct {
	Attributes     VehicleAttributes
	Licensing      Licensing
	Page, PageSize int
}

type Reading interface {
	GetByID(ctx context.Context, id int64) (*Vehicle, error)
	GetByPlate(ctx context.Context, plate string) (*Vehicle, error)
	GetByRenavam(ctx context.Context, renavam string) (*Vehicle, error)
	List(ctx context.Context, specification *VehicleSpectification) (*[]Vehicle, error)
}

type Writing interface {
	Create(ctx context.Context, v *Vehicle) (int64, error)
	Update(ctx context.Context, v *Vehicle) error
	Delete(ctx context.Context, id int64) error
}

type Repository interface {
	Reading
	Writing
}

type UseCase interface {
	GetByID(ctx context.Context, id int64) (*Vehicle, error)
	GetByPlate(ctx context.Context, plate string) (*Vehicle, error)
	GetByRenavam(ctx context.Context, renavam string) (*Vehicle, error)
	List(ctx context.Context, specification *VehicleSpectification) (*[]Vehicle, error)
	Create(ctx context.Context, v *Vehicle) (int64, error)
	Update(ctx context.Context, v *Vehicle) error
	Delete(ctx context.Context, id int64) error
}
