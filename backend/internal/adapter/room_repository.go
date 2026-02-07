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

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) port.RoomPort {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) FindByHotelID(ctx context.Context, hotelID string) ([]domain.Room, error) {
	var gormRooms []entity.Room

	if err := r.db.WithContext(ctx).Preload("Benefit").Where("hotel_id = ?", hotelID).Find(&gormRooms).Error; err != nil {
		slog.Error("[ADAPTER]", "message", "error while inquiry rooms by hotel id", "hotel_id", hotelID, "error", err.Error())
		return nil, err
	}

	domainRooms := mapper.ToDomainRooms(gormRooms)
	return domainRooms, nil
}

func (r *RoomRepository) FindByRoomID(ctx context.Context, roomID string) (*domain.Room, error) {
	var gormRoom entity.Room

	if err := r.db.WithContext(ctx).Preload("Benefit").First(&gormRoom, "room_id = ?", roomID).Error; err != nil {
		slog.Error("[ADAPTER]", "message", "error while inquiry room by room id", "room_id", roomID, "error", err.Error())
		return nil, err
	}

	domainRoom := mapper.ToDomainRoom(&gormRoom)
	return domainRoom, nil
}
