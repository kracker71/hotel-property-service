package roomdto

type InquiryRoomsRequest struct {
	HotelID string `param:"hotelID" validate:"required,uuid4"`
}

type InquiryRoomRequest struct {
	HotelID string `param:"hotelID" validate:"required,uuid4"`
	RoomID  string `param:"roomID" validate:"required,uuid4"`
}
