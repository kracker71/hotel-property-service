package database

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

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
	// add some randomness so seeded values vary across runs
	rand.Seed(time.Now().UnixNano())

	return db.Transaction(func(tx *gorm.DB) error {
		// curated sample data to feel more realistic
		hotelNames := []string{
			"The Riverside Grand",
			"Sukhumvit Boutique Hotel",
			"Bangkok Skyline Hotel",
			"Chiang Mai Garden Inn",
			"Phuket Sea Breeze Resort",
		}
		streets := []string{"Sukhumvit Rd", "Silom Rd", "Rama IV Rd", "Nimmanhemin Rd", "Patong Beach Rd"}
		cities := []string{"Bangkok", "Chiang Mai", "Phuket", "Pattaya", "Hua Hin"}

		facilityCatalog := []struct{ Name, Desc string }{
			{"Swimming Pool", "Outdoor heated pool"},
			{"Fitness Center", "24/7 gym with modern equipment"},
			{"Spa", "Full service spa and massage"},
			{"Rooftop Bar", "City views and signature cocktails"},
			{"Airport Shuttle", "Scheduled airport transfers"},
			{"Free Parking", "Complimentary on-site parking"},
			{"Kids Club", "Supervised activities for children"},
		}

		roomTemplates := []struct{ Name, Type string }{
			{"Superior King", "standard"},
			{"Deluxe Twin", "deluxe"},
			{"Executive Suite", "suite"},
			{"Family Room", "family"},
		}

		benefitPool := []struct{ Name, Desc string }{
			{"Breakfast", "Complimentary breakfast buffet"},
			{"WiFi", "Unlimited high-speed internet"},
			{"Late Checkout", "Guaranteed late checkout until 2PM"},
			{"Welcome Drink", "One complimentary welcome drink"},
			{"Airport Pickup", "One-way airport pickup"},
			{"City Tour Discount", "10% off partnered city tours"},
			{"Mini Bar Credit", "Small minibar credit included"},
		}

		// create a few hotels
		for i := 0; i < 3; i++ {
			hotelID := uuid.NewString()
			name := hotelNames[rand.Intn(len(hotelNames))]
			street := streets[rand.Intn(len(streets))]
			city := cities[rand.Intn(len(cities))]
			number := 100 + rand.Intn(900)
			addr := fmt.Sprintf("%d %s, %s", number, street, city)

			hotel := entity.Hotel{
				HotelID:  hotelID,
				Name:     name,
				Address:  addr,
				IsActive: true,
			}
			if err := tx.Create(&hotel).Error; err != nil {
				return err
			}

			// attach 3-5 unique facilities chosen from catalog (no duplicates)
			facilities := []entity.Facility{}
			fcount := 3 + rand.Intn(3)
			if fcount > len(facilityCatalog) {
				fcount = len(facilityCatalog)
			}
			picks := map[int]struct{}{}
			for len(picks) < fcount {
				idx := rand.Intn(len(facilityCatalog))
				if _, ok := picks[idx]; ok {
					continue
				}
				picks[idx] = struct{}{}
				pick := facilityCatalog[idx]
				facilities = append(facilities, entity.Facility{
					FacilityID:  uuid.NewString(),
					HotelID:     hotelID,
					Name:        pick.Name,
					Description: pick.Desc,
					IsActive:    true,
				})
			}
			if err := tx.Create(&facilities).Error; err != nil {
				return err
			}

			// create 3-6 rooms with varied names/prices
			rooms := []entity.Room{}
			rcount := 3 + rand.Intn(4)
			for ri := 0; ri < rcount; ri++ {
				physID := uuid.NewString()
				tmpl := roomTemplates[rand.Intn(len(roomTemplates))]
				// base price varies by hotel index and room type
				base := 1500 + rand.Intn(6000) + i*1000 + ri*500
				rooms = append(rooms, entity.Room{
					RoomID:             uuid.NewString(),
					PhysicalRoomID:     physID,
					HotelID:            hotelID,
					Name:               tmpl.Name,
					Description:        fmt.Sprintf("%s with modern amenities", tmpl.Name),
					Type:               tmpl.Type,
					BasePrice:          int64(base),
					Currency:           "THB",
					CancellationPolicy: []string{cancellationpolicy.FreeCancellation, cancellationpolicy.NonRefundable}[rand.Intn(2)],
					IsActive:           true,
				})
			}
			if err := tx.Create(&rooms).Error; err != nil {
				return err
			}

			// assign 1-3 benefits per room, sampled from pool
			benefits := []entity.Benefit{}
			for _, r := range rooms {
				bcount := 1 + rand.Intn(3)
				// pick bcount unique benefits
				picks := map[int]struct{}{}
				for b := 0; b < bcount; b++ {
					for {
						idx := rand.Intn(len(benefitPool))
						if _, ok := picks[idx]; ok {
							continue
						}
						picks[idx] = struct{}{}
						bp := benefitPool[idx]
						benefits = append(benefits, entity.Benefit{
							BenefitID:      uuid.NewString(),
							PhysicalRoomID: r.PhysicalRoomID,
							Name:           bp.Name,
							Description:    bp.Desc,
							IsActive:       true,
						})
						break
					}
				}
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
