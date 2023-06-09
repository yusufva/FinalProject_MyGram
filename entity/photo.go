package entity

import (
	"final-project/dto"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        int    `gorm:"primaryKey;not null" json:"id"`
	Title     string `gorm:"not null;type:varchar(255)" json:"title"`
	Caption   string `gorm:"not null;type:text" json:"caption"`
	Photo_Url string `gorm:"not null;type:text" json:"photo_url"`
	UserID    int    `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	// if userid fails add ;type:int;foreignKey:user_id;references:id
	CreatedAt time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (p *Photo) EntityToProductResponseDto() dto.PhotoResponse {
	return dto.PhotoResponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		Photo_Url: p.Photo_Url,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
