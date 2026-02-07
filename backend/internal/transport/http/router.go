package http

import (
	"github.com/chayutK/hotel-property-service/internal/transport/http/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	hotelHandler *handler.HotelHandler,
	roomHandler *handler.RoomHandler,
	pricingHandler *handler.PricingHandler,
) {
	apiGroup := e.Group("/api/v1")

	hotelHandler.RegisterRoutes(apiGroup)
	roomHandler.RegisterRoutes(apiGroup)
	pricingHandler.RegisterRoutes(apiGroup)
}
