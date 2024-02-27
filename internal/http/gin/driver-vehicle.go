package gin

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	drivervehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	"github.com/gin-gonic/gin"
)

type driverVehicleOutputDTO struct {
	DriverID  int       `json:"driver_id"`
	VehicleID int       `json:"vehicle_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type driverVehicleInputDTO struct {
	DriverID  int `json:"driver_id" binding:"required"`
	VehicleID int `json:"vehicle_id" binding:"required"`
}

type driverVehicleSpectificationInputDTO struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func mapDriverVehicleToOutputDTO(driverVehicle drivervehicle.DriverVehicle) *driverVehicleOutputDTO {
	return &driverVehicleOutputDTO{
		DriverID:  driverVehicle.DriverID,
		VehicleID: driverVehicle.VehicleID,
		CreatedAt: driverVehicle.CreatedAt,
		UpdatedAt: driverVehicle.UpdatedAt,
		DeletedAt: driverVehicle.DeletedAt,
	}
}

func mapInputDTOToDriverVehicle(input driverVehicleInputDTO) *drivervehicle.DriverVehicle {
	return &drivervehicle.DriverVehicle{
		DriverID:  input.DriverID,
		VehicleID: input.VehicleID,
	}
}

func listDriversByVehicleID(ctx context.Context, service *drivervehicle.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID, err := strconv.Atoi(c.Param("vehicle_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ds driverVehicleSpectificationInputDTO
		if err := c.ShouldBindQuery(&ds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driverVehicleSpecification := &drivervehicle.DriverVehicleSpecification{
			VehicleID: vehicleID,
			Page:      ds.Page,
			PageSize:  ds.PageSize,
		}

		drivers, err := service.GetDriverListByVehicleID(ctx, driverVehicleSpecification)
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

func listVehiclesByDriverID(ctx context.Context, service *drivervehicle.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverID, err := strconv.Atoi(c.Param("driver_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ds driverVehicleSpectificationInputDTO
		if err := c.ShouldBindQuery(&ds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driverVehicleSpecification := &drivervehicle.DriverVehicleSpecification{
			DriverID: driverID,
			Page:     ds.Page,
			PageSize: ds.PageSize,
		}

		vehicles, err := service.GetVehicleListByDriverID(ctx, driverVehicleSpecification)
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

func createDriverVehicle(ctx context.Context, service *drivervehicle.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto driverVehicleInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driverVehicle := mapInputDTOToDriverVehicle(dto)

		driverVehicle, err := service.Create(ctx, driverVehicle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := mapDriverVehicleToOutputDTO(*driverVehicle)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func deleteDriverVehicle(ctx context.Context, service *drivervehicle.Service, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverID, err := strconv.Atoi(c.Param("driver_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicleID, err := strconv.Atoi(c.Param("vehicle_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.Delete(ctx, driverID, vehicleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
