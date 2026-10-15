package mapper

import "github.com/yohagos/multi-content-management/internal/core/domain"

func MapBuilding(existing, update *domain.Building) {
	if existing == nil || update == nil {
		return
	}

	existing.TenantID = update.TenantID
	existing.Name = update.Name
	existing.Slug = update.Slug
	existing.Street = update.Street
	existing.HouseNumber = update.HouseNumber
	existing.PostalCode = update.PostalCode
	existing.City = update.City
	existing.Country = update.Country
	existing.Latitude = update.Latitude
	existing.Longitude = update.Longitude
	existing.YearBuild = update.YearBuild
	existing.TotalFloors = update.TotalFloors
	existing.TotalUnits = update.TotalUnits
	existing.Description = update.Description
	existing.Status = update.Status
	existing.Published = update.Published
	existing.PublishedAt = &update.UpdatedAt
}