package database

import (
	"github.com/teezzan/commitspy/account"
)

func CreateProject(p *account.Project) error {
	result := db.Create(p)
	return result.Error

}

func GetUserProjectByNameOrURL(id int64, name string, url string) (*account.Project, error) {
	var p account.Project
	result := db.Where(&account.Project{UserID: id, Name: name}).Or(&account.Project{UserID: id, URL: url}).Limit(1).Find(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &p, nil
}
