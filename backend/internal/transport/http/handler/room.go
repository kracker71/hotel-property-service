package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/chayutK/hotel-property-service/internal/service"
	"github.com/chayutK/hotel-property-service/internal/transport/http/dto/mapperdto"
	"github.com/chayutK/hotel-property-service/internal/transport/http/dto/roomdto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	roomService *service.RoomService
	validate    *validator.Validate
}

func NewRoomHandler(roomService *service.RoomService, validate *validator.Validate) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
		validate:    validate,
	}
}

func (h *RoomHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/hotels/:hotelID/rooms", h.GetRooms)
	g.GET("/hotels/:hotelID/rooms/:roomID", h.GetRoomByID)
}

// GetRooms godoc
// @Summary List rooms by hotel
// @Description Get rooms for a given hotel
// @Tags rooms
// @Produce json
// @Param hotelID path string true "Hotel ID"
// @Success 200 {object} roomdto.InquiryRoomsResponse
// @Router /hotels/{hotelID}/rooms [get]
func (h *RoomHandler) GetRooms(c echo.Context) error {
	var (
		req  roomdto.InquiryRoomsRequest
		resp roomdto.InquiryRoomsResponse
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

	rooms, err := h.roomService.GetRoomsByHotelID(ctx, req.HotelID)
	if err != nil {
		return err
	}

	resp.Rooms = mapperdto.ToRoomsDTO(rooms)
	return c.JSON(200, &resp)
}

// GetRoomByID godoc
// @Summary Get room by id
// @Description Get room details by hotel id and room id
// @Tags rooms
// @Produce json
// @Param hotelID path string true "Hotel ID"
// @Param roomID path string true "Room ID"
// @Success 200 {object} roomdto.InquiryRoomResponse
// @Router /hotels/{hotelID}/rooms/{roomID} [get]
func (h *RoomHandler) GetRoomByID(c echo.Context) error {
	var (
		req  roomdto.InquiryRoomRequest
		resp roomdto.InquiryRoomResponse
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

	room, err := h.roomService.GetRoomByRoomID(ctx, req.HotelID, req.RoomID)
	if err != nil {
		return err
	}

	resp.Room = *mapperdto.ToRoomDTO(room)
	return c.JSON(200, &resp)
}
