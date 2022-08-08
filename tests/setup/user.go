package setup

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/account"
)

type UserDetailsResponse struct {
	Data struct {
		User account.User
	}
}

type ProjectDetailsResponse struct {
	Data struct {
		Project account.Project
	}
}

var Request = HTTPRequest

func UserAccount(router *gin.Engine) *error {
	var res UserDetailsResponse

	statusCode, err := Request(router, "GET", "/api/user/login", nil,
		gin.H{"Authorization": "TestToken"}, &res)

	if (err != nil) || ((*statusCode != 200) && (*statusCode != 201)) {

		return &err
	}
	return nil
}

func Project(router *gin.Engine) (*string, *error) {
	var res ProjectDetailsResponse

	body := []byte(`{
		"url":"https://github.com/memme/",
		"name": "Mememe",
		"type": 1
		}`)

	statusCode, err := Request(router,
		"POST",
		"/api/project/create",
		bytes.NewReader(body),
		gin.H{"Authorization": "TestToken"},
		&res)

	if (err != nil) || ((*statusCode != 200) && (*statusCode != 201)) {
		return nil, &err
	}
	return &res.Data.Project.ID, nil
}
