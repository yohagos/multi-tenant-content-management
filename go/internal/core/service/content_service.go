package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

var (
	ErrContentNotFound = errors.New("content not found")
	ErrContentExists   = errors.New("content already exists")
)

type ContentService struct {
	contentRepo port.ContentRepository
	cacheRepo   port.CacheRepository
	publisher   port.EventPublisher
}

func NewContentService(contentRepo port.ContentRepository, cacheRepo port.CacheRepository, publisher port.EventPublisher) *ContentService {
	return &ContentService{
		contentRepo: contentRepo,
		cacheRepo:   cacheRepo,
		publisher:   publisher,
	}
}

func (s *ContentService) Create(ctx context.Context, content *domain.Content) error {
	if content.ID == "" {
		content.ID = generateID()
	}
	content.Slug = generateSlug(content.Title)
	content.CreatedAt = time.Now()
	content.UpdatedAt = time.Now()
	content.Published = false
	content.Status = "draft"

	existing, _ := s.contentRepo.GetBySlug(ctx, content.TenantID, content.Slug)
	if existing != nil {
		return ErrContentExists
	}

	if err := s.contentRepo.Create(ctx, content); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("content:%s", content.TenantID)
	_ = s.cacheRepo.Delete(ctx, cacheKey)
	cacheKeyID := fmt.Sprintf("content:id:%s", content.ID)
	_ = s.cacheRepo.Delete(ctx, cacheKeyID)

	_ = s.publisher.PublishContentCreated(ctx, content.TenantID, content.ID)

	return nil
}

func (s *ContentService) GetByID(ctx context.Context, id string) (*domain.Content, error) {
	cacheKey := fmt.Sprintf("content:id:%s", id)
	if cached, err := s.cacheRepo.Get(ctx, cacheKey); err == nil && cached != "" {
		var content domain.Content
		if err := json.Unmarshal([]byte(cached), &content); err == nil {
			return &content, nil
		}
	}

	content, err := s.contentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if content == nil {
		return nil, ErrContentNotFound
	}

	cacheData, _ := json.Marshal(content)
	_ = s.cacheRepo.Set(ctx, cacheKey, string(cacheData), 5*time.Minute)

	return content, nil
}

func (s *ContentService) GetBySlug(ctx context.Context, tenantID, slug string) (*domain.Content, error) {
	cacheKey := fmt.Sprintf("content:slug:%s:%s", tenantID, slug)
	if cached, err := s.cacheRepo.Get(ctx, cacheKey); err == nil && cached != "" {
		var content domain.Content
		if err := json.Unmarshal([]byte(cached), &content); err == nil {
			return &content, nil
		}
	}

	content, err := s.contentRepo.GetBySlug(ctx, tenantID, slug)
	if err != nil {
		return nil, err
	}
	if content == nil {
		return nil, ErrContentNotFound
	}

	cacheData, _ := json.Marshal(content)
	_ = s.cacheRepo.Set(ctx, cacheKey, string(cacheData), 5*time.Minute)

	return content, nil
}

func (s *ContentService) List(ctx context.Context, filter domain.ContentFilter) ([]domain.Content, int, error) {
	return s.contentRepo.List(ctx, filter)
}

func (s *ContentService) Update(ctx context.Context, content *domain.Content) error {
	existing, err := s.contentRepo.GetByID(ctx, content.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrContentNotFound
	}

	content.UpdatedAt = time.Now()
	content.CreatedAt = existing.CreatedAt
	content.Published = existing.Published
	content.PublishedAt = existing.PublishedAt

	if err := s.contentRepo.Update(ctx, content); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("content:id:%s", content.ID)
	_ = s.cacheRepo.Delete(ctx, cacheKey)
	cacheKeySlug := fmt.Sprintf("content:slug:%s:%s", existing.TenantID, existing.Slug)
	_ = s.cacheRepo.Delete(ctx, cacheKeySlug)
	cacheKeyTenant := fmt.Sprintf("content:%s", content.TenantID)
	_ = s.cacheRepo.Delete(ctx, cacheKeyTenant)

	return nil
}

func (s *ContentService) Delete(ctx context.Context, id string) error {
	existing, err := s.contentRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrContentNotFound
	}

	if err := s.contentRepo.Delete(ctx, id); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("content:id:%s", id)
	_ = s.cacheRepo.Delete(ctx, cacheKey)
	cacheKeySlug := fmt.Sprintf("content:slug:%s:%s", existing.TenantID, existing.Slug)
	_ = s.cacheRepo.Delete(ctx, cacheKeySlug)
	cacheKeyTenant := fmt.Sprintf("content:%s", existing.TenantID)
	_ = s.cacheRepo.Delete(ctx, cacheKeyTenant)

	return nil
}

func (s *ContentService) Publish(ctx context.Context, id string) error {
	existing, err := s.contentRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrContentNotFound
	}

	if err := s.contentRepo.Publish(ctx, id); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("content:id:%s", id)
	_ = s.cacheRepo.Delete(ctx, cacheKey)
	cacheKeySlug := fmt.Sprintf("content:slug:%s:%s", existing.TenantID, existing.Slug)
	_ = s.cacheRepo.Delete(ctx, cacheKeySlug)
	cacheKeyTenant := fmt.Sprintf("content:%s", existing.TenantID)
	_ = s.cacheRepo.Delete(ctx, cacheKeyTenant)

	_ = s.publisher.PublishContentPublished(ctx, existing.TenantID, id)

	return nil
}

func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, ".", "-")
	slug = strings.ReplaceAll(slug, ",", "-")
	slug = strings.ReplaceAll(slug, ":", "-")
	slug = strings.ReplaceAll(slug, ";", "-")
	slug = strings.ReplaceAll(slug, "!", "")
	slug = strings.ReplaceAll(slug, "?", "")
	slug = strings.ReplaceAll(slug, "&", "")
	slug = strings.ReplaceAll(slug, "(", "")
	slug = strings.ReplaceAll(slug, ")", "")

	if len(slug) > 50 {
		slug = slug[:50]
	}

	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err == nil {
		slug = fmt.Sprintf("%s-%s", slug, hex.EncodeToString(bytes))
	}

	return slug
}
