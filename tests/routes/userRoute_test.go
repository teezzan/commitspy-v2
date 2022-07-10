package routes_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/teezzan/commitspy/tests/setup"
)

var router = setup.Router()
var Request = setup.HTTPRequest

func TestPingRoute(t *testing.T) {
	Convey("Should return 200 for ping route", t, func() {

		req, w := Request("GET", "/api/user/ping", nil,
			gin.H{"Authorization": "TestToken"})
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

	})

}
