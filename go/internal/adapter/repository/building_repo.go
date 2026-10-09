package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

type buildingRepository struct {
	db *sqlx.DB
}

func NewBuildingRepository(db *sqlx.DB) port.BuildingRepository {
	return &buildingRepository{
		db: db,
	}
}

func (r *buildingRepository) Create(ctx context.Context, building *domain.Building) error {
	query := `
		INSERT INTO buildings (name, slug, street, house_number, postal_code, city, country, latitude, longitude, year_build, total_floors, total_units, description, status, published, published_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	`

	_, err := r.db.ExecContext(ctx, query,
		building.Name, building.Slug, building.Street, building.HouseNumber,
		building.PostalCode, building.City, building.Country, building.Latitude,
		building.Longitude, building.YearBuild, building.TotalFloors,
		building.TotalUnits, building.Description, building.Status,
		building.Published, building.PublishedAt, building.CreatedAt,
		building.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *buildingRepository) GetBySlug(ctx context.Context, slug string) (*domain.Building, error) {
	var building domain.Building
	query := `
	SELECT 
		id, tenant_id, name, slug, street, house_number, postal_code, 
		city, country, latitude, longitude, year_build, total_floors, 
		total_units, description, status, published, published_at, 
		created_at, updated_at, deleted_at
	FROM buildings
	WHERE slug = $1 AND deleted_at IS NULL`

	err := r.db.GetContext(ctx, &building, query, slug)

	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &building, nil
}

func (r *buildingRepository) List(ctx context.Context, filter domain.BuildingFilter) ([]domain.Building, int, error) {
	conditions := []string{"deleted_at IS NULL"}
	args := []interface{}{}
	argIndex := 1

	if filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR slug ILIKE $%d)", argIndex, argIndex+1))
		args = append(args, "%"+filter.Search+"%", "%"+filter.Search+"%")
		argIndex += 2
	}

	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, filter.Status)
		argIndex++
	}

	if filter.TenantID != "" {
		conditions = append(conditions, fmt.Sprintf("tenant_id = $%d", argIndex))
		args = append(args, filter.TenantID)
		argIndex++
	}

	if filter.Published != false {
		conditions = append(conditions, fmt.Sprintf("published = $%d", argIndex))
		args = append(args, filter.Published)
		argIndex++
	}

	whereClause := "WHERE " + strings.Join(conditions, " AND ")

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM buildings %s", whereClause)
	var total int
	if err := r.db.GetContext(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}

	query := fmt.Sprintf(`
		SELECT 
			id, tenant_id, name, slug, street, house_number, postal_code, 
			city, country, latitude, longitude, year_build, total_floors, 
			total_units, description, status, published, published_at, 
			created_at, updated_at, deleted_at
		FROM buildings %s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)

	args = append(args, filter.Limit, filter.Offset)

	var buildings []domain.Building

	err := r.db.SelectContext(ctx, &buildings, query, args...)
	if err != nil {
		return nil, 0, err
	}

	return buildings, total, nil
}

func (r *buildingRepository) GetByID(ctx context.Context, id string) (*domain.Building, error) {
	query := `
	SELECT 
		id, tenant_id, name, slug, street, house_number, postal_code, 
		city, country, latitude, longitude, year_build, total_floors, 
		total_units, description, status, published, published_at, 
		created_at, updated_at, deleted_at
	FROM buildings
	WHERE id = $1 AND deleted_at IS NULL`

	var building domain.Building

	err := r.db.GetContext(ctx, &building, query, id)

	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}

	return &building, nil
}

func (r *buildingRepository) Update(ctx context.Context, building *domain.Building) error {
	if building.TenantID != "" {
		query := `
			UPDATE buildings
			SET 
				tenant_id = $2, name = $3, slug = $4, street = $5, house_number = $6, postal_code = $7, 
				city = $8, country = $9, latitude = $10, longitude = $11, year_build = $12, total_floors = $13, 
				total_units = $14, description = $15, status = $16, published = $17, published_at = $18, 
				created_at = $19, updated_at = $20
			WHERE id = $1 AND deleted_at IS NULL
		`
		result, err := r.db.ExecContext(ctx, query,
			building.ID, building.TenantID, building.Name, building.Slug, building.HouseNumber, building.PostalCode, building.City, building.Country,
			building.Latitude, building.Longitude, building.YearBuild, building.TotalFloors, building.TotalUnits,
			building.Description, building.Status, building.Published, building.PublishedAt, building.CreatedAt, building.UpdatedAt,
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
	} else {
		query := `
			UPDATE buildings
			SET 
				name = $2, slug = $3, street = $4, house_number = $5, postal_code = $6, 
				city = $7, country = $8, latitude = $9, longitude = $10, year_build = $11, total_floors = $12, 
				total_units = $13, description = $14, status = $15, published = $16, published_at = $17, 
				created_at = $18, updated_at = $19
			WHERE id = $1 AND deleted_at IS NULL
		`
		result, err := r.db.ExecContext(ctx, query,
			building.Name, building.Slug, building.HouseNumber, building.PostalCode, building.City, building.Country,
			building.Latitude, building.Longitude, building.YearBuild, building.TotalFloors, building.TotalUnits,
			building.Description, building.Status, building.Published, building.PublishedAt, building.CreatedAt, building.UpdatedAt,
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
}

func (r *buildingRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE buildings
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}
