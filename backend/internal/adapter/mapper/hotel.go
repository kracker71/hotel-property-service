package mapper

import (
	"github.com/chayutK/hotel-property-service/internal/adapter/entity"
	"github.com/chayutK/hotel-property-service/internal/domain"
)

func ToDomainHotels(es []entity.Hotel) []domain.Hotel {
	domains := make([]domain.Hotel, len(es))
	for i, e := range es {
		domains[i] = *ToDomainHotel(&e)
	}
	return domains
}

func ToDomainHotel(e *entity.Hotel) *domain.Hotel {
	if e == nil {
		return nil
	}

	facilities := make([]domain.Facility, len(e.Facility))
	for i, f := range e.Facility {
		facilities[i] = *ToDomainFacility(&f)
	}

	return &domain.Hotel{
		ID:       e.HotelID,
		Name:     e.Name,
		Address:  e.Address,
		IsActive: e.IsActive,
		Facility: facilities,
	}
}

func ToDomainFacility(e *entity.Facility) *domain.Facility {
	if e == nil {
		return nil
	}

	return &domain.Facility{
		ID:          e.FacilityID,
		HotelID:     e.HotelID,
		Name:        e.Name,
		Description: e.Description,
		IsActive:    e.IsActive,
	}
}
