package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID                   int64          `gorm:"id, primarykey, autoincrement" json:"id"`
	ExternalID           string         `gorm:"index:idx_ext_id,unique" json:"-"`
	Name                 string         `json:"name"`
	CommitGoal           int64          `json:"commit_goal"`
	CommitGoalTimeWindow int64          `json:"commit_goal_time_window"`
	User                 User           `json:"-"`
	UpdatedAt            time.Time      `json:"-"`
	CreatedAt            time.Time      `json:"-"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
}
