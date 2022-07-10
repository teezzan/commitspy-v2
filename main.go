package main

import (
	"github.com/teezzan/commitspy/config"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/env"
	"github.com/teezzan/commitspy/routes"
)

func main() {
	env.SetEnviroment("")
	config.InitConfig()
	config.InitFirebase()
	database.InitDB()
	routes.Run()
}
