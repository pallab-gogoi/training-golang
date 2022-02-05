// initialize connection b/w database and this code
package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pallab-gogoi/covid_vaccine/model"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
func setUp() {
	host := "localhost"
	port := "5432"
	dbName := "vaccine"
	userName := "postgres"
	passWord := "postgres"
	arg := "host=" + host + " port=" + port + " user=" + userName + " dbname=" + dbName + " sslmode=disable password=" + passWord
	db, err := gorm.Open("postgres", arg)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(model.Details{})
	DB = db
}
