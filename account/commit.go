package account

import (
	"time"
)

type Commit struct {
	ID          int64     `gorm:"id, primarykey, autoincrement" json:"-"`
	Number      int64     `json:"number"`
	ProjectID   int64     `json:"-"`
	Project     Project   `json:"-"`
	CreatedAt   time.Time `json:"created"`
}
