package handler

import (
	"context"
	"time"

	"github.com/chayutK/hotel-property-service/internal/service"
	pricingdto "github.com/chayutK/hotel-property-service/internal/transport/http/dto/pricing"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PricingHandler struct {
	pricingService *service.PricingService
	validate       *validator.Validate
}

func NewPricingHandler(pricingService *service.PricingService, validate *validator.Validate) *PricingHandler {
	return &PricingHandler{
		pricingService: pricingService,
		validate:       validate,
	}
}

func (h *PricingHandler) RegisterRoutes(g *echo.Group) {
	g.POST("/price", h.CalculateRoomPrice)
}

func (h *PricingHandler) CalculateRoomPrice(c echo.Context) error {
	var (
		req  pricingdto.CalculatePricingRequest
		resp pricingdto.CalculatePricingResponse
	)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	price, err := h.pricingService.CalculateRoomPrice(ctx, req.HotelID, req.RoomID, req.Nights)
	if err != nil {
		return err
	}

	resp.TotalPrice = price
	return c.JSON(200, &resp)
}
