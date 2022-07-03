package test_utils

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/config"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/routes"
	"github.com/teezzan/commitspy/utils"
)

func init() {
	os.Setenv("ENV", "TEST")
}

func SetupRouter() *gin.Engine {
	utils.SetEnviroment(".test")
	config.InitConfig()
	config.InitFirebase()
	database.InitDB()
	router := routes.LoadRoutesAndReturnRouter()
	return router
}

func Request(method string, url string, body io.Reader, headers gin.H) (*http.Request, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/user/ping", nil)
	if err != nil {
		log.Fatalf("error occured %v", err)
	}
	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Set(key, val.(string))
		}
	}
	return req, w
}
