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
	//SetDefult()
	fmt.Println("migration is success")
}
func SetDefult() {
	dronearr := []models.Drone{
		//{
		//	Creator:    "Drone",
		//	Name:       "test1",
		//	Maxspeed:   85,
		//	Batterycap: 3500,
		//	Height:     3,
		//	Range:      15,
		//	Fpv:        false,
		//	Vizor:      false,
		//	Text:       "Гарний дрон",
		//	Imgpath:    "drone",
		//},
		//{
		//	Creator:    "Drone",
		//	Name:       "test2",
		//	Maxspeed:   85,
		//	Batterycap: 3500,
		//	Height:     3,
		//	Range:      15,
		//	Fpv:        false,
		//	Vizor:      false,
		//	Text:       "Гарний дрон",
		//	Imgpath:    "drone",
		//},
		{
			Creator:    "Drone1",
			Name:       "test2424",
			Maxspeed:   85,
			Batterycap: 3500,
			Height:     3,
			Range:      15,
			Fpv:        false,
			Vizor:      false,
			Text:       "Дуже Гарний дрон",
			Imgpath:    "drone",
		},
	}
	for _, drone := range dronearr {
		DB.Create(&drone)
	}
	fmt.Println("Success")
}
