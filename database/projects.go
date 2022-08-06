package database

import (
	"fmt"
	"strconv"

	"github.com/teezzan/commitspy-v2/account"
)

func CreateProject(p *account.Project) error {
	result := db.Create(p)
	return result.Error

}

func UpdateProject(p *account.Project) error {

	result := db.Save(p)
	return result.Error

}

func DeleteProject(p *account.Project) error {

	result := db.Delete(p)
	return result.Error

}

func GetUserProjectByNameOrURL(userId int64, name string, url string) (*account.Project, error) {
	var p account.Project
	result := db.Where(&account.Project{UserID: userId, Name: name}).Or(&account.Project{UserID: userId, URL: url}).Limit(1).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &p, nil
}

func GetUserProjectById(userId int64, projectId int64) (*account.Project, error) {
	var p account.Project
	result := db.Where(&account.Project{UserID: userId, ID: projectId}).Limit(1).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 1 {
		return nil, fmt.Errorf("expected 1 match for project Id %d: found %d matches", projectId, result.RowsAffected)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &p, nil
}

func GetProjectByUUID(projectUUID string) (*account.Project, error) {
	var p account.Project

	uuid, err := strconv.ParseInt(projectUUID, 10, 64)

	if err != nil {
		return nil, err
	}

	result := db.Where(&account.Project{ID: uuid}).Limit(1).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 1 {
		return nil,
			fmt.Errorf("expected 1 match for project Id %s: found %d matches", projectUUID, result.RowsAffected)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &p, nil
}

func GetUserProjects(userId int64) (*[]account.Project, error) {
	var p []account.Project
	result := db.Where(&account.Project{UserID: userId}).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &p, nil
}

