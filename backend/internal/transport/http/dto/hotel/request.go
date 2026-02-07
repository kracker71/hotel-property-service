package hoteldto

type InquiryHotelRequest struct {
	HotelID string `param:"hotel_id" validate:"required,uuid4"`
}
