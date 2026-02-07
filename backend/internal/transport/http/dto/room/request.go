package roomdto

type InquiryRoomsRequest struct {
	HotelID string `param:"hotel_id" validate:"required,uuid4"`
}

type InquiryRoomRequest struct {
	HotelID string `param:"hotel_id" validate:"required,uuid4"`
	RoomID  string `param:"room_id" validate:"required,uuid4"`
}

