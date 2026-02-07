package port

import (
	"context"

	"github.com/chayutK/hotel-property-service/internal/domain"
)

type RoomPort interface {
	FindByHotelID(ctx context.Context, hotelID string) ([]domain.Room, error)
	FindByRoomID(ctx context.Context, roomID string) (*domain.Room, error)
}
