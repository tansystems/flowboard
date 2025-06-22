package models

import "gorm.io/gorm"

type Deal struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CustomerID  uint           `json:"customer_id"`
	Customer    Customer       `gorm:"foreignKey:CustomerID"`
	StatusID    uint           `json:"status_id"`
	Status      Status         `gorm:"foreignKey:StatusID"`
	CreatedAt   int64          `json:"created_at"`
	UpdatedAt   int64          `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
