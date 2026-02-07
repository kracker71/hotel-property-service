package entity

type Benefit struct {
	BenefitID      string `gorm:"column:benefit_id;primaryKey"`
	PhysicalRoomID string `gorm:"column:physical_room_id;index"`
	Name           string `gorm:"column:name"`
	Description    string `gorm:"column:description"`
	IsActive       bool   `gorm:"column:is_active"`
	CreatedAt      int64  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      int64  `gorm:"column:updated_at;autoUpdateTime"`
}
