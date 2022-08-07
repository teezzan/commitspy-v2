package account

import (
	"time"
)

type Commit struct {
	ID        int64     `gorm:"id, primarykey, autoincrement" json:"-"`
	Number    int64     `json:"number"`
	ProjectID string    `json:"-"`
	Project   Project   `json:"-"`
	CreatedAt time.Time `json:"created"`
}

type CommitCohort struct {
	Number         int64     `json:"number_of_commit"`
	CommitDeadline time.Time `json:"commit_deadline"`
}
