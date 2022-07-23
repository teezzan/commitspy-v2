package user_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"github.com/teezzan/commitspy-v2/account"
	"github.com/teezzan/commitspy-v2/auth"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/tests/setup"
)

type UserRouteTestSuite struct {
	suite.Suite
}

func (suite *UserRouteTestSuite) SetupTest() {
	database.DropUserTable()
}

type UserDetailsResponse struct {
	Data struct {
		User account.User
	}
}

var router = setup.Router()

func (suite *UserRouteTestSuite) TestMissingToken() {
	Convey("Should return 400 for missing token", suite.T(), func() {
		var res interface{}
		statusCode, err := setup.HTTPRequest(router, "GET", "/api/user/ping", nil,
			nil, &res)

		if err != nil {
			panic(err)
		}
		So(*statusCode, ShouldEqual, 400)
	})
}

func (suite *UserRouteTestSuite) TestCreateOrLoginUser() {
	Convey("Should return user details after creating or login", suite.T(), func() {

		var res UserDetailsResponse

		statusCode, err := setup.HTTPRequest(router, "GET", "/api/user/login", nil,
			gin.H{"Authorization": "TestToken"}, &res)

		if err != nil {
			panic(err)
		}
		So(*statusCode, ShouldEqual, 201)
		So(res.Data.User.Avatar, ShouldEqual, auth.TestUserStub.Avatar)
		So(res.Data.User.Name, ShouldEqual, auth.TestUserStub.Name)
		So(res.Data.User.Email, ShouldEqual, auth.TestUserStub.Email)

		var res2 UserDetailsResponse

		statusCode, err = setup.HTTPRequest(router, "GET", "/api/user/login", nil,
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

func (suite *UserRouteTestSuite) TestGetUserDetails() {
	Convey("Should return user details", suite.T(), func() {

		error := setup.UserAccount(router)
		So(error, ShouldBeNil)

		var res UserDetailsResponse

		statusCode, err := setup.HTTPRequest(router, "GET", "/api/user/details", nil,
			gin.H{"Authorization": "TestToken"}, &res)

		if err != nil {
			panic(err)
		}
		So(*statusCode, ShouldEqual, 200)
		So(res.Data.User.Avatar, ShouldEqual, auth.TestUserStub.Avatar)
		So(res.Data.User.Name, ShouldEqual, auth.TestUserStub.Name)
		So(res.Data.User.Email, ShouldEqual, auth.TestUserStub.Email)

	})

	database.DropUserTable()

	Convey("Should not return details for unauthorised user", suite.T(), func() {
		var res UserDetailsResponse

		statusCode, err := setup.HTTPRequest(router, "GET", "/api/user/details", nil,
			gin.H{"Authorization": "TestToken"}, &res)

		if err != nil {
			panic(err)
		}
		So(*statusCode, ShouldEqual, 404)

	})

}

func TestUserRouteSuite(t *testing.T) {
	suite.Run(t, new(UserRouteTestSuite))
}
