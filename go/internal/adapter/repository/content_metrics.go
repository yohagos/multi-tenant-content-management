package repository

import (
	"context"

	"github.com/yohagos/multi-content-management/pkg/metrics"
)

func (r *contentRepository) UpdateContentMetrics(ctx context.Context) error {
	stats, err := r.getContentStats(ctx)
	if err != nil {
		return err
	}

	for tenantID, statusCodes := range stats {
		for status, count := range statusCodes {
			metrics.ContentsTotal.WithLabelValues(tenantID, status).Set(float64(count))
		}
	}

	return nil
}

func (r *contentRepository) getContentStats(ctx context.Context) (map[string]map[string]int, error) {
	query := `
		SELECT tenant_id, status, COUNT(*) as count
		FROM contents
		WHERE deleted_at IS NULL
		GROUP BY tenant_id, status
	`

	var results []struct {
		TenantID string `db:"tenant_id"`
		Status string `db:"status"`
		Count int `db:"count"`
	}

	err := r.db.SelectContext(ctx, &results, query)
	if err != nil {
		return nil, err
	}

	stats := make(map[string]map[string]int)
	for _, result := range results {
		if _, exists := stats[result.TenantID]; !exists {
			stats[result.TenantID] = make(map[string]int)
		}
		stats[result.TenantID][result.Status] = result.Count
	}

	return stats, nil
}