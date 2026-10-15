package service

import (
	"context"
	"errors"

	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

var (
	ErrInvalidAmenity         = errors.New("invalid amenity")
	ErrInvalidAmenityCategory = errors.New("invalid amenity category")
	ErrAmenityNotFound        = errors.New("amenity not found")
)

type AmenityService struct {
	amenityRepo port.AmenityRepository
}

func NewAmenityService(amenityRepo port.AmenityRepository) *AmenityService {
	return &AmenityService{
		amenityRepo: amenityRepo,
	}
}

func (s *AmenityService) Create(ctx context.Context, amenity *domain.Amenity) error {
	if amenity.Name == "" {
		return ErrInvalidAmenity
	}

	if !amenity.Category.IsValid() {
		return ErrInvalidAmenityCategory
	}

	return s.Create(ctx, amenity)
}

func (s *AmenityService) GetByID(ctx context.Context, id string) (*domain.Amenity, error) {
	if id == "" {
		return nil, ErrAmenityNotFound
	}

	return s.amenityRepo.GetByID(ctx, id)
}

func (s *AmenityService) List(ctx context.Context) ([]domain.Amenity, error) {
	return s.amenityRepo.List(ctx)
}

func (s *AmenityService) Update(ctx context.Context, amenity *domain.Amenity) error {
	_, err := s.amenityRepo.GetByID(ctx, amenity.ID)
	if err != nil {
		return err
	}

	return s.amenityRepo.Update(ctx, amenity)
}

func (s *AmenityService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidAmenity
	}

	return s.amenityRepo.Delete(ctx, id)
}
