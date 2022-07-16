package project_test

import (
	"bytes"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"github.com/teezzan/commitspy/account"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/tests/setup"
)

type ProjectRouteTestSuite struct {
	suite.Suite
}

func (suite *ProjectRouteTestSuite) SetupTest() {
	database.DropProjectTable()
}

type ProjectDetailsResponse struct {
	Data struct {
		Project account.Project
	}
}

var router = setup.Router()

func (suite *ProjectRouteTestSuite) TestProjectCreateRoute() {

	Convey("Should create project and return details", suite.T(), func() {

		var res ProjectDetailsResponse

		error := setup.UserAccount(router)
		So(error, ShouldBeNil)

		body := []byte(`{
			"url":"https://github.com/memme/",
			"name": "Mememe",
			"type": 1
			}`)

		statusCode, err := setup.HTTPRequest(router,
			"POST",
			"/api/project/create",
			bytes.NewReader(body),
			gin.H{"Authorization": "TestToken"},
			&res)

		So(err, ShouldBeNil)
		So(*statusCode, ShouldEqual, 201)
		So(res.Data.Project.ID, ShouldBeGreaterThanOrEqualTo, 1)
		So(res.Data.Project.URL, ShouldEqual, "https://github.com/memme/")
		So(res.Data.Project.Name, ShouldEqual, "Mememe")
		So(res.Data.Project.Type, ShouldEqual, 1)
	})

	database.DropProjectTable()

	Convey("Should prevent duplicate project creation", suite.T(), func() {

		var res ProjectDetailsResponse

		body := []byte(`{
			"url":"https://github.com/memme/",
			"name": "Mememe",
			"type": 1
			}`)

		statusCode, err := setup.HTTPRequest(router,
			"POST",
			"/api/project/create",
			bytes.NewReader(body),
			gin.H{"Authorization": "TestToken"},
			&res)

		So(err, ShouldBeNil)
		So(*statusCode, ShouldEqual, 201)

		var res2 map[string]interface{}

		statusCode, err = setup.HTTPRequest(router,
			"POST",
			"/api/project/create",
			bytes.NewReader(body),
			gin.H{"Authorization": "TestToken"},
			&res2)

		So(err, ShouldBeNil)
		So(*statusCode, ShouldEqual, 409)


	})
}
func TestProjectRouteSuite(t *testing.T) {
	suite.Run(t, new(ProjectRouteTestSuite))
}
