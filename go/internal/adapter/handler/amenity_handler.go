package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/service"
	"github.com/yohagos/multi-content-management/pkg/mapper"
)

type AmenityHandler struct {
	amenityService *service.AmenityService
}

func NewAmenityHandler(amenityService *service.AmenityService) *AmenityHandler {
	return &AmenityHandler{
		amenityService: amenityService,
	}
}

func (h *AmenityHandler) Create(c *gin.Context) {
	var amenity domain.Amenity

	if err := c.ShouldBindJSON(&amenity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.amenityService.Create(c.Request.Context(), &amenity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, amenity)
}

func (h *AmenityHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	amenity, err := h.amenityService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrAmenityNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "tenant not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, amenity)
}

func (h *AmenityHandler) List(c *gin.Context) {
	amenities, err := h.amenityService.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  amenities,
		"total": len(amenities),
	})
}

func (h *AmenityHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var amenity domain.Amenity
	if err := c.ShouldBindJSON(&amenity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existing, err := h.amenityService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrAmenityNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "amenity not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mapper.MapAmenity(existing, &amenity)

	if err := h.amenityService.Update(c.Request.Context(), existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func (h *AmenityHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.amenityService.Delete(c.Request.Context(), id); err != nil {
		if err == service.ErrAmenityNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "amenity not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
