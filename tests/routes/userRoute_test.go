package routes_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/teezzan/commitspy/tests/test_utils"
	"gopkg.in/h2non/baloo.v3"
)

var test *baloo.Client

func init() {
	test_utils.SetupRouter()
	test = baloo.New("http://localhost:3000/api")
}
func TestPingRoute(t *testing.T) {
	Convey("Should return 200 for ping route", t, func() {
		test.Get("/user/ping").
			SetHeader("Foo", "Bar").
			Expect(t).
			Status(200)

	})
}
