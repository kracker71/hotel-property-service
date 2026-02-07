package mapper

import (
	"github.com/chayutK/hotel-property-service/internal/adapter/entity"
	"github.com/chayutK/hotel-property-service/internal/domain"
)

func ToDomainBenefits(es []entity.Benefit) []domain.Benefit {
	domains := make([]domain.Benefit, len(es))
	for i, e := range es {
		domains[i] = *ToDomainBenefit(&e)
	}
	return domains
}

func ToDomainBenefit(e *entity.Benefit) *domain.Benefit {
	if e == nil {
		return nil
	}

	return &domain.Benefit{
		ID:             e.BenefitID,
		PhysicalRoomID: e.PhysicalRoomID,
		Name:           e.Name,
		Description:    e.Description,
		IsActive:       e.IsActive,
	}
}

func ToDomainRooms(es []entity.Room) []domain.Room {
	domains := make([]domain.Room, len(es))
	for i, e := range es {
		domains[i] = *ToDomainRoom(&e)
	}
	return domains
}

func ToDomainRoom(e *entity.Room) *domain.Room {
	if e == nil {
		return nil
	}

	benefits := make([]domain.Benefit, len(e.Benefit))
	for i, b := range e.Benefit {
		benefits[i] = *ToDomainBenefit(&b)
	}

	return &domain.Room{
		ID:                 e.RoomID,
		PhysicalRoomID:     e.PhysicalRoomID,
		HotelID:            e.HotelID,
		Name:               e.Name,
		Description:        e.Description,
		BasePrice:          float64(e.BasePrice),
		Type:               e.Type,
		Currency:           e.Currency,
		CancellationPolicy: e.CancellationPolicy,
		Benefit:            benefits,
		IsActive:           e.IsActive,
	}
}
