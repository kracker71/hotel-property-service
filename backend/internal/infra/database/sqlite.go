package database

import (
	"log/slog"
	"os"

	"github.com/chayutK/hotel-property-service/internal/adapter/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(dsn string, migration, seeding bool) (*gorm.DB, error) {
	err := ensureDir("./data")
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("[INFRA]", "message", "Error while connecting to database", "error", err.Error())
		return nil, err
	}

	if migration {
		if err := Migrate(db); err != nil {
			slog.Error("[INFRA]", "message", "Error while migrating database", "error", err.Error())
			return nil, err
		}
	}

	// Seeding can be implemented here if needed
	if seeding {
		if err := RunSeeder(db); err != nil {
			slog.Error("[INFRA]", "message", "Error while seeding database", "error", err.Error())
			return nil, err
		}
	}

	slog.Info("[INFRA]", "message", "Connecting to database successfully!!")
	return db, nil
}

func ensureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// Migrate runs database migrations
func Migrate(db *gorm.DB) error {
	slog.Info("[INFRA]", "message", "Running database migrations...")

	err := db.AutoMigrate(
		&entity.Facility{},
		&entity.Benefit{},
		&entity.Hotel{},
		&entity.Room{},
	)

	if err != nil {
		slog.Error("[INFRA]", "message", "Failed to migrate database", "error", err.Error())
		return err
	}

	slog.Info("[INFRA]", "message", "Database migrations completed successfully!")
	return nil
}
