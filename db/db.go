package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
	"gitlab.com/dapo/crypto-trader/model"
)

var Db *gorm.DB

func InitDB(dbUser string, dbPwd string, dbName string, dbHost string, dbPort string) {
    var err error
    dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)
    Db, err = gorm.Open("mysql", dbUrl)
    if err != nil {
        fmt.Printf("Failed to connect to database %v", err)
        log.Panic("Database connection failed")
    } else {
        fmt.Println("Database connection successful")
        runDatabaseMigration()
    }
}

func runDatabaseMigration() {
	Db.AutoMigrate(
		&model.Trade{},
	)
}
