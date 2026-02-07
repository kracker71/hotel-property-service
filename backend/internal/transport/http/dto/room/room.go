package roomdto

type RoomDTO struct {
	RoomID             string       `json:"roomID"`
	HotelID            string       `json:"hotelID"`
	Name               string       `json:"name"`
	Description        string       `json:"description"`
	Type               string       `json:"type"`
	BasePrice          float64      `json:"basePrice"`
	Currency           string       `json:"currency"`
	Benefit            []BenefitDTO `json:"benefit"`
	CancellationPolicy string       `json:"cancellationPolicy"`
}

type BenefitDTO struct {
	BenefitID      string `json:"benefitID"`
	PhysicalRoomID string `json:"physicalRoomID"`
	Description    string `json:"description"`
}
