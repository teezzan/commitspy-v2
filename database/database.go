package database

import (
	"fmt"
	"log"

	"github.com/teezzan/commitspy/account"
	"github.com/teezzan/commitspy/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dsn string

func InitDB() {
	var err error

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Cfg.Database.Host,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Name,
		config.Cfg.Database.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		MigrateDBModels()
	}
}

func GetDB() *gorm.DB {
	return db
}

func MigrateDBModels() {
	db.AutoMigrate(&account.User{}, &account.Project{})
}

func DropUserTable() {
	db.Exec("DELETE FROM users")
}

func DropProjectTable() {
	db.Exec("DELETE FROM projects")
}