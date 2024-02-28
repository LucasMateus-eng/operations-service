package gin

import (
	"context"
	"net/http"
	"strconv"

	drivervehicle "github.com/LucasMateus-eng/operations-service/driver-vehicle"
	gin_dto "github.com/LucasMateus-eng/operations-service/internal/http/gin/dto"
	gin_mapping "github.com/LucasMateus-eng/operations-service/internal/http/gin/mapping"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/gin-gonic/gin"
)

func listDriversByVehicleID(ctx context.Context, service *drivervehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ds gin_dto.DriverVehicleSpectificationInputDTO
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

		driversDTO := make([]gin_dto.DriverOutputDTO, 0, len(*drivers))
		for _, dv := range *drivers {
			driversDTO = append(driversDTO, *gin_mapping.MapDriverToOutputDTO(dv))
		}

		c.JSON(http.StatusOK, driversDTO)
	}
}

func listVehiclesByDriverID(ctx context.Context, service *drivervehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ds gin_dto.DriverVehicleSpectificationInputDTO
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

		vehiclesDTO := make([]gin_dto.VehicleOutputDTO, 0, len(*vehicles))
		for _, v := range *vehicles {
			vehiclesDTO = append(vehiclesDTO, *gin_mapping.MapVehicleToOutputDTO(v))
		}

		c.JSON(http.StatusOK, vehiclesDTO)
	}
}

func createDriverVehicle(ctx context.Context, service *drivervehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto gin_dto.DriverVehicleInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driverVehicle := gin_mapping.MapInputDTOToDriverVehicle(dto)

		driverVehicle, err := service.Create(ctx, driverVehicle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := gin_mapping.MapDriverVehicleToOutputDTO(*driverVehicle)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func deleteDriverVehicle(ctx context.Context, service *drivervehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		driverID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
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
