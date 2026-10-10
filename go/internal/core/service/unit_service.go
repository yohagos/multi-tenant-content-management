package service

import (
	"context"
	"errors"
	"time"

	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

var (
	ErrBuildingIdMissing = errors.New("building id is missing")
	ErrUnitExists        = errors.New("unit already exists")
	ErrUnitIdMissing     = errors.New("unit id missing")
	ErrUnitMissing       = errors.New("unit is missing")
	ErrUnitNotFound      = errors.New("unit not found")
)

type UnitService struct {
	unitRepo port.UnitRepository
}

func NewUnitRepository(unitRepo port.UnitRepository) *UnitService {
	return &UnitService{
		unitRepo: unitRepo,
	}
}

func (s *UnitService) Create(ctx context.Context, unit *domain.Unit) error {
	if unit.BuildingID == "" {
		return ErrBuildingIdMissing
	}

	unit.CreatedAt = time.Now()
	unit.UpdatedAt = time.Now()

	existing, _ := s.unitRepo.GetByTitleUnitNumber(ctx, unit.Title, unit.UnitNumber)
	if existing != nil {
		return ErrUnitExists
	}

	if err := s.unitRepo.Create(ctx, unit); err != nil {
		return err
	}
	return nil
}

func (s *UnitService) GetByID(ctx context.Context, id string) (*domain.Unit, error) {
	unit, err := s.unitRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return unit, nil
}

func (s *UnitService) List(ctx context.Context, filter domain.UnitFilter) ([]domain.Unit, int, error) {
	return s.unitRepo.List(ctx, &filter)
}

func (s *UnitService) Update(ctx context.Context, unit *domain.Unit) error {
	if unit == nil {
		return ErrUnitMissing
	}
	existing, err := s.unitRepo.GetByID(ctx, unit.ID)
	if err != nil {
		return err
	}

	unit.UpdatedAt = time.Now()
	unit.CreatedAt = existing.CreatedAt

	if err := s.unitRepo.Update(ctx, unit); err != nil {
		return err
	}
	return nil
}

func (s *UnitService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrUnitIdMissing
	}

	existing, err := s.unitRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrUnitNotFound
	}

	if err := s.unitRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
