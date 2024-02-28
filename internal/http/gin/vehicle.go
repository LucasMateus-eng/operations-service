package gin

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	gin_dto "github.com/LucasMateus-eng/operations-service/internal/http/gin/dto"
	gin_mapping "github.com/LucasMateus-eng/operations-service/internal/http/gin/mapping"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/LucasMateus-eng/operations-service/vehicle"
	"github.com/gin-gonic/gin"
)

var (
	ErrEmptyVehicleList = errors.New("no vehicle records found for the query parameters used")
)

func listVehicles(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vs gin_dto.VehicleSpecificationInputDTO
		if err := c.ShouldBindQuery(&vs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicleSpecification := gin_mapping.MapInputDTOToVehicleSpecification(vs)

		vehicles, err := service.List(ctx, vehicleSpecification)
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

func getVehicle(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicle, err := service.GetByVehicleId(ctx, vehicleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := gin_mapping.MapVehicleToOutputDTO(*vehicle)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func createVehicle(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto gin_dto.VehicleInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicle := gin_mapping.MapInputDTOToVehicle(dto)

		vehicleID, err := service.Create(ctx, vehicle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin_dto.VehicleOutputDTO{ID: vehicleID})
	}
}

func updateVehicle(ctx context.Context, service *vehicle.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dto gin_dto.VehicleInputDTO
		if err = c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vehicle := gin_mapping.MapInputDTOToVehicle(dto)
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
		vehicleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
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
