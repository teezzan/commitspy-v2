package routes_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/tests/setup"
)

type ProjectRouteTestSuite struct {
	suite.Suite
}

func (suite *ProjectRouteTestSuite) SetupTest() {
	database.DropUserTable()
}

type ProjectDetailsResponse struct {
	Data struct {
		User struct {
			Avatar string
			Email  string
			Name   string
		}
	}
}

func (suite *ProjectRouteTestSuite) TestPingRoute() {

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

func TestProjectRouteSuite(t *testing.T) {
	suite.Run(t, new(ProjectRouteTestSuite))
}
