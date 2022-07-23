package account

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID               int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	ExternalID       string         `gorm:"index:idx_ext_id" json:"-"`
	URL              string         `gorm:"index:idx_url,unique" json:"url"`
	Name             string         `json:"name"`
	Type             int16          `json:"type"`
	CommitGoal       int64          `json:"commit_goal"`
	CommitTimeWindow int64          `json:"commit_time_window"`
	CommitDeadline   *time.Time     `gorm:"default:null" json:"commit_deadline"`
	UserID           int64          `json:"-"`
	User             User           `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	CreatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
