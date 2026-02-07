package hoteldto

type HotelDTO struct {
	HotelID  string        `json:"hotel_id"`
	Name     string        `json:"name"`
	Address  string        `json:"address"`
	Facility []FacilityDTO `json:"facility"`
}

type FacilityDTO struct {
	FacilityID  string `json:"facility_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
