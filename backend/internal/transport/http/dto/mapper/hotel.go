package mapperdto

import (
	"github.com/chayutK/hotel-property-service/internal/domain"
	hoteldto "github.com/chayutK/hotel-property-service/internal/transport/http/dto/hotel"
)

func ToHotelsDTO(hotels []domain.Hotel) []hoteldto.HotelDTO {
	hotelDTOs := make([]hoteldto.HotelDTO, len(hotels))
	for i, hotel := range hotels {
		hotelDTOs[i] = *ToHotelDTO(&hotel)
	}
	return hotelDTOs
}

func ToHotelDTO(hotel *domain.Hotel) *hoteldto.HotelDTO {
	if hotel == nil {
		return nil
	}

	facilities := make([]hoteldto.FacilityDTO, len(hotel.Facility))
	for i, facility := range hotel.Facility {
		facilities[i] = *ToFacilityDTO(&facility)
	}

	return &hoteldto.HotelDTO{
		HotelID:  hotel.ID,
		Name:     hotel.Name,
		Address:  hotel.Address,
		Facility: facilities,
	}
}

func ToFacilityDTO(facility *domain.Facility) *hoteldto.FacilityDTO {
	if facility == nil {
		return nil
	}

	return &hoteldto.FacilityDTO{
		FacilityID:  facility.ID,
		Name:        facility.Name,
		Description: facility.Description,
	}
}
