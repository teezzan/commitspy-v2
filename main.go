package main

import (
	"github.com/teezzan/commitspy/config"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/routes"
	"github.com/teezzan/commitspy/utils"
)

func main() {
	utils.SetEnviroment("")
	config.InitConfig()
	config.InitFirebase()
	database.InitDB()
	routes.Run()
}
