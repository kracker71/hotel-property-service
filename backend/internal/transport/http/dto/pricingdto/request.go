package pricingdto

type CalculatePricingRequest struct {
	HotelID string `json:"hotelID" validate:"required,uuid4"`
	RoomID  string `json:"roomID" validate:"required,uuid4"`
	Nights  int    `json:"nights" validate:"required,min=1"`
}
