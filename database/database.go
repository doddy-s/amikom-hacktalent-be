package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"amikom-hacktalent-be/models"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Schedule{})
	DBConn = db
}
