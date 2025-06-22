package models

import "gorm.io/gorm"

type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	DealID    uint           `json:"deal_id"`
	Deal      Deal           `gorm:"foreignKey:DealID"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignKey:UserID"`
	Content   string         `json:"content"`
	CreatedAt int64          `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
