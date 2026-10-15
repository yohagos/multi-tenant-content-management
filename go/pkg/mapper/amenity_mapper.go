package mapper

import "github.com/yohagos/multi-content-management/internal/core/domain"

func MapAmenity(existing, update *domain.Amenity) {
	if existing == nil || update == nil {
		return
	}

	existing.TenantID = update.TenantID
	existing.Code = update.Code
	existing.Name = update.Name
	existing.Category = update.Category
	existing.Icon = update.Icon
	existing.IsGlobal = update.IsGlobal
}