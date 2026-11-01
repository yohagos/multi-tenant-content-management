package repository

import (
	"context"

	"github.com/yohagos/multi-content-management/pkg/metrics"
)

func (r *userRepository) UpdateUserMetrics(ctx context.Context) error {
	var activeUsers int
	query := `
		SELECT COUNT(*)
		FROM users
		WHERE active = true AND deleted_at IS NULL
	`

	if err := r.db.GetContext(ctx, &activeUsers, query); err != nil {
		return err
	}
	metrics.AuthActiveUsers.Set(float64(activeUsers))

	roleQuery := `
		SELECT role, COUNT(*) as count
		FROM users
		WHERE deleted_at IS NULL
		GROUP BY role
	`
	var results []struct {
		Role string `db:"role"`
		Count int `db:"count"`
	}
	if err := r.db.SelectContext(ctx, &results, roleQuery); err !=nil {
		return err
	}

	for _, result := range results {
		metrics.AuthUsersByRole.WithLabelValues(result.Role).Set(float64(result.Count))
	}

	return nil
}