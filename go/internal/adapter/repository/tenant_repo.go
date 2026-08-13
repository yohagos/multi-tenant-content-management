package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
	"github.com/yohagos/multi-content-management/pkg/metrics"
)

type tenantRepository struct {
	db *sqlx.DB
}

func NewTenantRepository(db *sqlx.DB) port.TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	start := time.Now()

	query := `
		INSERT INTO tenants (id, name, slug, domain, config, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	configJSON, err := json.Marshal(tenant.Config)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query,
		tenant.ID, tenant.Name, tenant.Slug, tenant.Domain,
		configJSON, tenant.Active, tenant.CreatedAt, tenant.UpdatedAt)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("create", "tenants").Observe(duration)

	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("create", "tenants", "error").Inc()
		return err
	}

	metrics.DbQueryTotal.WithLabelValues("create", "tenants", "success").Inc()
	return nil
}

func (r *tenantRepository) GetByID(ctx context.Context, id string) (*domain.Tenant, error) {
	start := time.Now()

	query := `SELECT id, name, slug, domain, config, active, created_at, updated_at FROM tenants WHERE id = $1 AND deleted_at IS NULL`

	var tenant domain.Tenant

	var tenantWrapper struct {
		ID        string    `db:"id"`
		Name      string    `db:"name"`
		Slug      string    `db:"slug"`
		Domain    string    `db:"domain"`
		Config    []byte    `db:"config"`
		Active    bool      `db:"active"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	err := r.db.GetContext(ctx, &tenantWrapper, query, id)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("get_by_id", "tenants").Observe(duration)

	if err == sql.ErrNoRows {
		metrics.DbQueryTotal.WithLabelValues("get_by_id", "tenants", "not_found").Inc()
		return nil, nil
	}
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("get_by_id", "tenants", "error").Inc()
		return nil, err
	}
	metrics.DbQueryTotal.WithLabelValues("get_by_id", "tenants", "success").Inc()

	tenant.ID = tenantWrapper.ID
	tenant.Name = tenantWrapper.Name
	tenant.Slug = tenantWrapper.Slug
	tenant.Domain = tenantWrapper.Domain
	tenant.Active = tenantWrapper.Active
	tenant.CreatedAt = tenantWrapper.CreatedAt
	tenant.UpdatedAt = tenantWrapper.UpdatedAt

	if len(tenantWrapper.Config) > 0 {
		if err := json.Unmarshal(tenantWrapper.Config, &tenant.Config); err != nil {
			return nil, err
		}
	}

	return &tenant, nil
}

func (r *tenantRepository) GetBySlug(ctx context.Context, slug string) (*domain.Tenant, error) {
	start := time.Now()

	query := `SELECT id, name, slug, domain, config, active, created_at, updated_at FROM tenants WHERE slug = $1 AND deleted_at IS NULL`

	var tenant domain.Tenant
	var tenantWrapper struct {
		ID        string    `db:"id"`
		Name      string    `db:"name"`
		Slug      string    `db:"slug"`
		Domain    string    `db:"domain"`
		Config    []byte    `db:"config"`
		Active    bool      `db:"active"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	err := r.db.GetContext(ctx, &tenantWrapper, query, slug)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("get_by_slug", "tenants").Observe(duration)

	if err == sql.ErrNoRows {
		metrics.DbQueryTotal.WithLabelValues("get_by_slug", "tenants", "not_found").Inc()
		return nil, err
	}
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("get_by_slug", "tenants", "error").Inc()
		return nil, err
	}

	metrics.DbQueryTotal.WithLabelValues("get_by_slug", "tenants", "success").Inc()

	tenant.ID = tenantWrapper.ID
	tenant.Name = tenantWrapper.Name
	tenant.Slug = tenantWrapper.Slug
	tenant.Domain = tenantWrapper.Domain
	tenant.Active = tenantWrapper.Active
	tenant.CreatedAt = tenantWrapper.CreatedAt
	tenant.UpdatedAt = tenantWrapper.UpdatedAt

	if len(tenantWrapper.Config) > 0 {
		if err := json.Unmarshal(tenantWrapper.Config, &tenant.Config); err != nil {
			return nil, err
		}
	}

	return &tenant, nil
}

func (r *tenantRepository) GetByDomain(ctx context.Context, dom string) (*domain.Tenant, error) {
	start := time.Now()

	query := `SELECT id, name, slug, domain, config, active, created_at, updated_at FROM tenants WHERE domain = $1 AND deleted_at IS NULL`

	var tenant domain.Tenant
	var tenantWrapper struct {
		ID        string    `db:"id"`
		Name      string    `db:"name"`
		Slug      string    `db:"slug"`
		Domain    string    `db:"domain"`
		Config    []byte    `db:"config"`
		Active    bool      `db:"active"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	err := r.db.GetContext(ctx, &tenantWrapper, query, dom)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("get_by_domain", "tenants").Observe(duration)

	if err == sql.ErrNoRows {
		metrics.DbQueryTotal.WithLabelValues("get_by_domain", "tenants", "not_found").Inc()
		return nil, err
	}
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("get_by_domain", "tenants", "error").Inc()
		return nil, err
	}
	metrics.DbQueryTotal.WithLabelValues("get_by_domain", "tenants", "success").Inc()

	tenant.ID = tenantWrapper.ID
	tenant.Name = tenantWrapper.Name
	tenant.Slug = tenantWrapper.Slug
	tenant.Domain = tenantWrapper.Domain
	tenant.Active = tenantWrapper.Active
	tenant.CreatedAt = tenantWrapper.CreatedAt
	tenant.UpdatedAt = tenantWrapper.UpdatedAt

	if len(tenantWrapper.Config) > 0 {
		if err := json.Unmarshal(tenantWrapper.Config, &tenant.Config); err != nil {
			return nil, err
		}
	}

	return &tenant, nil
}

func (r *tenantRepository) List(ctx context.Context, filter domain.TenantFilter) ([]domain.Tenant, int, error) {
	start := time.Now()

	conditions := []string{"deleted_at IS NULL"}
	args := []interface{}{}
	argIndex := 1

	if filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR slug ILIKE $%d)", argIndex, argIndex+1))
		args = append(args, "%"+filter.Search+"%", "%"+filter.Search+"%")
		argIndex += 2
	}

	if filter.Active != false {
		conditions = append(conditions, fmt.Sprintf("active = $%d", argIndex))
		args = append(args, filter.Active)
		argIndex++
	}

	whereClause := "WHERE " + strings.Join(conditions, " AND ")

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM tenants %s", whereClause)
	var total int
	if err := r.db.GetContext(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, err
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}
	query := fmt.Sprintf(`
		SELECT id, name, slug, domain, config, active, created_at, updated_at 
		FROM tenants %s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)

	args = append(args, filter.Limit, filter.Offset)

	var tenantWrappers []struct {
		ID        string    `db:"id"`
		Name      string    `db:"name"`
		Slug      string    `db:"slug"`
		Domain    string    `db:"domain"`
		Config    []byte    `db:"config"`
		Active    bool      `db:"active"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	err := r.db.SelectContext(ctx, &tenantWrappers, query, args...)
	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("list", "tenants").Observe(duration)
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("list", "tenants", "error").Inc()
		return nil, 0, err
	}
	metrics.DbQueryTotal.WithLabelValues("list", "tenants", "success").Inc()

	tenants := make([]domain.Tenant, len(tenantWrappers))
	for i, w := range tenantWrappers {
		tenants[i].ID = w.ID
		tenants[i].Name = w.Name
		tenants[i].Slug = w.Slug
		tenants[i].Domain = w.Domain
		tenants[i].Active = w.Active
		tenants[i].CreatedAt = w.CreatedAt
		tenants[i].UpdatedAt = w.UpdatedAt
		if len(w.Config) > 0 {
			if err := json.Unmarshal(w.Config, &tenants[i].Config); err != nil {
				return nil, 0, err
			}
		}
	}

	return tenants, total, nil
}

func (r *tenantRepository) Update(ctx context.Context, tenant *domain.Tenant) error {
	start := time.Now()

	query := `
		UPDATE tenants 
		SET name = $1, slug = $2, domain = $3, config = $4, active = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	configJSON, err := json.Marshal(tenant.Config)
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, query,
		tenant.Name, tenant.Slug, tenant.Domain,
		configJSON, tenant.Active, tenant.UpdatedAt, tenant.ID)
	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("update", "tenants").Observe(duration)
	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("update", "tenants", "error").Inc()
		return err
	}
	metrics.DbQueryTotal.WithLabelValues("update", "tenants", "success").Inc()

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *tenantRepository) Delete(ctx context.Context, id string) error {
	start := time.Now()
	query := `UPDATE tenants SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := r.db.ExecContext(ctx, query, id)

	duration := time.Since(start).Seconds()
	metrics.DbQueryDuration.WithLabelValues("delete", "tenants").Observe(duration)

	if err != nil {
		metrics.DbQueryTotal.WithLabelValues("delete", "tenants", "error").Inc()
		return err
	}
	metrics.DbQueryTotal.WithLabelValues("delete", "tenants", "success").Inc()

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
