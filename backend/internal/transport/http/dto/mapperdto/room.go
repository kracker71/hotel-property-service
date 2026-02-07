package mapperdto

import (
	"github.com/chayutK/hotel-property-service/internal/domain"
	"github.com/chayutK/hotel-property-service/internal/transport/http/dto/roomdto"
)

func ToRoomsDTO(rooms []domain.Room) []roomdto.RoomDTO {
	roomDTOs := make([]roomdto.RoomDTO, len(rooms))
	for i, room := range rooms {
		roomDTOs[i] = *ToRoomDTO(&room)
	}
	return roomDTOs
}

func ToRoomDTO(room *domain.Room) *roomdto.RoomDTO {
	if room == nil {
		return nil
	}

	benefits := make([]roomdto.BenefitDTO, len(room.Benefit))
	for i, benefit := range room.Benefit {
		benefits[i] = *ToBenefitDTO(&benefit)
	}

	return &roomdto.RoomDTO{
		RoomID:             room.ID,
		HotelID:            room.HotelID,
		Name:               room.Name,
		Description:        room.Description,
		Type:               room.Type,
		BasePrice:          room.BasePrice,
		Currency:           room.Currency,
		Benefit:            benefits,
		CancellationPolicy: room.CancellationPolicy,
	}
}

func ToBenefitDTO(benefit *domain.Benefit) *roomdto.BenefitDTO {
	return &roomdto.BenefitDTO{
		BenefitID:      benefit.ID,
		PhysicalRoomID: benefit.PhysicalRoomID,
		Description:    benefit.Description,
	}

}
