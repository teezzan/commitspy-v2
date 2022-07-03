package test_utils

import (
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
