package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Role int64

const (
	UNDEFINED Role = iota
	ADMINISTRATOR
	EMPLOYEE
	DRIVER
)

var (
	roleMap = map[string]Role{
		"ADMINISTRATOR": ADMINISTRATOR,
		"EMPLOYEE":      EMPLOYEE,
		"DRIVER":        DRIVER,
	}

	roleList = []Role{
		ADMINISTRATOR,
		EMPLOYEE,
		DRIVER,
	}
)

func GetRole(name string) (Role, error) {
	for key, value := range roleMap {
		if strings.EqualFold(key, name) {
			return value, nil
		}
	}

	return UNDEFINED, fmt.Errorf("the given role [%s] is non-existent in the map of valid values", name)
}

func (r Role) Change(new Role) (*Role, error) {
	if new == UNDEFINED {
		return nil, errors.New("the new role cannot be equal to undefined")
	}

	if new == r {
		return nil, errors.New("the new role cannot be the same as the old one")
	}

	if !slices.Contains(roleList, new) {
		return nil, fmt.Errorf("the new role [%v] is not present in the available list of valid values", new)
	}

	return &new, nil
}

func (r Role) String() string {
	switch r {
	case ADMINISTRATOR:
		return "ADMINISTRATOR"
	case EMPLOYEE:
		return "EMPLOYEE"
	case DRIVER:
		return "DRIVER"
	}

	return "UNDEFINED"
}

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Role) UnmarshalJSON(data []byte) (err error) {
	var role string
	if err := json.Unmarshal(data, &role); err != nil {
		return err
	}
	if *r, err = GetRole(role); err != nil {
		return err
	}

	return nil
}

func (r Role) IsAdministrator() bool {
	return r == ADMINISTRATOR
}

func (r Role) IsEmployee() bool {
	return r == EMPLOYEE
}

func (r Role) IsDriver() bool {
	return r == DRIVER
}

type User struct {
	ID             int
	Username       string
	HashedPassword string
	Role           Role
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type Reading interface {
	GetById(ctx context.Context, id int) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByRole(ctx context.Context, role Role) (*User, error)
}

type Writing interface {
	Create(ctx context.Context, u *User) (int, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Reading
	Writing
}

type UseCase interface {
	GetById(ctx context.Context, id int) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByRole(ctx context.Context, role Role) (*User, error)
	Create(ctx context.Context, u *User) (int, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id int) error
}
