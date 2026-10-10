package service

import (
	"context"
	"errors"
	"time"

	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

var (
	ErrNameCityStreetMissing = errors.New("necessary information are missing")
	ErrBuildingExists = errors.New("building already exists")
	ErrBuildingNotFound = errors.New("building not found")
)

type BuildingService struct {
	buildingRepo port.BuildingRepository
}

func NewBuildingService(buildingRepo port.BuildingRepository) *BuildingService {
	return &BuildingService{
		buildingRepo: buildingRepo,
	}
}

func (s *BuildingService) Create(ctx context.Context, building *domain.Building) error {
	if building.Name == "" || building.City == "" || building.Street == "" {
		return ErrNameCityStreetMissing
	}

	building.CreatedAt = time.Now()
	building.UpdatedAt = time.Now()

	existing, _ := s.buildingRepo.GetBySlug(ctx, building.Slug)
	if existing != nil {
		return ErrBuildingExists
	}

	if err := s.buildingRepo.Create(ctx, building); err != nil {
		return err
	}

	return nil
}

func (s *BuildingService) GetBySlug(ctx context.Context, slug string) (*domain.Building, error) {
	building, err := s.buildingRepo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	if building == nil {
		return nil, ErrBuildingNotFound
	}

	return building, nil
}

func (s *BuildingService) GetByID(ctx context.Context, id string) (*domain.Building, error) {
	building, err := s.buildingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return building, nil
}

func (s *BuildingService) List(ctx context.Context, filter domain.BuildingFilter) ([]domain.Building, int, error) {
	return s.buildingRepo.List(ctx, filter)
}

func (s *BuildingService) Update(ctx context.Context, building *domain.Building) error {
	existing, err := s.buildingRepo.GetByID(ctx, building.ID)
	if err != nil {
		return err
	}
	if building == nil {
		return ErrBuildingNotFound
	}

	building.UpdatedAt = time.Now()
	building.CreatedAt = existing.CreatedAt

	if err := s.buildingRepo.Update(ctx, building); err != nil {
		return err
	}

	return nil
}

func (s *BuildingService) Delete(ctx context.Context, id string) error {
	existing, err := s.buildingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrBuildingNotFound
	}

	if err := s.buildingRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}