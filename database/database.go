package database

import (
	"fmt"
	"log"

	"github.com/teezzan/commitspy-v2/account"
	"github.com/teezzan/commitspy-v2/config"
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

func _() *gorm.DB {
	return db
}

func MigrateDBModels() {
	err := db.AutoMigrate(&account.User{}, &account.Project{}, &account.Commit{})
	if err != nil {
		return
	}
}

func DropUserTable() {
	db.Exec("TRUNCATE TABLE users CASCADE;")
}

func DropProjectTable() {
	db.Exec("TRUNCATE TABLE projects CASCADE;")
}

func DropCommitTable() {
	db.Exec("TRUNCATE TABLE commits CASCADE;")
}
