package adapter

import (
	"context"
	"log/slog"

	"github.com/chayutK/hotel-property-service/internal/adapter/entity"
	"github.com/chayutK/hotel-property-service/internal/adapter/mapper"
	"github.com/chayutK/hotel-property-service/internal/domain"
	"github.com/chayutK/hotel-property-service/internal/port"
	"gorm.io/gorm"
)

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) port.HotelPort {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) FindAll(ctx context.Context) ([]domain.Hotel, error) {
	var gormHotels []entity.Hotel

	if err := r.db.WithContext(ctx).Preload("Facility").Find(&gormHotels).Error; err != nil {
		slog.Error("[ADAPTER]", "message", "error while inquiry hotels", "error", err.Error())
		return nil, err
	}

	domainHotels := mapper.ToDomainHotels(gormHotels)
	return domainHotels, nil
}

func (r *hotelRepository) FindByID(ctx context.Context, id string) (*domain.Hotel, error) {
	var gormHotel entity.Hotel

	if err := r.db.WithContext(ctx).Preload("Facility").First(&gormHotel, "hotel_id = ?", id).Error; err != nil {
		slog.Error("[ADAPTER]", "message", "error while inquiry hotel by id", "hotel_id", id, "error", err.Error())
		return nil, err
	}

	domainHotel := mapper.ToDomainHotel(&gormHotel)
	return domainHotel, nil
}
