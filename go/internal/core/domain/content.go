package domain

import "time"

type Content struct {
	ID          string     `db:"id" json:"id"`
	TenantID    string     `db:"tenant_id" json:"tenant_id"`
	Title       string     `db:"title" json:"title"`
	Slug        string     `db:"slug" json:"slug"`
	Body        string     `db:"body" json:"body"`
	Status      string     `db:"status" json:"status"`
	Published   bool     `db:"published" json:"published"`
	PublishedAt *time.Time `db:"published_at" json:"published_at"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type ContentFilter struct {
	TenantID  string `json:"tenant_id"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Search    string `json:"search"`
	Status    string `json:"status"`
	Published *bool  `json:"published"`
}
