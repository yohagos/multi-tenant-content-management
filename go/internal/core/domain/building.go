package domain

import "time"

type Building struct {
	ID          string     `db:"id" json:"id"`
	TenantID    string     `db:"tenant_id" json:"tenant_id"`
	Name        string     `db:"name" json:"name"`
	Slug        string     `db:"slug" json:"slug"`
	Street      string     `db:"street" json:"street"`
	HouseNumber string     `db:"house_number" json:"house_number"`
	PostalCode  string     `db:"postal_code" json:"postal_code"`
	City        string     `db:"city" json:"city"`
	Latitude    *float64   `db:"latitude" json:"latitude"`
	Longitude   *float64   `db:"longitude" json:"longitude"`
	YearBuild   *int       `db:"year_build" json:"year_build"`
	TotalFloors *int       `db:"total_floors" json:"total_floors"`
	TotalUnits  int        `db:"total_units" json:"total_units"`
	Description string     `db:"description" json:"description"`
	Status      string     `db:"status" json:"status"`
	Published   bool       `db:"published" json:"published"`
	PublishedAt *time.Time `db:"published_at" jso:"published_at"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type BuildingFilter struct {
	TenantID  string `json:"tenant_id"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Search    string `json:"search"`
	City      string `json:"city"`
	Status    string `json:"status"`
	Published bool   `json:"published"`
}
