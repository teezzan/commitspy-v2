package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/account"
)

type UserDetailsResponse struct {
	Data struct {
		User account.User
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
