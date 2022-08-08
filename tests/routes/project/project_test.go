package project_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"github.com/teezzan/commitspy-v2/account"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/tests/setup"
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
type ProjectsDetailsResponse struct {
	Data struct {
		Projects []account.Project
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
		So(res.Data.Project.ID, ShouldNotBeBlank)
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

func (suite *ProjectRouteTestSuite) TestProjectFetchRoute() {

	Convey("Should fetch all projects for specific user", suite.T(), func() {
		var projectID string
		Convey("Should create a project for user", func() {
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

			body2 := []byte(`{
				"url":"https://github.com/body2",
				"name": "Body2",
				"type": 2
				}`)

			statusCode, err = setup.HTTPRequest(router,
				"POST",
				"/api/project/create",
				bytes.NewReader(body2),
				gin.H{"Authorization": "TestToken"},
				&res)

			So(err, ShouldBeNil)
			So(*statusCode, ShouldEqual, 201)
		})

		Convey("Should fetch projects of specific user", func() {
			var res ProjectsDetailsResponse

			error := setup.UserAccount(router)
			So(error, ShouldBeNil)

			statusCode, err := setup.HTTPRequest(router,
				"GET",
				"/api/project/",
				nil,
				gin.H{"Authorization": "TestToken"},
				&res)

			So(err, ShouldBeNil)
			So(*statusCode, ShouldEqual, 200)
			So(res.Data.Projects, ShouldHaveLength, 2)
			projectID = res.Data.Projects[0].ID
			Convey("Should fetch one project of specific user", func() {
				var res ProjectDetailsResponse

				error := setup.UserAccount(router)
				So(error, ShouldBeNil)

				statusCode, err := setup.HTTPRequest(router,
					"GET",
					fmt.Sprintf("/api/project/%s", projectID),
					nil,
					gin.H{"Authorization": "TestToken"},
					&res)

				So(err, ShouldBeNil)
				So(*statusCode, ShouldEqual, 200)
				So(res.Data.Project.ID, ShouldEqual, projectID)

			})
		})

	})
}

func (suite *ProjectRouteTestSuite) TestProjectUpdateRoute() {

	Convey("Should edit one project of specific user", suite.T(), func() {

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
		projectID := res.Data.Project.ID

		var res2 ProjectDetailsResponse

		body = []byte(`{
				"name": "Mememe Again"
				}`)

		So(error, ShouldBeNil)

		statusCode, err = setup.HTTPRequest(router,
			"POST",
			fmt.Sprintf("/api/project/%s", projectID),
			bytes.NewReader(body),
			gin.H{"Authorization": "TestToken"},
			&res2)

		So(err, ShouldBeNil)
		So(*statusCode, ShouldEqual, 202)
		So(res2.Data.Project.ID, ShouldEqual, projectID)

	})

}

func (suite *ProjectRouteTestSuite) TestProjectDeleteRoute() {

	Convey("Should delete one project of specific user", suite.T(), func() {

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
		projectID := res.Data.Project.ID

		var res2 ProjectDetailsResponse

		error = setup.UserAccount(router)
		So(error, ShouldBeNil)

		statusCode, err = setup.HTTPRequest(router,
			"DELETE",
			fmt.Sprintf("/api/project/%s", projectID),
			nil,
			gin.H{"Authorization": "TestToken"},
			&res2)

		So(err, ShouldBeNil)
		So(*statusCode, ShouldEqual, 200)

	})

}

func TestProjectRouteSuite(t *testing.T) {
	suite.Run(t, new(ProjectRouteTestSuite))
}
