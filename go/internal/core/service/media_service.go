package service

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

type UploadMediaInput struct {
	TenantID   string
	EntityType domain.MediaEntityType
	EntityID   string
	FileName   string
	MimeType   string
	File       io.Reader
	AltText    string
	SortOrder  int
	IsPrimary  bool
}

var (
	ErrInvalidMedia         = errors.New("invalid media")
	ErrInvalidMediaType     = errors.New("invalid media entity type")
	ErrMediaNotFound        = errors.New("media not found")
	ErrMediaTooLarge        = errors.New("media exceeds maximum size")
	ErrUnsupportedMediaType = errors.New("unsupported media type")
)

const maxMediaSize int64 = 10 << 20

var allowedMediaTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
}

type MediaService struct {
	mediaRepo port.MediaRepository
	storage   port.MediaStorage
}

func NewMediaService(
	mediaRepo port.MediaRepository,
	storage port.MediaStorage,
) *MediaService {
	return &MediaService{
		mediaRepo: mediaRepo,
		storage:   storage,
	}
}

func (s *MediaService) Upload(
	ctx context.Context,
	input UploadMediaInput,
) (*domain.Media, error) {
	if input.EntityID == "" || input.FileName == "" || input.File == nil {
		return nil, ErrInvalidMedia
	}

	if input.EntityType != domain.MediaEntityTypeBuilding &&
		input.EntityType != domain.MediaEntityTypeUnit {
		return nil, ErrInvalidMediaType
	}

	input.FileName = filepath.Base(input.FileName)
	input.MimeType = strings.ToLower(strings.TrimSpace(input.MimeType))

	if !allowedMediaTypes[input.MimeType] {
		return nil, ErrUnsupportedMediaType
	}

	limitedReader := &io.LimitedReader{
		R: input.File,
		N: maxMediaSize + 1,
	}

	filePath, fileSize, err := s.storage.Save(
		ctx,
		limitedReader,
		input.FileName,
	)
	if err != nil {
		return nil, err
	}

	if fileSize == 0 {
		return nil, ErrInvalidMedia
	}

	if fileSize > maxMediaSize {
		return nil, ErrMediaTooLarge
	}

	media := &domain.Media{
		EntityType: input.EntityType,
		EntityID:   input.EntityID,
		FileName:   input.FileName,
		FilePath:   filePath,
		FileSize:   fileSize,
		MimeType:   input.MimeType,
		AltText:    input.AltText,
		SortOrder:  input.SortOrder,
		IsPrimary:  input.IsPrimary,
		CreatedAt:  time.Now(),
	}

	if err := s.mediaRepo.Create(ctx, media); err != nil {
		_ = s.storage.Delete(ctx, filePath)
		return nil, err
	}

	return media, nil
}

func (s *MediaService) ListByEntity(
	ctx context.Context,
	entityType domain.MediaEntityType,
	entityID string,
) ([]domain.Media, error) {
	if entityID == "" {
		return nil, ErrInvalidMedia
	}

	if entityType != domain.MediaEntityTypeBuilding &&
		entityType != domain.MediaEntityTypeUnit {
		return nil, ErrInvalidMediaType
	}

	return s.mediaRepo.ListByEntity(ctx, entityType, entityID)
}

func (s *MediaService) GetFile(
	ctx context.Context,
	id string,
) (*domain.Media, io.ReadCloser, error) {
	if id == "" {
		return nil, nil, ErrInvalidMedia
	}

	media, err := s.mediaRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	if media == nil {
		return nil, nil, ErrMediaNotFound
	}

	file, err := s.storage.Open(ctx, media.FilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil, ErrMediaNotFound
		}
		return nil, nil, err
	}

	return media, file, nil
}

func (s *MediaService) Delete(
	ctx context.Context,
	id string,
) error {
	if id == "" {
		return ErrInvalidMedia
	}

	media, err := s.mediaRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if media == nil {
		return ErrMediaNotFound
	}

	if err := s.mediaRepo.Delete(ctx, id); err != nil {
		return err
	}

	return s.storage.Delete(ctx, media.FilePath)
}
