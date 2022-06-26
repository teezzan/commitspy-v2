package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	ExternalID string         `gorn:"user_id" json:"-"`
	Email      string         `gorm:"email" json:"email"`
	Name       string         `gorm:"name" json:"name"`
	Avatar     string         `gorm:"avatar" json:"avatar"`
	UpdatedAt  time.Time      `gorm:"updated_at" json:"-"`
	CreatedAt  time.Time      `gorm:"created_at" json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
