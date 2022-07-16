package setup

import (
	"github.com/gin-gonic/gin"
)

type UserDetailsResponse struct {
	Data struct {
		User struct {
			Avatar string
			Email  string
			Name   string
		}
	}
}

var Request = HTTPRequest

func UserAccount(router *gin.Engine) {
	var res UserDetailsResponse

	statusCode, err := Request(router, "GET", "/api/user/login", nil,
		gin.H{"Authorization": "TestToken"}, &res)

	if err != nil || *statusCode != 201 {
		panic(err)
	}

}
