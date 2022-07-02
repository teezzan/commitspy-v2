package routes_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/teezzan/commitspy/tests/test_utils"
	"gopkg.in/h2non/baloo.v3"
)

var test *baloo.Client

func TestPingRoute(t *testing.T) {
	test_utils.SetupRouter()
	test = baloo.New("http://localhost:5000/api")
	Convey("Should return 200 for ping route", t, func() {
		test.Get("/user/ping").
			SetHeader("Foo", "Bar").
			Expect(t).
			Status(200)

	})
}
