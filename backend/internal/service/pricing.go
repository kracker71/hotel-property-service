package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/chayutK/hotel-property-service/internal/port"
)

type PricingService struct {
	hotelRepository port.HotelPort
	roomRepository  port.RoomPort
}

func NewPricingService(hotelRepository port.HotelPort, roomRepository port.RoomPort) *PricingService {
	return &PricingService{
		hotelRepository: hotelRepository,
		roomRepository:  roomRepository,
	}
}

func (s *PricingService) CalculateRoomPrice(ctx context.Context, hotelID, roomID string, nights int) (float64, error) {
	room, err := s.roomRepository.FindByRoomID(ctx, roomID)
	if err != nil {
		return 0, err
	}

	if room.HotelID != hotelID {
		slog.Error("[SERVICE]", "message", fmt.Sprintf("hotelID does not match with room, room.HotelID:%s, hotelID:%s", room.HotelID, hotelID))
		return 0, fmt.Errorf("hotelID does not match with room")
	}

	// Example pricing logic: base price multiplied by number of nights
	totalPrice := room.CalculatePrice(nights)
	return totalPrice, nil
}
