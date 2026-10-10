package domain

import "time"

type UnitStatus string

const (
	UnitStatusDraft UnitStatus = "draft"
	UnitStatusAvailable UnitStatus = "available"
	UnitStatusReserved UnitStatus = "reserved"
	UnitStatusRented UnitStatus = "rented"
	UnitStatusArchived UnitStatus = "archived"
)

func (s UnitStatus) IsValid() bool {
	switch s {
	case UnitStatusDraft, UnitStatusAvailable, UnitStatusReserved, UnitStatusRented, UnitStatusArchived:
		return true
	}
	return false
}

type Unit struct {
	ID             string     `db:"id" json:"id"`
	TenantID       string     `db:"tenant_id" json:"tenant_id,omitempty"`
	BuildingID     string     `db:"building_id" json:"building_id"`
	UnitNumber     string     `db:"unit_number" json:"unit_number"`
	Floor          *int       `db:"floor" json:"floor,omitempty"`
	Rooms          float64    `db:"rooms" json:"rooms"`
	Bedrooms       int        `db:"bedrooms" json:"bedrooms"`
	Bathrooms      float64    `db:"bathrooms" json:"bathrooms"`
	LivingAreaSqm  *float64   `db:"living_area_sqm" json:"living_area_sqm,omitempty"`
	TotalAreaSqm   *float64   `db:"total_area_sqm" json:"total_area_sqm,omitempty"`
	BalconyAreaSqm *float64   `db:"balcony_area_sqm" json:"balcony_area_sqm,omitempty"`
	CeilingHeightM *float64   `db:"ceiling_height_m" json:"ceiling_height_m,omitempty"`
	ColdRent       *float64   `db:"cold_rent" json:"cold_rent,omitempty"`
	WarmRent       *float64   `db:"warm_rent" json:"warm_rent,omitempty"`
	Deposit        *float64   `db:"deposit" json:"deposit,omitempty"`
	AvailableFrom  *time.Time `db:"available_from" json:"available_from,omitempty"`
	Status         UnitStatus     `db:"status" json:"status"`
	Published      bool       `db:"published" json:"published"`
	PublishedAt    *time.Time `db:"published_at" json:"published_at,omitempty"`
	Title          string     `db:"title" json:"title"`
	Description    string     `db:"description" json:"description"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type UnitFilter struct {
	TenantID   string   `json:"tenant_id"`
	BuildingID string   `json:"building_id"`
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	Search     string   `json:"search"`
	Status     string   `json:"status"`
	MinRooms   *float64 `json:"min_rooms"`
	MaxRooms   *float64 `json:"max_rooms"`
	MinRent    *float64 `json:"min_rent"`
	MaxRent    *float64 `json:"max_rent"`
	Published  *bool    `json:"published"`
}