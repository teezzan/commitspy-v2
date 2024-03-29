package setup

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/config"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/env"
	"github.com/teezzan/commitspy-v2/routes"
)

func init() {
	err := os.Setenv("ENV", "TEST")
	if err != nil {
		return
	}
}

func Router() *gin.Engine {
	err := env.SetEnviroment(".test")
	if err != nil {
		return nil
	}
	config.InitConfig()
	config.InitFirebase()
	database.InitDB()
	router := routes.LoadRoutesAndReturnRouter()
	return router
}

type LoginResponse struct {
	data struct {
		user struct {
			avater string
			email  string
			name   string
		}
	}
}

func HTTPRequest(router *gin.Engine, method string, url string, body io.Reader, headers gin.H, target interface{}) (*int, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("error occured %v", err)
	}
	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Set(key, val.(string))
		}
	}

	router.ServeHTTP(w, req)

	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}
	// log.Println("unmash:", target)
	return &(w.Code), nil

}
