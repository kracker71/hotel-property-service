package service

import (
	"context"

	"github.com/chayutK/hotel-property-service/internal/domain"
	"github.com/chayutK/hotel-property-service/internal/port"
)

type HotelService struct {
	hotelRepository port.HotelPort
}

func NewHotelService(hotelRepository port.HotelPort) *HotelService {
	return &HotelService{
		hotelRepository: hotelRepository,
	}
}

func (s *HotelService) GetAllHotels(ctx context.Context) ([]domain.Hotel, error) {
	hotels, err := s.hotelRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (s *HotelService) GetHotelByID(ctx context.Context, hotelID string) (*domain.Hotel, error) {
	hotel, err := s.hotelRepository.FindByID(ctx, hotelID)
	if err != nil {
		return nil, err
	}

	return hotel, nil
}
