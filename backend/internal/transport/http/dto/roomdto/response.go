package roomdto

type InquiryRoomsResponse struct {
	Rooms []RoomDTO `json:"rooms"`
}

type InquiryRoomResponse struct {
	Room RoomDTO `json:"room"`
}
