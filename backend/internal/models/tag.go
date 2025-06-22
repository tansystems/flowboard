package models

import "gorm.io/gorm"

type Tag struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Deals     []Deal         `gorm:"many2many:deal_tags;"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
