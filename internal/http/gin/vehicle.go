package gin

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	"github.com/gin-gonic/gin"
)

var (
	ErrEmptyVehicleList = errors.New("no vehicle records found for the query parameters used")
)

type vehicleOutputDTO struct {
	ID                  int                     `json:"id"`
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

type vehicleInputDTO struct {
	ID                  int                     `json:"id" binding:"required"`
	Brand               string                  `json:"brand" binding:"required"`
	Model               string                  `json:"model" binding:"required"`
	YearOfManufacture   time.Time               `json:"year_of_manufacture" binding:"required"`
	Plate               string                  `json:"plate" binding:"required"`
	Renavam             string                  `json:"renavam" binding:"required"`
	LicensingExpiryDate time.Time               `json:"licensing_expiry_date" binding:"required"`
	LicensingStatus     vehicle.LicensingStatus `json:"licensing_status" binding:"required"`
}

type vehicleSpecificationInputDTO struct {
	Brand               string                  `form:"brand"`
	Model               string                  `form:"model"`
	YearOfManufacture   time.Time               `form:"year_of_manufacture"`
	LicensingExpiryDate time.Time               `form:"licensing_expiry_date"`
	LicensingStatus     vehicle.LicensingStatus `form:"licensing_status"`
	Page                int                     `form:"page" binding:"required"`
	PageSize            int                     `form:"pageSize" binding:"required"`
}

func mapInputDTOToVehicleSpecification(input vehicleSpecificationInputDTO) *vehicle.VehicleSpectification {
	return &vehicle.VehicleSpectification{
		Attributes: vehicle.VehicleAttributes{
			Brand:             input.Brand,
			Model:             input.Model,
			YearOfManufacture: input.YearOfManufacture,
		},
		Licensing: vehicle.Licensing{
			ExpiryDate: input.LicensingExpiryDate,
			Status:     input.LicensingStatus,
		},
		Page:     input.Page,
		PageSize: input.PageSize,
	}
}

func mapVehicleToOutputDTO(vehicle vehicle.Vehicle) *vehicleOutputDTO {
	return &vehicleOutputDTO{
		ID:                  vehicle.ID,
		Brand:               vehicle.Attributes.Brand,
		Model:               vehicle.Attributes.Model,
		YearOfManufacture:   vehicle.Attributes.YearOfManufacture,
		Plate:               vehicle.LegalInformation.Plate,
		Renavam:             vehicle.LegalInformation.Renavam,
		LicensingExpiryDate: vehicle.LegalInformation.Licensing.ExpiryDate,
		LicensingStatus:     vehicle.LegalInformation.Licensing.Status,
		CreatedAt:           vehicle.CreatedAt,
		UpdatedAt:           vehicle.UpdatedAt,
		DeletedAt:           vehicle.DeletedAt,
	}
}

func mapInputDTOToVehicle(input vehicleInputDTO) *vehicle.Vehicle {
	return &vehicle.Vehicle{
		ID: input.ID,
		Attributes: vehicle.VehicleAttributes{
			Brand:             input.Brand,
			Model:             input.Model,
			YearOfManufacture: input.YearOfManufacture,
		},
		LegalInformation: vehicle.VehicleLegalInformation{
			Plate:   input.Plate,
			Renavam: input.Renavam,
			Licensing: vehicle.Licensing{
				ExpiryDate: input.LicensingExpiryDate,
				Status:     input.LicensingStatus,
			},
		},
	}
}

func mapVehicleListToOutputDTO(vehicles []vehicle.Vehicle) []vehicleOutputDTO {
	vehicleDTOs := make([]vehicleOutputDTO, 0, len(vehicles))
	for i, v := range vehicles {
		vehicleDTOs[i] = *mapVehicleToOutputDTO(v)
	}
	return vehicleDTOs
}

func listVehicles(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vs vehicleSpecificationInputDTO
		if err := c.ShouldBindQuery(&vs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicleSpecification := mapInputDTOToVehicleSpecification(vs)

		vehicles, err := service.List(ctx, vehicleSpecification)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(*vehicles) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrEmptyVehicleList.Error()})
			return
		}

		vehiclesDTO := make([]vehicleOutputDTO, 0, len(*vehicles))
		for _, v := range *vehicles {
			vehiclesDTO = append(vehiclesDTO, *mapVehicleToOutputDTO(v))
		}

		c.JSON(http.StatusOK, vehiclesDTO)
	}
}

func getVehicle(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicle, err := service.GetByVehicleId(ctx, vehicleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := mapVehicleToOutputDTO(*vehicle)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func createVehicle(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto vehicleInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicle := mapInputDTOToVehicle(dto)

		vehicleID, err := service.Create(ctx, vehicle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, vehicleOutputDTO{ID: vehicleID})
	}
}

func updateVehicle(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dto vehicleInputDTO
		if err = c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicle := mapInputDTOToVehicle(dto)
		vehicle.ID = vehicleID

		err = service.Update(ctx, vehicle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func deleteVehicle(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.Delete(ctx, vehicleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
