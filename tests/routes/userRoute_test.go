package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/teezzan/commitspy/tests/test_utils"
)

var router = test_utils.SetupRouter()

func TestPingRoute(t *testing.T) {
	Convey("Should return 200 for ping route", t, func() {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/user/ping", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

	})

}
