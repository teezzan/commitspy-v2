package webhook_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
	"github.com/teezzan/commitspy-v2/account"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/tests/setup"
)

type WebhookRouteTestSuite struct {
	suite.Suite
}

type WebhookResponse struct {
	Data   interface{}
	Status string
}

type ProjectDetailsResponse struct {
	Data struct {
		Project account.Project
	}
}

func (suite *WebhookRouteTestSuite) SetupTest() {
	database.DropProjectTable()
	database.DropCommitTable()
}

var router = setup.Router()

func (suite *WebhookRouteTestSuite) TestGitlabWebhookRoute() {

	Convey("Should authenticate gitlab project", suite.T(), func() {
		var res WebhookResponse

		error := setup.UserAccount(router)
		So(error, ShouldBeNil)

		projectID, error := setup.Project(router)
		So(error, ShouldBeNil)

		statusCode, err := setup.HTTPRequest(router,
			"POST",
			fmt.Sprint("/api/webhooks/gl/", *projectID),
			bytes.NewReader(GitlabBody),
			gin.H{
				"x-gitlab-event": "Push Hook",
				"x-gitlab-token": *projectID,
			},
			&res)

		So(err, ShouldBeNil)
		So(*statusCode, ShouldEqual, 202)

		Convey("Should update gitlab project", func() {
			var res2 ProjectDetailsResponse

			body := []byte(`{
				"commit_goal": 5,
				"commit_time_window": 24
				}`)

			So(error, ShouldBeNil)

			statusCode, err = setup.HTTPRequest(router,
				"POST",
				fmt.Sprintf("/api/project/%s", *projectID),
				bytes.NewReader(body),
				gin.H{"Authorization": "TestToken"},
				&res2)

			So(err, ShouldBeNil)
			Convey("Should register push event for gitlab project", func() {
				var res2 WebhookResponse

				statusCode, err := setup.HTTPRequest(router,
					"POST",
					fmt.Sprint("/api/webhooks/gl/", *projectID),
					bytes.NewReader(GitlabBody),
					gin.H{
						"x-gitlab-event": "Push Hook",
						"x-gitlab-token": *projectID,
					},
					&res2)

				So(err, ShouldBeNil)
				So(*statusCode, ShouldEqual, 200)

				Convey("Should fetch project details", func() {
					var res ProjectDetailsResponse
					So(error, ShouldBeNil)

					statusCode, err := setup.HTTPRequest(router,
						"GET",
						fmt.Sprintf("/api/project/%s", *projectID),
						nil,
						gin.H{"Authorization": "TestToken"},
						&res)

					So(err, ShouldBeNil)
					So(*statusCode, ShouldEqual, 200)
					So(res.Data.Project.ID, ShouldEqual, *projectID)
					So(res.Data.Project.CurrentCohort.Number, ShouldEqual, 3)
				})

			})

		})

	})

}

func (suite *WebhookRouteTestSuite) TestGithubWebhookRoute() {

	Convey("Should authenticate github project", suite.T(), func() {
		var res WebhookResponse

		error := setup.UserAccount(router)
		So(error, ShouldBeNil)

		projectID, error := setup.Project(router)
		So(error, ShouldBeNil)

		statusCode, err := setup.HTTPRequest(router,
			"POST",
			fmt.Sprint("/api/webhooks/gh/", *projectID),
			bytes.NewReader(GithubBody),
			gin.H{
				"x-github-event":      "ping",
				"x-hub-signature-256": fmt.Sprint("sha256=", generateSHA256(*projectID, string(GithubBody))),
			},
			&res)

		So(err, ShouldBeNil)
		So(*statusCode, ShouldEqual, 202)

		Convey("Should update github project", func() {
			var res2 ProjectDetailsResponse

			body := []byte(`{
				"commit_goal": 5,
				"commit_time_window": 24
				}`)

			So(error, ShouldBeNil)

			statusCode, err = setup.HTTPRequest(router,
				"POST",
				fmt.Sprintf("/api/project/%s", *projectID),
				bytes.NewReader(body),
				gin.H{"Authorization": "TestToken"},
				&res2)

			So(err, ShouldBeNil)
			Convey("Should register push event for github project", func() {
				var res2 WebhookResponse

				statusCode, err := setup.HTTPRequest(router,
					"POST",
					fmt.Sprint("/api/webhooks/gh/", *projectID),
					bytes.NewReader(GithubBody),
					gin.H{
						"x-github-event":      "push",
						"x-hub-signature-256": fmt.Sprint("sha256=", generateSHA256(*projectID, string(GithubBody))),
					},
					&res2)

				So(err, ShouldBeNil)
				So(*statusCode, ShouldEqual, 200)

				Convey("Should fetch project details", func() {
					var res ProjectDetailsResponse
					So(error, ShouldBeNil)

					statusCode, err := setup.HTTPRequest(router,
						"GET",
						fmt.Sprintf("/api/project/%s", *projectID),
						nil,
						gin.H{"Authorization": "TestToken"},
						&res)

					So(err, ShouldBeNil)
					So(*statusCode, ShouldEqual, 200)
					So(res.Data.Project.ID, ShouldEqual, *projectID)
					So(res.Data.Project.CurrentCohort.Number, ShouldEqual, 1)
				})

			})

		})

	})

}

func generateSHA256(token string, body string) string {
	h := hmac.New(sha256.New, []byte(token))
	h.Write([]byte(body))
	return hex.EncodeToString(h.Sum(nil))
}
func TestWebhookRouteSuite(t *testing.T) {
	suite.Run(t, new(WebhookRouteTestSuite))
}
