package gin

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LucasMateus-eng/operations-service/driver"
	"github.com/gin-gonic/gin"
)

const (
	EMPTY_LIST_SIZE = 0
)

var (
	ErrEmptyDriverList = errors.New("no driver records found for the query parameters used")
)

type driverOutputDTO struct {
	ID            int                `json:"id"`
	UserID        int                `json:"user_id,omitempty"`
	Name          string             `json:"name,omitempty"`
	DateOfBirth   time.Time          `json:"date_of_birth,omitempty"`
	RG            string             `json:"rg,omitempty"`
	CPF           string             `json:"cpf,omitempty"`
	DriverLicense string             `json:"driver_license,omitempty"`
	CellPhone     string             `json:"cell_phone,omitempty"`
	Email         string             `json:"email,omitempty"`
	Address       *addressOutputDTO  `json:"address,omitempty"`
	Vehicles      []vehicleOutputDTO `json:"vehicles,omitempty"`
	CreatedAt     time.Time          `json:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty"`
	DeletedAt     time.Time          `json:"deleted_at,omitempty"`
}

type driverInputDTO struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	DateOfBirth   time.Time `json:"date_of_birth" binding:"required"`
	RG            string    `json:"rg" binding:"required"`
	CPF           string    `json:"cpf" binding:"required"`
	DriverLicense string    `json:"driver_license" binding:"required"`
	CellPhone     string    `json:"cell_phone" binding:"required"`
	Email         string    `json:"email" binding:"required"`
}

func mapDriverToOutputDTO(driver driver.Driver) *driverOutputDTO {
	return &driverOutputDTO{
		ID:            driver.ID,
		UserID:        driver.UserID,
		Name:          driver.Attributes.Name,
		DateOfBirth:   driver.Attributes.DateOfBirth,
		RG:            driver.LegalInformation.RG,
		CPF:           driver.LegalInformation.CPF,
		DriverLicense: driver.LegalInformation.DriverLicense,
		CellPhone:     driver.Contact.CellPhone,
		Email:         driver.Contact.Email,
		Address:       mapAddressToOutputDTO(*driver.Address),
		Vehicles:      mapVehicleListToOutputDTO(driver.Vehicles),
		CreatedAt:     driver.CreatedAt,
		UpdatedAt:     driver.UpdatedAt,
		DeletedAt:     driver.DeletedAt,
	}
}

func mapInputDTOToDriver(input driverInputDTO) *driver.Driver {
	return &driver.Driver{
		ID:     input.ID,
		UserID: input.UserID,
		Attributes: driver.DriverAttributes{
			Name:        input.Name,
			DateOfBirth: input.DateOfBirth,
		},
		LegalInformation: driver.DriverLegalInformation{
			RG:            input.RG,
			CPF:           input.CPF,
			DriverLicense: input.DriverLicense,
		},
		Contact: driver.Contact{
			CellPhone: input.CellPhone,
			Email:     input.Email,
		},
	}
}

type driverSpecificationInputDTO struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func mapInputDTOToDriverSpecification(input driverSpecificationInputDTO) *driver.DriverSpecification {
	return &driver.DriverSpecification{
		Page:     input.Page,
		PageSize: input.PageSize,
	}
}

func listDrivers(ctx context.Context, service *driver.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ds driverSpecificationInputDTO
		if err := c.ShouldBindQuery(&ds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		eagerLoadingHeader := c.GetHeader("eager-loading")
		if len(strings.TrimSpace(eagerLoadingHeader)) == 0 {
			eagerLoadingHeader = "false"
		}

		isEagerLoading, err := strconv.ParseBool(eagerLoadingHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		driverSpecification := mapInputDTOToDriverSpecification(ds)

		var drivers *[]driver.Driver
		if isEagerLoading {
			drivers, err = service.ListWithEagerLoading(ctx, driverSpecification)
		} else {
			drivers, err = service.List(ctx, driverSpecification)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(*drivers) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrEmptyDriverList.Error()})
			return
		}

		driversDTO := make([]driverOutputDTO, 0, len(*drivers))
		for _, dv := range *drivers {
			driversDTO = append(driversDTO, *mapDriverToOutputDTO(dv))
		}

		c.JSON(http.StatusOK, driversDTO)
	}
}

func createDriver(ctx context.Context, service *driver.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto driverInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driver := mapInputDTOToDriver(dto)

		driverID, err := service.Create(ctx, driver)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, driverOutputDTO{ID: driverID})
	}
}

func getDriver(ctx context.Context, service *driver.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driver, err := service.GetByDriverId(ctx, driverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := mapDriverToOutputDTO(*driver)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func updateDriver(ctx context.Context, service *driver.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dto driverInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driver := mapInputDTOToDriver(dto)
		driver.ID = driverID

		err = service.Update(ctx, driver)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func deleteDriver(ctx context.Context, service *driver.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.Delete(ctx, driverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
