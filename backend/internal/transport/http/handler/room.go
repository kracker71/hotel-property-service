package handler

import (
	"context"
	"time"

	"github.com/chayutK/hotel-property-service/internal/service"
	mapperdto "github.com/chayutK/hotel-property-service/internal/transport/http/dto/mapper"
	roomdto "github.com/chayutK/hotel-property-service/internal/transport/http/dto/room"
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

func (h *RoomHandler) GetRooms(c echo.Context) error {
	var (
		req  roomdto.InquiryRoomsRequest
		resp roomdto.InquiryRoomsResponse
	)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	rooms, err := h.roomService.GetRoomsByHotelID(ctx, req.HotelID)
	if err != nil {
		return err
	}

	resp.Rooms = mapperdto.ToRoomsDTO(rooms)
	return c.JSON(200, &resp)
}

func (h *RoomHandler) GetRoomByID(c echo.Context) error {
	var (
		req  roomdto.InquiryRoomRequest
		resp roomdto.InquiryRoomResponse
	)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	room, err := h.roomService.GetRoomByRoomID(ctx, req.HotelID, req.RoomID)
	if err != nil {
		return err
	}

	resp.Room = *mapperdto.ToRoomDTO(room)
	return c.JSON(200, &resp)
}
