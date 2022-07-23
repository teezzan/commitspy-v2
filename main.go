package main

import (
	"github.com/teezzan/commitspy/config"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/env"
	"github.com/teezzan/commitspy/routes"
)

func main() {
	err := env.SetEnviroment("")
	if err != nil {
		return
	}
	config.InitConfig()
	config.InitFirebase()
	database.InitDB()
	routes.Run()
}
