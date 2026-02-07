package hoteldto

type InquiryHotelResponse struct {
	Hotel HotelDTO `json:"hotel"`
}

type InquiryHotelsResponse struct {
	Hotels []HotelDTO `json:"hotels"`
}
