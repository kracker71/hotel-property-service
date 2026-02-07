package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/chayutK/hotel-property-service/internal/domain"
	"github.com/chayutK/hotel-property-service/internal/port"
)

type RoomService struct {
	roomRepository port.RoomPort
}

func NewRoomService(roomRepository port.RoomPort) *RoomService {
	return &RoomService{
		roomRepository: roomRepository,
	}
}

func (s *RoomService) GetRoomsByHotelID(ctx context.Context, hotelID string) ([]domain.Room, error) {
	rooms, err := s.roomRepository.FindByHotelID(ctx, hotelID)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *RoomService) GetRoomByRoomID(ctx context.Context, hotelID, roomID string) (*domain.Room, error) {
	room, err := s.roomRepository.FindByRoomID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	if room.HotelID != hotelID {
		slog.Error("[SERVICE]", "message", fmt.Sprintf("hotelID does not match with room, room.HotelID:%s, hotelID:%s", room.HotelID, hotelID))
		return nil, fmt.Errorf("hotelID does not match with room")
	}

	return room, nil
}
