package gin

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/LucasMateus-eng/operations-service/driver"
	gin_dto "github.com/LucasMateus-eng/operations-service/internal/http/gin/dto"
	gin_mapping "github.com/LucasMateus-eng/operations-service/internal/http/gin/mapping"
	"github.com/LucasMateus-eng/operations-service/internal/logging"
	"github.com/gin-gonic/gin"
)

const (
	EMPTY_LIST_SIZE = 0
)

var (
	ErrEmptyDriverList = errors.New("no driver records found for the query parameters used")
)

func listDrivers(ctx context.Context, service *driver.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("List drivers", nil)

		var ds gin_dto.DriverSpecificationInputDTO
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

		driverSpecification := gin_mapping.MapInputDTOToDriverSpecification(ds)

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

		driversDTO := make([]gin_dto.DriverOutputDTO, 0, len(*drivers))
		for _, dv := range *drivers {
			driversDTO = append(driversDTO, *gin_mapping.MapDriverToOutputDTO(dv))
		}

		c.JSON(http.StatusOK, driversDTO)
	}
}

func createDriver(ctx context.Context, service *driver.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Create driver", nil)

		var dto gin_dto.DriverInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driver := gin_mapping.MapInputDTOToDriver(dto)

		driverID, err := service.Create(ctx, driver)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin_dto.DriverOutputDTO{ID: driverID})
	}
}

func getDriver(ctx context.Context, service *driver.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Get driver", nil)

		driverID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driver, err := service.GetByID(ctx, driverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputDTO := gin_mapping.MapDriverToOutputDTO(*driver)

		c.JSON(http.StatusOK, outputDTO)
	}
}

func updateDriver(ctx context.Context, service *driver.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Update driver", nil)

		driverID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dto gin_dto.DriverInputDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		driver := gin_mapping.MapInputDTOToDriver(dto)
		driver.ID = driverID

		err = service.Update(ctx, driver)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func deleteDriver(ctx context.Context, service *driver.Service, logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Delete driver", nil)

		driverID, err := strconv.ParseInt(c.Param("id"), 10, 64)
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
