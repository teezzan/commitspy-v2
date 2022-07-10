package database

import (
	"fmt"

	"github.com/teezzan/commitspy/account"
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
