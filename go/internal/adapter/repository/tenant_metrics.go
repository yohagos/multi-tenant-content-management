package repository

import (
	"context"

	"github.com/yohagos/multi-content-management/pkg/metrics"
)

func (r *tenantRepository) UpdateTenantMetrics(ctx context.Context) error {
	total, active, err := r.getTenantStats(ctx)
	if err != nil {
		return err
	}

	metrics.TenantsTotal.Set(float64(total))
	metrics.TenantsActive.Set(float64(active))

	return nil
}

func (r *tenantRepository) getTenantStats(ctx context.Context) (int, int, error) {
	var total, active int

	queryTotal := `SELECT COUNT(*) FROM tenants WHERE deleted_at IS NULL`
	err := r.db.GetContext(ctx, &total, queryTotal)
	if err != nil {
		return 0, 0, err
	}

	queryActive := `SELECT COUNT(*) FROM tenants WHERE active = true AND deleted_at IS NULL`
	err = r.db.GetContext(ctx, &active, queryActive)
	if err != nil {
		return 0, 0, err
	}

	return total, active, nil
}
