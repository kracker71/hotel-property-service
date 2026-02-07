package handler

import (
	"context"
	"time"

	"github.com/chayutK/hotel-property-service/internal/service"
	hoteldto "github.com/chayutK/hotel-property-service/internal/transport/http/dto/hotel"
	mapperdto "github.com/chayutK/hotel-property-service/internal/transport/http/dto/mapper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HotelHandler struct {
	hotelService *service.HotelService
	validate     *validator.Validate
}

func NewHotelHandler(hotelService *service.HotelService, validate *validator.Validate) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
		validate:     validate,
	}
}

func (h *HotelHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/hotels", h.GetAllHotels)
	g.GET("/hotel/:hotel_id", h.GetHotelByID)
}

func (h *HotelHandler) GetAllHotels(c echo.Context) error {
	var (
		resp hoteldto.InquiryHotelsResponse
	)
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	hotels, err := h.hotelService.GetAllHotels(ctx)
	if err != nil {
		return err
	}

	resp.Hotels = mapperdto.ToHotelsDTO(hotels)
	return c.JSON(200, &resp)
}

func (h *HotelHandler) GetHotelByID(c echo.Context) error {
	var (
		req  hoteldto.InquiryHotelRequest
		resp hoteldto.InquiryHotelResponse
	)

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	hotel, err := h.hotelService.GetHotelByID(ctx, req.HotelID)
	if err != nil {
		return err
	}

	resp.Hotel = *mapperdto.ToHotelDTO(hotel)
	return c.JSON(200, &resp)
}
