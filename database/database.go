package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"amikom-hacktalent-be/models"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	DB_USER := os.Getenv("MYSQL_USER")
	DB_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DB_DBNAME := os.Getenv("MYSQL_DBNAME")
	DB_HOST := os.Getenv("MYSQL_HOST")
	DB_PORT := "3306"

	var err error
	dsn := DB_USER +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_DBNAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		os.Exit(1)
	}

	db.AutoMigrate(&models.User{}, &models.Schedule{})
	fmt.Println("Database connected")
	DBConn = db
}