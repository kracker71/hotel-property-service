package entity

type Room struct {
	RoomID             string    `gorm:"column:room_id;primaryKey"`
	PhysicalRoomID     string    `gorm:"column:physical_room_id;index"`
	HotelID            string    `gorm:"column:hotel_id;index"`
	Name               string    `gorm:"column:name"`
	Description        string    `gorm:"column:description"`
	Type               string    `gorm:"column:type"`
	BasePrice          int64     `gorm:"column:base_price"`
	Currency           string    `gorm:"column:currency"`
	CancellationPolicy string    `gorm:"column:cancellation_policy"`
	Benefit            []Benefit `gorm:"foreignKey:PhysicalRoomID;references:PhysicalRoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsActive           bool      `gorm:"column:is_active"`
	CreatedAt          int64     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          int64     `gorm:"column:updated_at;autoUpdateTime"`
}
