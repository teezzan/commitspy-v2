package database

import (
	"fmt"

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

func GetUserProjectById(userId int64, projectId string) (*account.Project, error) {
	var p account.Project
	result := db.Where(&account.Project{UserID: userId, ID: projectId}).Limit(1).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 1 {
		return nil, fmt.Errorf("expected 1 match for project Id %s: found %d matches", projectId, result.RowsAffected)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &p, nil
}

func GetProjectByUUID(projectID string) (*account.Project, error) {
	var p account.Project

	result := db.Where(&account.Project{ID: projectID}).Limit(1).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 1 {
		return nil,
			fmt.Errorf("expected 1 match for project Id %s: found %d matches", projectID, result.RowsAffected)
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
