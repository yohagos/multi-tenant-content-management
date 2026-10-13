package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/service"
	"github.com/yohagos/multi-content-management/pkg/mapper"
)

type UnitHandler struct {
	unitService *service.UnitService
}

func NewUnitHandler(unitService *service.UnitService) *UnitHandler {
	return &UnitHandler{
		unitService: unitService,
	}
}

func (h *UnitHandler) Create(c *gin.Context) {
	var req domain.Unit
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.unitService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (h *UnitHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	unit, err := h.unitService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrUnitNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "unit not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, unit)
}

func (h *UnitHandler) List(c *gin.Context) {
	var filter domain.UnitFilter

	if limit := c.Query("limit"); limit != "" {
		filter.Limit = parseInt(limit, 20)
	}
	if offset := c.Query("offset"); offset != "" {
		filter.Offset = parseInt(offset, 0)
	}
	filter.Search = c.Query("search")
	if published := c.Query("published"); published != "" {
		publishedBool := published == "true"
		filter.Published = &publishedBool
	}
	if buildingID := c.Query("building_id"); buildingID != "" {
		filter.BuildingID = buildingID
	}
	if status := c.Query("status"); status != "" {
		filter.Status = status
	}
	if minRent := c.Query("min_rent"); minRent != "" {
		minRentFloat, _ := strconv.ParseFloat(strings.TrimSpace(minRent), 64)
		filter.MinRent = &minRentFloat
	}
	if maxRent := c.Query("max_rent"); maxRent != "" {
		maxRentFloat, _ := strconv.ParseFloat(strings.TrimSpace(maxRent), 64)
		filter.MaxRent = &maxRentFloat
	}
	if minRooms := c.Query("min_rooms"); minRooms != "" {
		minRoomsFloat, _ := strconv.ParseFloat(strings.TrimSpace(minRooms), 64)
		filter.MinRooms = &minRoomsFloat
	}
	if maxRooms := c.Query("max_rooms"); maxRooms != "" {
		maxRoomsFloat, _ := strconv.ParseFloat(strings.TrimSpace(maxRooms), 64)
		filter.MaxRooms = &maxRoomsFloat
	}

	units, total, err := h.unitService.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   units,
		"total":  total,
		"limit":  filter.Limit,
		"offset": filter.Offset,
	})
}

func (h *UnitHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req domain.Unit
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, err := h.unitService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrUnitNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "unit not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mapper.MapUnit(existing, &req)

	if err := h.unitService.Update(c.Request.Context(), existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func (h *UnitHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.unitService.Delete(c.Request.Context(), id); err != nil {
		if err == service.ErrUnitNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "unit not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
