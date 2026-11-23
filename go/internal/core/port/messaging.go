package port

import "context"

type EventPublisher interface {
	Publish(ctx context.Context, exchange, routingKey string, body []byte) error
	PublishContentCreated(ctx context.Context, tenantID, contentID string) error
	PublishContentPublished(ctx context.Context, tenantID, contentID string) error
	PublishTenantCreated(ctx context.Context, tenantID string) error
}