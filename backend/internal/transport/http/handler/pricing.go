package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/chayutK/hotel-property-service/internal/service"
	"github.com/chayutK/hotel-property-service/internal/transport/http/dto/pricingdto"
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

// CalculateRoomPrice godoc
// @Summary Calculate room price
// @Description Calculate total price for requested hotel room and nights
// @Tags pricing
// @Accept json
// @Produce json
// @Param request body pricingdto.CalculatePricingRequest true "Pricing request"
// @Success 200 {object} pricingdto.CalculatePricingResponse
// @Router /price [post]
func (h *PricingHandler) CalculateRoomPrice(c echo.Context) error {
	var (
		req  pricingdto.CalculatePricingRequest
		resp pricingdto.CalculatePricingResponse
	)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := c.Bind(&req); err != nil {
		slog.Error("[HANDLER]", "message", "error binding request", "error", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request"})
	}

	if err := h.validate.Struct(&req); err != nil {
		slog.Error("[HANDLER]", "message", "error validating request", "error", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request"})
	}

	price, err := h.pricingService.CalculateRoomPrice(ctx, req.HotelID, req.RoomID, req.Nights)
	if err != nil {
		return err
	}

	resp.TotalPrice = price
	return c.JSON(200, &resp)
}
