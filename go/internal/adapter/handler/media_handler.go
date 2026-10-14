package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/service"
)

type MediaHandler struct {
	mediaService *service.MediaService
}

func NewMediaHandler(mediaService *service.MediaService) *MediaHandler {
	return &MediaHandler{
		mediaService: mediaService,
	}
}

func (h *MediaHandler) Upload(c *gin.Context) {
	entityType := domain.MediaEntityType(c.PostForm("entity_type"))
	entityID := c.PostForm("entity_id")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open uploaded file"})
		return
	}
	defer file.Close()

	sortOrder, err := strconv.Atoi(c.DefaultPostForm("sort_order", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sort_order"})
		return
	}

	isPrimary, err := strconv.ParseBool(c.DefaultPostForm("is_primary", "false"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid is_primary"})
		return
	}

	input := service.UploadMediaInput{
		TenantID:   c.PostForm("tenant_id"),
		EntityType: entityType,
		EntityID:   entityID,
		FileName:   fileHeader.Filename,
		MimeType:   fileHeader.Header.Get("Content-Type"),
		File:       file,
		AltText:    c.PostForm("alt_text"),
		SortOrder:  sortOrder,
		IsPrimary:  isPrimary,
	}

	media, err := h.mediaService.Upload(c.Request.Context(), input)
	if err != nil {
		switch err {
		case service.ErrInvalidMedia, service.ErrInvalidMediaType:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case service.ErrUnsupportedMediaType:
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": err.Error()})
		case service.ErrMediaTooLarge:
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, media)
}


func (h *MediaHandler) ListByEntity(c *gin.Context) {
	entityType := domain.MediaEntityType(c.Query("entity_type"))
	entityID := c.Query("entity_id")

	media, err := h.mediaService.ListByEntity(
		c.Request.Context(),
		entityType,
		entityID,
	)
	if err != nil {
		switch err {
		case service.ErrInvalidMedia, service.ErrInvalidMediaType:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  media,
		"total": len(media),
	})
}

func (h *MediaHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.mediaService.Delete(c.Request.Context(), id); err != nil {
		if err == service.ErrMediaNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
			return
		}

		if err == service.ErrInvalidMedia {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}