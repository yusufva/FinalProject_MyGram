package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `gorm:"primaryKey;not null" json:"id"`
	UserID    int            `gorm:"not null;foreignKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	PhotoID   int            `gorm:"not null;foreignKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo_id"`
	Message   int            `gorm:"not null;type:text" json:"message"`
	CreatedAt time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
