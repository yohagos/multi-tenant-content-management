package messaging

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/yohagos/multi-content-management/internal/core/port"
)

type eventPublisher struct {
	conn *amqp.Connection
}

type ContentCreatedEvent struct {
	TenantID  string `json:"tenant_id"`
	ContentID string `json:"content_id"`
	Timestamp string `json:"timestamp"`
}

type ContentPublishedEvent struct {
	TenantID  string `json:"tenant_id"`
	ContentID string `json:"content_id"`
	Timestamp string `json:"timestamp"`
}

type TenantCreatedEvent struct {
	TenantID  string `json:"tenant_id"`
	Timestamp string `json:"timestamp"`
}

func NewEventPublisher(conn *amqp.Connection) port.EventPublisher {
	return &eventPublisher{conn: conn}
}

func (p *eventPublisher) Publish(ctx context.Context, exchange, routingKey string, body []byte) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return ch.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
		},
	)
}

func (p *eventPublisher) PublishContentCreated(ctx context.Context, tenantID, contentID string) error {
	event := ContentCreatedEvent{
		TenantID:  tenantID,
		ContentID: contentID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.Publish(ctx, "content-events", "content.created", body)
}

func (p *eventPublisher) PublishContentPublished(ctx context.Context, tenantID, contentID string) error {
	event := ContentPublishedEvent{
		TenantID:  tenantID,
		ContentID: contentID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.Publish(ctx, "content-events", "content.published", body)
}

func (p *eventPublisher) PublishTenantCreated(ctx context.Context, tenantID string) error {
	event := TenantCreatedEvent{
		TenantID:  tenantID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.Publish(ctx, "tenant-events", "tenant.created", body)
}
