// database/database.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"inventory_control/config"
	"inventory_control/model"
	"log"
	"strings"
)

var DB *gorm.DB
var err error

func init() {
	// Get DSN from config based on DevMode
	var dsn string
	if config.AppConfig.DevMode {
		dsn = config.AppConfig.DSN.Dev
	} else {
		dsn = config.AppConfig.DSN.Production
	}
	if strings.HasPrefix(dsn, "DSN=") {
		// dsn = strings.TrimPrefix(dsn, "DSN=") // "DSN="がある場合は取り除く
		log.Fatalln("DSN are not expected to contain the 'DSN=' prefix. Please correct the value in config.json")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + " database can't connect")
	}

	// Using AutoMigrate to create tables for BookInfo and BookStock
	err = DB.AutoMigrate(
		&model.TypeBranch{},
		&model.TypeGenre{},
		&model.BookInfo{},
		&model.BookStock{})
	if err != nil {
		log.Fatalln("Failed to auto-migrate tables")
	}
}
