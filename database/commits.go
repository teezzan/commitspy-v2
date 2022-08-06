package database

import (
	"strconv"

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

func GetCommitsByProjectUUID(projectUUID string) (*[]account.Commit, error) {
	var c []account.Commit

	uuid, err := strconv.ParseInt(projectUUID, 10, 64)

	if err != nil {
		return nil, err
	}

	result := db.Where(&account.Commit{ProjectID: uuid}).Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &c, nil
}

func CountCommitsByProjectUUID(projectUUID string) (*int64, error) {
	var count int64

	uuid, err := strconv.ParseInt(projectUUID, 10, 64)

	if err != nil {
		return nil, err
	}

	result := db.Model(&account.Commit{}).Where("project_id", uuid).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	return &count, nil
}
