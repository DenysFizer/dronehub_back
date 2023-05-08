package database

import (
	"dron_hub_back_/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "postgres://postgres:password@localhost:5432/dronehub?sslmode=disable"

var DB *gorm.DB

func Connection() {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	conn.AutoMigrate(&models.User{}, &models.Drone{})
	DB = conn
	fmt.Println("migration is success")
}
