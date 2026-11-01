package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/service"
	"github.com/yohagos/multi-content-management/pkg/mapper"
)

type BuildingHandler struct {
	buildingService *service.BuildingService
}

func NewBuildingHandler(buildingService *service.BuildingService) *BuildingHandler {
	return &BuildingHandler{
		buildingService: buildingService,
	}
}

func (h *BuildingHandler) Create(c *gin.Context) {
	var building domain.Building

	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.buildingService.Create(c.Request.Context(), &building); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *BuildingHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	building, err := h.buildingService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrBuildingNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "building not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, building)
}

func (h *BuildingHandler) List(c *gin.Context) {
	var filter domain.BuildingFilter

	if tenantID := c.Query("tenant_id"); tenantID != "" {
		filter.TenantID = tenantID
	}
	if search := c.Query("search"); search != "" {
		filter.Search = search
	}
	if city := c.Query("city"); city != "" {
		filter.City = city
	}
	if status := c.Query("status"); status != "" {
		filter.Status = status
	}
	if published := c.Query("published"); published != "" {
		pubBool := published == "true"
		filter.Published = pubBool
	}
	if limit := c.Query("limit"); limit != "" {
		filter.Limit = parseInt(limit, 20)
	}
	if offset := c.Query("offset"); offset != "" {
		filter.Offset = parseInt(offset, 0)
	}

	buildings, total, err := h.buildingService.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   buildings,
		"total":  total,
		"limit":  filter.Limit,
		"offset": filter.Offset,
	})
}

func (h *BuildingHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var update domain.Building
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, err := h.buildingService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrBuildingNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "building not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mapper.MapBuilding(existing, &update)

	if err := h.buildingService.Update(c.Request.Context(), existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *BuildingHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.buildingService.Delete(c.Request.Context(), id); err != nil {
		if err == service.ErrBuildingNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "building not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
