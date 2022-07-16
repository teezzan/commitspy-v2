package routes_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/teezzan/commitspy/auth"
	"github.com/teezzan/commitspy/tests/setup"
)

var router = setup.Router()
var Request = setup.HTTPRequest

func TestPingRoute(t *testing.T) {
	Convey("Should return 200 for ping route", t, func() {
		var res map[string]interface{}

		statusCode, err := Request(router, "GET", "/api/user/ping", nil,
			gin.H{"Authorization": "TestToken"}, &res)

		if err != nil {
			panic(err)
		}

		So(*statusCode, ShouldEqual, 200)
		So(res["message"], ShouldEqual, "Ping")

	})
}

type LoginResponse struct {
	Data struct {
		User struct {
			Avatar string
			Email  string
			Name   string
		}
	}
}

func TestCreateOrLoginUser(t *testing.T) {
	Convey("Should return user details after creating or login", t, func() {

		var res LoginResponse

		statusCode, err := Request(router, "GET", "/api/user/login", nil,
			gin.H{"Authorization": "TestToken"}, &res)

		if err != nil {
			panic(err)
		}
		So(*statusCode, ShouldEqual, 200)
		So(res.Data.User.Avatar, ShouldEqual, auth.TestUserStub.Avatar)
		So(res.Data.User.Name, ShouldEqual, auth.TestUserStub.Name)
		So(res.Data.User.Email, ShouldEqual, auth.TestUserStub.Email)

		var res2 LoginResponse

		statusCode, err = Request(router, "GET", "/api/user/login", nil,
			gin.H{"Authorization": "TestToken"}, &res2)

		if err != nil {
			panic(err)
		}
		So(*statusCode, ShouldEqual, 200)
		So(res2.Data.User.Avatar, ShouldEqual, auth.TestUserStub.Avatar)
		So(res2.Data.User.Name, ShouldEqual, auth.TestUserStub.Name)
		So(res2.Data.User.Email, ShouldEqual, auth.TestUserStub.Email)

	})
}
