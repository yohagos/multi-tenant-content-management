package domain

import "time"

type AmenityCategory string

const (
	AmenityCategoryInterior      AmenityCategory = "interior"
	AmenityCategoryKitchen       AmenityCategory = "kitchen"
	AmenityCategoryBathroom      AmenityCategory = "bathroom"
	AmenityCategoryOutdoor       AmenityCategory = "outdoor"
	AmenityCategoryBuilding      AmenityCategory = "building"
	AmenityCategoryTechnology    AmenityCategory = "technology"
	AmenityCategoryParking       AmenityCategory = "parking"
	AmenityCategoryAccessibility AmenityCategory = "accessibility"
	AmenityCategoryOther         AmenityCategory = "other"
)

type MediaEntityType string

const (
	MediaEntityTypeBuilding MediaEntityType = "building"
	MediaEntityTypeUnit     MediaEntityType = "unit"
)

type Amenity struct {
	ID        string          `db:"id" json:"id"`
	TenantID  *string         `db:"tenant_id" json:"tenant_id,omitempty"`
	Code      string          `db:"code" json:"code"`
	Name      string          `db:"name" json:"name"`
	Category  AmenityCategory `db:"category" json:"category"`
	Icon      string          `db:"icon" json:"icon"`
	IsGlobal  bool            `db:"is_global" json:"is_global"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
}

type Media struct {
	ID         string          `db:"id" json:"id"`
	TenantID   string          `db:"tenant_id" json:"tenant_id"`
	EntityType MediaEntityType `db:"entity_type" json:"entity_type"`
	EntityID   string          `db:"entity_id" json:"entity_id"`
	FileName   string          `db:"file_name" json:"file_name"`
	FilePath   string          `db:"file_path" json:"file_path"`
	FileSize   int64           `db:"file_size" json:"file_size"`
	MimeType   string          `db:"mime_type" json:"mime_type"`
	AltText    string          `db:"alt_text" json:"alt_text"`
	SortOrder  int             `db:"sort_order" json:"sort_order"`
	IsPrimary  bool            `db:"is_primary" json:"is_primary"`
	CreatedAt  time.Time       `db:"created_at" json:"created_at"`
}
