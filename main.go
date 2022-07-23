package main

import (
	"github.com/teezzan/commitspy-v2/config"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/env"
	"github.com/teezzan/commitspy-v2/routes"
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
