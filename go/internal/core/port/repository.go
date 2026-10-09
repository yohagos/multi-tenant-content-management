package port

import (
	"context"

	"github.com/yohagos/multi-content-management/internal/core/domain"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *domain.Tenant) error
	GetByID(ctx context.Context, id string) (*domain.Tenant, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Tenant, error)
	GetByDomain(ctx context.Context, domain string) (*domain.Tenant, error)
	List(ctx context.Context, filter domain.TenantFilter) ([]domain.Tenant, int, error)
	Update(ctx context.Context, tenant *domain.Tenant) error
	Delete(ctx context.Context, id string) error
	UpdateTenantMetrics(ctx context.Context) error
}

type ContentRepository interface {
	Create(ctx context.Context, content *domain.Content) error
	GetByID(ctx context.Context, id string) (*domain.Content, error)
	GetBySlug(ctx context.Context, tenantID, slug string) (*domain.Content, error)
	List(ctx context.Context, filter domain.ContentFilter) ([]domain.Content, int, error)
	Update(ctx context.Context, content *domain.Content) error
	Delete(ctx context.Context, id string) error
	Publish(ctx context.Context, id string) error
	UpdateContentMetrics(ctx context.Context) error
}

type BuildingRepository interface {
	Create(ctx context.Context, building *domain.Building) error
	GetBySlug(ctx context.Context, slug string) (*domain.Building, error)
	GetByID(ctx context.Context, id string) (*domain.Building, error)
	List(ctx context.Context, filter domain.BuildingFilter) ([]domain.Building, int, error)
	Update(ctx context.Context, building *domain.Building) error
	Delete(ctx context.Context, id string) error
}

type UnitRepository interface {
	Create(ctx context.Context, unit *domain.Unit) error
	GetByID(ctx context.Context, id string) (*domain.Unit, error)
	List(ctx context.Context, filter *domain.UnitFilter) ([]domain.Unit, int, error)
	Update(ctx context.Context, unit *domain.Unit) error
	Delete(ctx context.Context, id string) error
}

type AmenityRepository interface {
}

type MediaRepository interface {
}
