package database

import (
	"fmt"
	"log/slog"

	"github.com/chayutK/hotel-property-service/internal/adapter/entity"
	"github.com/chayutK/hotel-property-service/internal/constants/cancellationpolicy"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RunSeeder seeds the database with sample data for hotels, facilities, rooms and benefits.
// It runs inside a transaction and is safe to call after migrations.
func RunSeeder(db *gorm.DB) error {
	slog.Info("[SEED]", "message", "Starting database seeder")

	// If there are already hotels in the database, skip seeding to avoid duplicates.
	var existing int64
	if err := db.Model(&entity.Hotel{}).Count(&existing).Error; err != nil {
		return err
	}
	if existing > 0 {
		slog.Info("[SEED]", "message", "Skipping seeding; data already exists", "hotels", existing)
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {
		// create multiple sample hotels
		for hi := 1; hi <= 3; hi++ {
			hotelID := uuid.NewString()
			hotel := entity.Hotel{
				HotelID:  hotelID,
				Name:     fmt.Sprintf("Demo Hotel %d", hi),
				Address:  fmt.Sprintf("%d Demo Street", hi),
				IsActive: true,
			}
			if err := tx.Create(&hotel).Error; err != nil {
				return err
			}

			// facilities per hotel
			facilities := []entity.Facility{}
			for fi := 1; fi <= 3; fi++ {
				facilities = append(facilities, entity.Facility{
					FacilityID:  uuid.NewString(),
					HotelID:     hotelID,
					Name:        fmt.Sprintf("Facility %d-%d", hi, fi),
					Description: "Auto-generated facility",
					IsActive:    true,
				})
			}
			if err := tx.Create(&facilities).Error; err != nil {
				return err
			}

			// rooms per hotel
			rooms := []entity.Room{}
			for ri := 1; ri <= 4; ri++ {
				physID := uuid.NewString()
				rooms = append(rooms, entity.Room{
					RoomID:             uuid.NewString(),
					PhysicalRoomID:     physID,
					HotelID:            hotelID,
					Name:               fmt.Sprintf("Room %d-%d", hi, ri),
					Description:        "Auto-generated room",
					Type:               []string{"standard", "deluxe", "suite"}[ri%3],
					BasePrice:          int64(5000 * ri * hi),
					Currency:           "THB",
					CancellationPolicy: []string{cancellationpolicy.FreeCancellation, cancellationpolicy.NonRefundable}[ri%2],
					IsActive:           true,
				})
			}
			if err := tx.Create(&rooms).Error; err != nil {
				return err
			}

			// benefits for each room (2 benefits per room)
			benefits := []entity.Benefit{}
			for _, r := range rooms {
				benefits = append(benefits, entity.Benefit{
					BenefitID:      uuid.NewString(),
					PhysicalRoomID: r.PhysicalRoomID,
					Name:           "Breakfast",
					Description:    "Complimentary breakfast",
					IsActive:       true,
				})
				benefits = append(benefits, entity.Benefit{
					BenefitID:      uuid.NewString(),
					PhysicalRoomID: r.PhysicalRoomID,
					Name:           "WiFi",
					Description:    "Free high-speed internet",
					IsActive:       true,
				})
			}
			if len(benefits) > 0 {
				if err := tx.Create(&benefits).Error; err != nil {
					return err
				}
			}
		}

		slog.Info("[SEED]", "message", "Database seeding completed")
		return nil
	})
}
