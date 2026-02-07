package domain

import "github.com/chayutK/hotel-property-service/internal/constants/cancellationpolicy"

type Room struct {
	ID                 string
	PhysicalRoomID     string
	HotelID            string
	Name               string
	Description        string
	Type               string
	BasePrice          float64
	Currency           string
	CancellationPolicy string
	Benefit            []Benefit
	IsActive           bool
}

func (r *Room) CalculatePrice(nights int) float64 {
	if r.CancellationPolicy == cancellationpolicy.FreeCancellation {
		return r.BasePrice * float64(nights) * 1.2 // 20% surcharge for free cancellation
	}
	return r.BasePrice * float64(nights)
}
