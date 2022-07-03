package config

import (
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		User     string
		Password string
		Name     string
		Port     string
	}
	Firebase struct {
		CredentialJSON string
	}
}

var Cfg Config

func InitConfig() {

	Cfg.Server.Port = os.Getenv("PORT")
	Cfg.Database.Host = os.Getenv("DB_HOST")
	Cfg.Database.User = os.Getenv("DB_USER")
	Cfg.Database.Password = os.Getenv("DB_PASS")
	Cfg.Database.Name = os.Getenv("DB_NAME")
	Cfg.Database.Port = os.Getenv("DB_PORT")
	Cfg.Firebase.CredentialJSON = os.Getenv("FIREBASE_CREDENTIAL")
}
