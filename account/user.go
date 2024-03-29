package account

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int64          `gorm:"id, primarykey, autoincrement" json:"-"`
	ExternalID string         `gorm:"index:idx_ext_id,unique" json:"-"`
	Email      string         `gorm:"unique" json:"email"`
	Name       string         `json:"name"`
	Avatar     string         `json:"avatar"`
	Projects   []Project      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	CreatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
