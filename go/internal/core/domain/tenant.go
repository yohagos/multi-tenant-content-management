package domain

import "time"

type Tenant struct {
	ID        string       `db:"id" json:"id"`
	Name      string       `db:"name" json:"name"`
	Slug      string       `db:"slug" json:"slug"`
	Domain    string       `db:"domain" json:"domain"`
	Config    TenantConfig `db:"config" json:"config"`
	Active    bool         `db:"active" json:"active"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time   `db:"deleted_at" json:"deleted_at,omitempty"`
}

type TenantConfig struct {
	Theme    string            `json:"theme"`
	Features []string          `json:"features"`
	Settings map[string]string `json:"settings"`
}

type TenantFilter struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Search string `json:"search"`
	Active bool   `json:"active"`
}
