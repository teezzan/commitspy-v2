package database

import (
	"github.com/teezzan/commitspy-v2/account"
)

func CreateCommit(c *account.Commit) error {
	result := db.Create(&c)
	return result.Error

}

func UpdateCommit(c *account.Commit) error {

	result := db.Save(c)
	return result.Error

}

func GetCommitsByProjectUUID(projectID string) (*[]account.Commit, error) {
	var c []account.Commit

	result := db.Where(&account.Commit{ProjectID: projectID}).Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &c, nil
}

func CountCommitsByProjectUUID(projectID string) (*int64, error) {
	var count int64

	result := db.Model(&account.Commit{}).Where("project_id", projectID).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	return &count, nil
}

func GetCurrentCommitCohort(project *account.Project) (*account.Commit, error) {
	// var cohort account.CommitCohort
	var commits account.Commit

	// startTimeWindow := project.CommitDeadline.Add(-1 * time.Hour * time.Duration(project.CommitTimeWindow))
	result := db.Model(&account.Commit{}).Where("created_at < ?", project.CommitDeadline).Find(&commits)
	if result.Error != nil {
		return nil, result.Error
	}
	return &commits, nil
}
