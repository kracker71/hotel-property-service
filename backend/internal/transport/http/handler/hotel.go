package handler

import (
	"context"
	"log/slog"
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

// GetAllHotels godoc
// @Summary List hotels
// @Description Get a list of hotels
// @Tags hotels
// @Produce json
// @Success 200 {object} hoteldto.InquiryHotelsResponse
// @Router /hotels [get]
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

// GetHotelByID godoc
// @Summary Get hotel by id
// @Description Get hotel details by hotel id
// @Tags hotels
// @Produce json
// @Param hotel_id path string true "Hotel ID"
// @Success 200 {object} hoteldto.InquiryHotelResponse
// @Router /hotel/{hotel_id} [get]
func (h *HotelHandler) GetHotelByID(c echo.Context) error {
	var (
		req  hoteldto.InquiryHotelRequest
		resp hoteldto.InquiryHotelResponse
	)

	if err := c.Bind(&req); err != nil {
		slog.Error("[HANDLER]", "message", "error binding request", "error", err.Error())
		return err
	}

	if err := h.validate.Struct(&req); err != nil {
		slog.Error("[HANDLER]", "message", "error validating request", "error", err.Error())
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
