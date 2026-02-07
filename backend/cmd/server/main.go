package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/chayutK/hotel-property-service/internal/adapter"
	"github.com/chayutK/hotel-property-service/internal/config"
	"github.com/chayutK/hotel-property-service/internal/infra/database"
	"github.com/chayutK/hotel-property-service/internal/service"
	"github.com/chayutK/hotel-property-service/internal/transport/http"
	"github.com/chayutK/hotel-property-service/internal/transport/http/handler"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	logger := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logger))
	validate := validator.New()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("[MAIN]", "message", "error while loading config", "error", err.Error())
		panic(err)
	}

	db, err := database.New(cfg.Database.DSN, cfg.Database.Migration, cfg.Database.Seeding)
	if err != nil {
		slog.Error("[MAIN]", "message", "error while connecting database", "error", err.Error())
		panic(err)
	}

	app := echo.New()
	app.Use(
		middleware.RequestLogger(),
		middleware.Recover(),
	)

	hotelRepo := adapter.NewHotelRepository(db)
	roomRepo := adapter.NewRoomRepository(db)

	hotelSvc := service.NewHotelService(hotelRepo)
	roomSvc := service.NewRoomService(roomRepo)
	priceSvc := service.NewPricingService(hotelRepo, roomRepo)

	hotelHandler := handler.NewHotelHandler(hotelSvc, validate)
	roomHandler := handler.NewRoomHandler(roomSvc, validate)
	pricingHandler := handler.NewPricingHandler(priceSvc, validate)

	http.RegisterRoutes(app, hotelHandler, roomHandler, pricingHandler)

	// Swagger documentation route
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	app.Start(fmt.Sprintf(":%d", cfg.Server.Port))
}
