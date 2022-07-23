package database

import (
	"fmt"

	"github.com/teezzan/commitspy/account"
)

func CreateUser(u *account.User) error {
	result := db.Create(u)
	return result.Error

}

func GetUserByExternalID(id string) (*account.User, error) {
	var u account.User
	result := db.Where(&account.User{ExternalID: id}).Limit(1).Find(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 1 {
		return nil, fmt.Errorf("expected 1 match for externalID %s: found %d matches", id, result.RowsAffected)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &u, nil
}
