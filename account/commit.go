package account

import (
	"time"
)

type Commit struct {
	ID          int64     `gorm:"id, primarykey, autoincrement" json:"-"`
	ExternalIDs []string  `gorm:"type:text" json:"commit_id"`
	Number      int64     `json:"number"`
	ProjectID   int64     `json:"-"`
	Project     Project   `json:"-"`
	CreatedAt   time.Time `json:"created"`
}
