package entity

import (
	"final-project/dto"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `gorm:"primaryKey;not null" json:"id"`
	UserID    int            `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	PhotoID   int            `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo_id"`
	Message   string         `gorm:"not null;type:text" json:"message"`
	CreatedAt time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (c *Comment) EntityToProductResponseDto() dto.CommentResponse {
	return dto.CommentResponse{
		ID:        c.ID,
		UserID:    c.UserID,
		PhotoID:   c.PhotoID,
		Message:   c.Message,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
