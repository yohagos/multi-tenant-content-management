// go/internal/adapter/repository/content.go
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
	"github.com/yohagos/multi-content-management/pkg/metrics"
)

type contentRepository struct {
	db *sqlx.DB
}

func NewContentRepository(db *sqlx.DB) port.ContentRepository {
	return &contentRepository{db: db}
}

func (r *contentRepository) Create(ctx context.Context, content *domain.Content) error {
	start := time.Now()

	query := `
		INSERT INTO contents (id, tenant_id, title, slug, body, status, published, published_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		content.ID, content.TenantID, content.Title, content.Slug,
		content.Body, content.Status, content.Published, content.PublishedAt,
		content.CreatedAt, content.UpdatedAt)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("create", "contents").Observe(duration)

	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("create", "contents", "error").Inc()
		return err
	}

	metrics.DbQueryTotal.WithLabelValues("create", "contents", "success").Inc()

	return nil
}

func (r *contentRepository) GetByID(ctx context.Context, id string) (*domain.Content, error) {
	start := time.Now()

	query := `SELECT id, tenant_id, title, slug, body, status, published, published_at, created_at, updated_at FROM contents WHERE id = $1 AND deleted_at IS NULL`

	var content domain.Content
	err := r.db.GetContext(ctx, &content, query, id)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("get_by_id", "contents").Observe(duration)

	if err == sql.ErrNoRows {
		metrics.DbQueryTotal.WithLabelValues("get_by_id", "contents", "not_found").Inc()
		return nil, nil
	}
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("get_by_id", "contents", "error").Inc()
		return nil, err
	}

	metrics.DbQueryTotal.WithLabelValues("get_by_id", "contents", "success").Inc()

	return &content, nil
}

func (r *contentRepository) GetBySlug(ctx context.Context, tenantID, slug string) (*domain.Content, error) {
	start := time.Now()

	query := `SELECT id, tenant_id, title, slug, body, status, published, published_at, created_at, updated_at FROM contents WHERE tenant_id = $1 AND slug = $2 AND deleted_at IS NULL`

	var content domain.Content
	err := r.db.GetContext(ctx, &content, query, tenantID, slug)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("get_by_slug", "contents").Observe(duration)

	if err == sql.ErrNoRows {
		metrics.DbQueryTotal.WithLabelValues("get_by_slug", "contents", "error").Inc()
		return nil, nil
	}
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("get_by_slug", "contents", "error").Inc()
		return nil, err
	}
	metrics.DbQueryTotal.WithLabelValues("get_by_slug", "contents", "success").Inc()

	return &content, nil
}

func (r *contentRepository) List(ctx context.Context, filter domain.ContentFilter) ([]domain.Content, int, error) {
	start := time.Now()

	conditions := []string{"deleted_at IS NULL"}
	args := []interface{}{}
	argIndex := 1

	if filter.TenantID != "" {
		conditions = append(conditions, fmt.Sprintf("tenant_id = $%d", argIndex))
		args = append(args, filter.TenantID)
		argIndex++
	}

	if filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(title ILIKE $%d OR slug ILIKE $%d)", argIndex, argIndex+1))
		args = append(args, "%"+filter.Search+"%", "%"+filter.Search+"%")
		argIndex += 2
	}

	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, filter.Status)
		argIndex++
	}

	if filter.Published != nil {
		conditions = append(conditions, fmt.Sprintf("published = $%d", argIndex))
		args = append(args, *filter.Published)
		argIndex++
	}

	whereClause := "WHERE " + strings.Join(conditions, " AND ")

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM contents %s", whereClause)
	var total int
	if err := r.db.GetContext(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}
	query := fmt.Sprintf(`
		SELECT id, tenant_id, title, slug, body, status, published, published_at, created_at, updated_at 
		FROM contents %s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)

	args = append(args, filter.Limit, filter.Offset)

	var contents []domain.Content
	err := r.db.SelectContext(ctx, &contents, query, args...)
	
	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("list", "contents").Observe(duration)

	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("list", "contents", "error").Inc()
		return nil, 0, err
	}
	metrics.DbQueryTotal.WithLabelValues("list", "contents", "success").Inc()

	return contents, total, nil
}

func (r *contentRepository) Update(ctx context.Context, content *domain.Content) error {
	start := time.Now()
	query := `
		UPDATE contents 
		SET title = $1, slug = $2, body = $3, status = $4, updated_at = $5
		WHERE id = $6 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		content.Title, content.Slug, content.Body,
		content.Status, content.UpdatedAt, content.ID)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("update", "contents").Observe(duration)

	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("update", "contents", "error").Inc()
		return err
	}
	metrics.DbQueryTotal.WithLabelValues("update", "contents", "success").Inc()

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *contentRepository) Delete(ctx context.Context, id string) error {
	start := time.Now()

	query := `UPDATE contents SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := r.db.ExecContext(ctx, query, id)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("delete", "contents").Observe(duration)

	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("delete", "contents", "error").Inc()
		return err
	}
	metrics.DbQueryTotal.WithLabelValues("delete", "contents", "success").Inc()

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *contentRepository) Publish(ctx context.Context, id string) error {
	now := time.Now()
	query := `
		UPDATE contents 
		SET published = true, status = 'published', published_at = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, now, now, id)
	
	duration := time.Since(now).Seconds()
	metrics.DbQueryDuration.WithLabelValues("publish", "contents").Observe(duration)
	
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("publish", "contents", "error").Inc()
		return err
	}
	metrics.DbQueryTotal.WithLabelValues("publish", "contents", "success").Inc()

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
