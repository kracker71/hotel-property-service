package port

import (
	"context"

	"github.com/chayutK/hotel-property-service/internal/domain"
)

type HotelPort interface {
	FindAll(ctx context.Context) ([]domain.Hotel, error)
	FindByID(ctx context.Context, id string) (*domain.Hotel, error)
}
