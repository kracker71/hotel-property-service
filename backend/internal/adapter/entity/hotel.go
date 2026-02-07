package entity

type Hotel struct {
	HotelID   string     `gorm:"column:hotel_id;primaryKey"`
	Name      string     `gorm:"column:name"`
	Address   string     `gorm:"column:address"`
	Facility  []Facility `gorm:"foreignKey:HotelID;references:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsActive  bool       `gorm:"column:is_active"`
	CreatedAt int64      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt int64      `gorm:"column:updated_at;autoUpdateTime"`
}
