package entity

type Facility struct {
	FacilityID  string `gorm:"column:facility_id;primaryKey"`
	HotelID     string `gorm:"column:hotel_id;index"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	IsActive    bool   `gorm:"column:is_active"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime"`
}
