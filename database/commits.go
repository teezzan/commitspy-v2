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
