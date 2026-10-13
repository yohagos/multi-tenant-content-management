package mapper

import "github.com/yohagos/multi-content-management/internal/core/domain"

func MapUnit(existing, updated *domain.Unit) {
	if existing == nil || updated == nil {
		return
	}

	existing.TenantID = updated.TenantID
	existing.BuildingID = updated.BuildingID
	existing.UnitNumber = updated.UnitNumber
	existing.Floor = updated.Floor
	existing.Rooms = updated.Rooms
	existing.Bedrooms = updated.Bedrooms
	existing.Bathrooms = updated.Bathrooms
	existing.LivingAreaSqm = updated.LivingAreaSqm
	existing.TotalAreaSqm = updated.TotalAreaSqm
	existing.BalconyAreaSqm = updated.BalconyAreaSqm
	existing.CeilingHeightM = updated.CeilingHeightM
	existing.ColdRent = updated.ColdRent
	existing.WarmRent = updated.WarmRent
	existing.Deposit = updated.Deposit
	existing.AvailableFrom = updated.AvailableFrom
	existing.Status = updated.Status
	existing.Published = updated.Published
	existing.PublishedAt = updated.PublishedAt
	existing.Title = updated.Title
	existing.Description = updated.Description
	existing.UpdatedAt = updated.UpdatedAt
	existing.DeletedAt = updated.DeletedAt
}
