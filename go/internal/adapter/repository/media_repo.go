package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

type mediaRepository struct {
	db sqlx.DB
}

func NewMediaRepository(db *sqlx.DB) port.MediaRepository {
	return &mediaRepository{
		db: *db,
	}
}

func (r *mediaRepository) Create(
	ctx context.Context,
	media *domain.Media,
) error {
	query := `
		INSERT INTO media(
			entity_type,
			entity_id,
			file_name,
			file_path,
			file_size,
			mime_type,
			alt_text,
			sort_order,
			is_primary
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at
	`

	return r.db.QueryRowContext(
		ctx, 
		query,
		media.EntityType,
		media.EntityID,
		media.FileName,
		media.FilePath,
		media.FileSize,
		media.MimeType,
		media.AltText,
		media.SortOrder,
		media.IsPrimary,
	).Scan(&media.ID, &media.CreatedAt)
}

func (r *mediaRepository) GetByID(
	ctx context.Context,
	id string,
) (*domain.Media, error) {
	var media domain.Media

	query := `
		SELECT *
		FROM media
		WHERE id = $1
	`

	if err := r.db.GetContext(ctx, &media, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &media, nil
}

func (r *mediaRepository) ListByEntity(
	ctx context.Context,
	entityType domain.MediaEntityType,
	entityID string,
) ([]domain.Media, error) {
	var media []domain.Media

	query := `
		SELECT *
		FROM media
		WHERE entity_type = $1 AND entity_id = $2
		ORDER BY sort_order ASC, created_at ASC
	`

	err := r.db.SelectContext(
		ctx,
		&media,
		query,
		entityType,
		entityID,
	)

	return media, err
}

func (r *mediaRepository) Delete(
	ctx context.Context,
	id string,
) error {
	query := `
		DELETE FROM media
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}