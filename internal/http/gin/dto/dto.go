package dto

import (
	"time"

	"github.com/LucasMateus-eng/operations-service/address"
	"github.com/LucasMateus-eng/operations-service/user"
	"github.com/LucasMateus-eng/operations-service/vehicle"
)

type AddressOutputDTO struct {
	ID           int64                  `json:"id"`
	UserID       int64                  `json:"user_id,omitempty"`
	Locality     string                 `json:"locality,omitempty"`
	Number       string                 `json:"number,omitempty"`
	Complement   string                 `json:"complement,omitempty"`
	Neighborhood string                 `json:"neighborhood,omitempty"`
	City         string                 `json:"city,omitempty"`
	State        address.BrazilianState `json:"state,omitempty"`
	CEP          string                 `json:"cep,omitempty"`
	Country      string                 `json:"country,omitempty"`
	CreatedAt    time.Time              `json:"created_at,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at,omitempty"`
	DeletedAt    time.Time              `json:"deleted_at,omitempty"`
}

type DriverVehicleOutputDTO struct {
	DriverID  int64     `json:"driver_id"`
	VehicleID int64     `json:"vehicle_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type DriverVehicleInputDTO struct {
	DriverID  int64 `json:"driver_id" binding:"required"`
	VehicleID int64 `json:"vehicle_id" binding:"required"`
}

type DriverVehicleSpectificationInputDTO struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

type DriverOutputDTO struct {
	ID            int64              `json:"id"`
	UserID        int64              `json:"user_id,omitempty"`
	Name          string             `json:"name,omitempty"`
	DateOfBirth   time.Time          `json:"date_of_birth,omitempty"`
	RG            string             `json:"rg,omitempty"`
	CPF           string             `json:"cpf,omitempty"`
	DriverLicense string             `json:"driver_license,omitempty"`
	CellPhone     string             `json:"cell_phone,omitempty"`
	Email         string             `json:"email,omitempty"`
	Address       *AddressOutputDTO  `json:"address,omitempty"`
	Vehicles      []VehicleOutputDTO `json:"vehicles,omitempty"`
	CreatedAt     time.Time          `json:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty"`
	DeletedAt     time.Time          `json:"deleted_at,omitempty"`
}

type DriverInputDTO struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name" binding:"required"`
	DateOfBirth   time.Time `json:"date_of_birth" binding:"required"`
	RG            string    `json:"rg" binding:"required"`
	CPF           string    `json:"cpf" binding:"required"`
	DriverLicense string    `json:"driver_license" binding:"required"`
	CellPhone     string    `json:"cell_phone" binding:"required"`
	Email         string    `json:"email" binding:"required"`
}

type DriverSpecificationInputDTO struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

type UserOutputDTO struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
	Role           user.Role `json:"role,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
}

type UserInputDTO struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username" binding:"required"`
	HashedPassword string    `json:"hashed_password" binding:"required"`
	Role           user.Role `json:"role" binding:"required"`
}

type VehicleOutputDTO struct {
	ID                  int64                   `json:"id"`
	Brand               string                  `json:"brand,omitempty"`
	Model               string                  `json:"model,omitempty"`
	YearOfManufacture   time.Time               `json:"year_of_manufacture,omitempty"`
	Plate               string                  `json:"plate,omitempty"`
	Renavam             string                  `json:"renavam,omitempty"`
	LicensingExpiryDate time.Time               `json:"licensing_expiry_date,omitempty"`
	LicensingStatus     vehicle.LicensingStatus `json:"licensing_status,omitempty"`
	CreatedAt           time.Time               `json:"created_at,omitempty"`
	UpdatedAt           time.Time               `json:"updated_at,omitempty"`
	DeletedAt           time.Time               `json:"deleted_at,omitempty"`
}

type VehicleInputDTO struct {
	ID                  int64                   `json:"id"`
	Brand               string                  `json:"brand" binding:"required"`
	Model               string                  `json:"model" binding:"required"`
	YearOfManufacture   time.Time               `json:"year_of_manufacture" binding:"required"`
	Plate               string                  `json:"plate" binding:"required"`
	Renavam             string                  `json:"renavam" binding:"required"`
	LicensingExpiryDate time.Time               `json:"licensing_expiry_date" binding:"required"`
	LicensingStatus     vehicle.LicensingStatus `json:"licensing_status" binding:"required"`
}

type VehicleSpecificationInputDTO struct {
	Brand               string                  `form:"brand"`
	Model               string                  `form:"model"`
	YearOfManufacture   time.Time               `form:"year_of_manufacture"`
	LicensingExpiryDate time.Time               `form:"licensing_expiry_date"`
	LicensingStatus     vehicle.LicensingStatus `form:"licensing_status"`
	Page                int                     `form:"page" binding:"required"`
	PageSize            int                     `form:"pageSize" binding:"required"`
}
