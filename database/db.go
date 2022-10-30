package database

import (
	"github.com/LucasGabrielBecker/golang-ecommerce/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecToDB() {

	db, err := gorm.Open(sqlite.Open("./database/ecommerce.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{})

	DB = db
}
