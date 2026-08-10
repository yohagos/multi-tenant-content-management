package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

var (
	ErrTenantNotFound = errors.New("tenant not found")
	ErrTenantExists   = errors.New("tenant already exists")
)

type TenantService struct {
	tenantRepo port.TenantRepository
	cacheRepo  port.CacheRepository
	publisher  port.EventPublisher
}

func NewTenantService(tenantRepo port.TenantRepository, cacheRepo port.CacheRepository, publisher port.EventPublisher) *TenantService {
	return &TenantService{
		tenantRepo: tenantRepo,
		cacheRepo:  cacheRepo,
		publisher:  publisher,
	}
}

func (s *TenantService) Create(ctx context.Context, tenant *domain.Tenant) error {
	if tenant.ID == "" {
		tenant.ID = generateID()
	}
	tenant.CreatedAt = time.Now()
	tenant.UpdatedAt = time.Now()
	tenant.Active = true

	existing, _ := s.tenantRepo.GetBySlug(ctx, tenant.Slug)
	if existing != nil {
		return ErrTenantExists
	}

	if err := s.tenantRepo.Create(ctx, tenant); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("tenant:%s", tenant.ID)
	_ = s.cacheRepo.Delete(ctx, cacheKey)

	_ = s.publisher.PublishTenantCreated(ctx, tenant.ID)

	return nil
}

func (s *TenantService) GetByID(ctx context.Context, id string) (*domain.Tenant, error) {
	cacheKey := fmt.Sprintf("tenant:%s", id)
	if cached, err := s.cacheRepo.Get(ctx, cacheKey); err == nil && cached != "" {
		var tenant domain.Tenant
		if err := json.Unmarshal([]byte(cached), &tenant); err == nil {
			return &tenant, nil
		}
	}

	tenant, err := s.tenantRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, ErrTenantNotFound
	}

	cacheData, _ := json.Marshal(tenant)
	_ = s.cacheRepo.Set(ctx, cacheKey, string(cacheData), 5*time.Minute)

	return tenant, nil
}

func (s *TenantService) GetBySlug(ctx context.Context, slug string) (*domain.Tenant, error) {
	cacheKey := fmt.Sprintf("tenant:slug:%s", slug)
	if cached, err := s.cacheRepo.Get(ctx, cacheKey); err == nil && cached != "" {
		var tenant domain.Tenant
		if err := json.Unmarshal([]byte(cached), &tenant); err == nil {
			return &tenant, nil
		}
	}

	tenant, err := s.tenantRepo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, ErrTenantNotFound
	}

	cacheData, _ := json.Marshal(tenant)
	_ = s.cacheRepo.Set(ctx, cacheKey, string(cacheData), 5*time.Minute)

	return tenant, nil
}

func (s *TenantService) GetByDomain(ctx context.Context, dom string) (*domain.Tenant, error) {
	cacheKey := fmt.Sprintf("tenant:domain:%s", dom)
	if cached, err := s.cacheRepo.Get(ctx, cacheKey); err == nil && cached != "" {
		var tenant domain.Tenant
		if err := json.Unmarshal([]byte(cached), &tenant); err == nil {
			return &tenant, nil
		}
	}

	tenant, err := s.tenantRepo.GetByDomain(ctx, dom)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, ErrTenantNotFound
	}

	cacheData, _ := json.Marshal(tenant)
	_ = s.cacheRepo.Set(ctx, cacheKey, string(cacheData), 5*time.Minute)

	return tenant, nil
}

func (s *TenantService) List(ctx context.Context, filter domain.TenantFilter) ([]domain.Tenant, int, error) {
	return s.tenantRepo.List(ctx, filter)
}

func (s *TenantService) Update(ctx context.Context, tenant *domain.Tenant) error {
	existing, err := s.tenantRepo.GetByID(ctx, tenant.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrTenantNotFound
	}

	tenant.UpdatedAt = time.Now()
	tenant.CreatedAt = existing.CreatedAt

	if err := s.tenantRepo.Update(ctx, tenant); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("tenant:%s", tenant.ID)
	_ = s.cacheRepo.Delete(ctx, cacheKey)
	cacheKeySlug := fmt.Sprintf("tenant:slug:%s", existing.Slug)
	_ = s.cacheRepo.Delete(ctx, cacheKeySlug)
	cacheKeyDomain := fmt.Sprintf("tenant:domain:%s", existing.Domain)
	_ = s.cacheRepo.Delete(ctx, cacheKeyDomain)

	return nil
}

func (s *TenantService) Delete(ctx context.Context, id string) error {
	existing, err := s.tenantRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrTenantNotFound
	}

	if err := s.tenantRepo.Delete(ctx, id); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("tenant:%s", id)
	_ = s.cacheRepo.Delete(ctx, cacheKey)
	cacheKeySlug := fmt.Sprintf("tenant:slug:%s", existing.Slug)
	_ = s.cacheRepo.Delete(ctx, cacheKeySlug)
	cacheKeyDomain := fmt.Sprintf("tenant:domain:%s", existing.Domain)
	_ = s.cacheRepo.Delete(ctx, cacheKeyDomain)

	return nil
}

func generateID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return hex.EncodeToString(bytes)
	}
	return hex.EncodeToString(bytes)
}