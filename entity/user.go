package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primaryKey;not null;" json:"id"`
	Username  string         `gorm:"unique;not null;type:varchar(255)" json:"username"`
	Email     string         `gorm:"unique;not null;type:varchar(255)" json:"email"`
	Password  string         `gorm:"not null;type:text;" json:"password"`
	Age       int            `gorm:"not null;type:int" json:"age"`
	CreatedAt time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
