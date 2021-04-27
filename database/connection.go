package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simrs-backend/models"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root@/db_simrs"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
