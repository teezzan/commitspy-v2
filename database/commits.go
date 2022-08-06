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
