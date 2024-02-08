package database

import (
	"log"
	database "sam0307204/jwt-Authentication/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	db, dbError := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("cannot connect to Instance")
	}
	log.Println("Connected to Database")
	Instance = db
}

func Migrate() {
	Instance.AutoMigrate(&database.User{})
	log.Println("Database Migrated Completed.")

}
