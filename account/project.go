package account

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID               string         `gorm:"id, primarykey" json:"id"`
	ExternalID       string         `gorm:"index:idx_ext_id" json:"-"`
	URL              string         `gorm:"index:idx_url,unique" json:"url"`
	Name             string         `json:"name"`
	Type             int16          `json:"type"`
	CommitGoal       int64          `json:"commit_goal"`
	CommitTimeWindow int64          `json:"commit_time_window"`
	CommitDeadline   *time.Time     `gorm:"default:null" json:"-"`
	CurrentCohort    *CommitCohort   `gorm:"-:all" json:"current_cohort"`
	UserID           int64          `json:"-"`
	User             User           `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	CreatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}

type CommitCohort struct {
	Number         int64     `json:"number_of_commit"`
	CommitDeadline time.Time `json:"commit_deadline"`
}
