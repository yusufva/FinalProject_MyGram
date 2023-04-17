package entity

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID               int            `gorm:"primaryKey;not null" json:"id"`
	Name             string         `gorm:"not null;type:varchar(255)" json:"name"`
	Social_Media_URL string         `gorm:"not null;type:text" json:"social_media_url"`
	UserID           int            `gorm:"not null;foreignKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	CreatedAt        time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"default:now()" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
