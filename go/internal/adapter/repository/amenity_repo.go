package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

type amenityRepository struct {
	db *sqlx.DB
}

var (
	ErrAmenityIdMissing = errors.New("amenity id missing")
	ErrAmenityNotFound  = errors.New("amenity not found")
)

func NewAmenityRepository(db *sqlx.DB) port.AmenityRepository {
	return &amenityRepository{
		db: db,
	}
}

func (r *amenityRepository) Create(ctx context.Context, amenity *domain.Amenity) error {
	amenity.CreatedAt = time.Now()

	query := `
		INSERT INTO amenities (
			code,
			name,
			category,
			icon,
			is_global,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		amenity.Code,
		amenity.Name,
		amenity.Category,
		amenity.Icon,
		amenity.IsGlobal,
		amenity.CreatedAt,
	).Scan(&amenity.ID)
}

func (r *amenityRepository) GetByID(ctx context.Context, id string) (*domain.Amenity, error) {
	if id == "" {
		return nil, ErrAmenityIdMissing
	}

	var amenity domain.Amenity

	query := `
		SELECT *
		FROM amenities
		WHERE id = $1
	`

	if err := r.db.GetContext(ctx, &amenity, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &amenity, nil
}

func (r *amenityRepository) List(ctx context.Context) ([]domain.Amenity, error) {
	var amenities []domain.Amenity

	query := `
		SELECT *
		FROM amenities
	`

	err := r.db.SelectContext(
		ctx,
		&amenities,
		query,
	)

	return amenities, err
}

func (r *amenityRepository) Update(ctx context.Context, amenity *domain.Amenity) error {
	if amenity.ID == "" {
		return ErrAmenityIdMissing
	}

	existing, err := r.GetByID(ctx, amenity.ID)
	if err != nil {
		return err
	}

	amenity.ID = existing.ID

	query := `
		UPDATE amenities
		SET 
			tenant_id = $2
			code = $3
			name = $4
			category = $5
			icon = $6
			is_global = $7
		WHERE id = $1
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		amenity.ID,
		amenity.TenantID,
		amenity.Code,
		amenity.Name,
		amenity.Category,
		amenity.Icon,
		amenity.IsGlobal,
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

func (r *amenityRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM amenities
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
