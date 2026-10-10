package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

var (
	ErrBuildingIdMissing = errors.New("building id is missing")
)

type unitRepository struct {
	db *sqlx.DB
}

func NewUnitRepository(db *sqlx.DB) port.UnitRepository {
	return &unitRepository{
		db: db,
	}
}

func (r *unitRepository) Create(ctx context.Context, unit *domain.Unit) error {
	if unit.BuildingID == "" {
		return ErrBuildingIdMissing
	}

	query := `
		INSERT INTO units (
			building_id, unit_number, floor, rooms, bedrooms, bathrooms, living_area_sqm, total_area_sqm,
			balcony_area_sqm, ceiling_height_m, cold_rent, warm_rent, deposit, available_from, status,
			published, published_at, title, description, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 
			$12, $13, $14, $15, $16, $17, $18, $19, $20, $21
		) 
	`

	_, err := r.db.ExecContext(ctx, query,
		unit.BuildingID, unit.UnitNumber, unit.Floor, unit.Rooms, unit.Bedrooms, unit.Bathrooms, unit.LivingAreaSqm,
		unit.TotalAreaSqm, unit.BalconyAreaSqm, unit.CeilingHeightM, unit.ColdRent, unit.WarmRent, unit.Deposit, unit.AvailableFrom,
		unit.Status, unit.Published, unit.PublishedAt, unit.Title, unit.Description, unit.CreatedAt, unit.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *unitRepository) GetByID(ctx context.Context, id string) (*domain.Unit, error) {
	query := `
		SELECT
			id, tenant_id, building_id, unit_number, floor, rooms, bedrooms, bathrooms, living_area_sqm, total_area_sqm,
			balcony_area_sqm, ceiling_height_m, cold_rent, warm_rent, deposit, available_from, status,
			published, published_at, title, description, created_at, updated_at
		FROM units
		WHERE id = $1 AND deleted_at IS NULL
	`

	var unit domain.Unit

	err := r.db.GetContext(ctx, &unit, query, id)

	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *unitRepository) GetByTitleUnitNumber(ctx context.Context, title string, unit_number string) (*domain.Unit, error) {
	query := `
		SELECT
			id, tenant_id, building_id, unit_number, floor, rooms, bedrooms, bathrooms, living_area_sqm, total_area_sqm,
			balcony_area_sqm, ceiling_height_m, cold_rent, warm_rent, deposit, available_from, status,
			published, published_at, title, description, created_at, updated_at
		FROM units
		WHERE title = $1 AND unit_number = $2 AND deleted_at IS NULL
	`

	var unit domain.Unit

	err := r.db.GetContext(ctx, &unit, query, title, unit_number)

	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *unitRepository) List(ctx context.Context, filter *domain.UnitFilter) ([]domain.Unit, int, error) {
	conditions := []string{"deleted IS NULL"}
	args := []interface{}{}
	argIndex := 1

	if filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(title ILIKE $%d)", argIndex))
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, filter.Status)
		argIndex++
	}

	if filter.MinRooms != nil {
		conditions = append(conditions, fmt.Sprintf("min_rooms <= $%d", argIndex))
		args = append(args, filter.MinRooms)
		argIndex++
	}

	if filter.MaxRooms != nil {
		conditions = append(conditions, fmt.Sprintf("max_rooms >= $%d", argIndex))
		args = append(args, filter.MaxRooms)
		argIndex++
	}

	if filter.MinRent != nil {
		conditions = append(conditions, fmt.Sprintf("cold_rent <= $%d", argIndex))
		args = append(args, filter.MinRent)
		argIndex++
	}

	if filter.MaxRent != nil {
		conditions = append(conditions, fmt.Sprintf("cold_rent >= $%d", argIndex))
		args = append(args, filter.MaxRent)
		argIndex++
	}

	if filter.Published != nil {
		conditions = append(conditions, fmt.Sprintf("published = $%d", argIndex))
		args = append(args, filter.Published)
		argIndex++
	}

	whereClause := "WHERE " + strings.Join(conditions, " AND ")

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM units %s", whereClause)
	var total int
	if err := r.db.GetContext(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	query := fmt.Sprintf(`
		SELECT
			id, tenant_id, building_id, unit_number, floor, rooms, bedrooms, bathrooms, living_area_sqm, total_area_sqm,
			balcony_area_sqm, ceiling_height_m, cold_rent, warm_rent, deposit, available_from, status,
			published, published_at, title, description, created_at, updated_at, deleted_at
		FROM units %s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)

	args = append(args, filter.Limit, filter.Offset)

	var units []domain.Unit

	err := r.db.SelectContext(ctx, &units, query, args...)
	if err != nil {
		return nil, 0, err
	}

	return units, total, nil
}

func (r *unitRepository) Update(ctx context.Context, unit *domain.Unit) error {
	unit.UpdatedAt = time.Now()

	query := `
		UPDATE units 
		SET
			tenant_id = $2, unit_number = $3, floor = $4, rooms = $5, 
			bedrooms = $6, bathrooms = $7, living_area_sqm = $8, total_area_sqm = $9,
			balcony_area_sqm = $10, ceiling_height_m = $11, cold_rent = $12, warm_rent = $13, 
			deposit = $14, available_from = $15, status = $16, published = $17, published_at = $18, 
			title = $19, description = $20, created_at = $21, updated_at = $22
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		unit.ID, unit.TenantID, unit.UnitNumber, unit.Floor, unit.Rooms, unit.Bedrooms, unit.Bathrooms,
		unit.LivingAreaSqm, unit.TotalAreaSqm, unit.BalconyAreaSqm, unit.CeilingHeightM, unit.ColdRent,
		unit.WarmRent, unit.Deposit, unit.AvailableFrom, unit.Status, unit.Published, unit.PublishedAt,
		unit.Title, unit.Description, unit.CreatedAt, unit.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *unitRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE units
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
