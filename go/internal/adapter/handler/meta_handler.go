package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yohagos/multi-content-management/internal/core/domain"
)

type MetaDataHandler struct{}

func NewMetaDataHandler() *MetaDataHandler {
	return &MetaDataHandler{}
}

func (h *MetaDataHandler) GetEnums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"building_status": []gin.H{
			{"value": string (domain.BuildingStatusDraft), "label": "Draft"},
			{"value": string (domain.BuildingStatusPublished), "label": "Published"},
			{"value": string (domain.BuildingStatusArchived), "label": "Archived"},
		},
		"unit_status": []gin.H{
			{"value": string (domain.UnitStatusDraft), "label": "Draft"},
			{"value": string (domain.UnitStatusAvailable), "label": "Available"},
			{"value": string (domain.UnitStatusReserved), "label": "Reserved"},
			{"value": string (domain.UnitStatusRented), "label": "Rented"},
			{"value": string (domain.UnitStatusArchived), "label": "Archived"},
		},
		"amenity_category": []gin.H{
			{"value": string (domain.AmenityCategoryInterior), "label": "Interior"},
			{"value": string (domain.AmenityCategoryKitchen), "label": "Kitchen"},
			{"value": string (domain.AmenityCategoryBathroom), "label": "Bathroom"},
			{"value": string (domain.AmenityCategoryOutdoor), "label": "Outdoor"},
			{"value": string (domain.AmenityCategoryBuilding), "label": "Building"},
			{"value": string (domain.AmenityCategoryTechnology), "label": "Technology"},
			{"value": string (domain.AmenityCategoryParking), "label": "Parking"},
			{"value": string (domain.AmenityCategoryAccessibility), "label": "Accessibility"},
			{"value": string (domain.AmenityCategoryOther), "label": "Other"},
		},
		"media_entity_type": []gin.H{
			{"value": string (domain.MediaEntityTypeBuilding), "label": "Building"},
			{"value": string (domain.MediaEntityTypeUnit), "label": "Unit"},
		},
	})
}