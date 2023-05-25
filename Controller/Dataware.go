package Controller

import (
	"dron_hub_back_/database"
	"dron_hub_back_/models"
	"github.com/gofiber/fiber/v2"
)

func Drones(c *fiber.Ctx) error {
	Drones := []models.Drone{}
	database.DB.Find(&Drones)
	return c.JSON(Drones)
}

func IdDrones(c *fiber.Ctx) error {
	var drone models.Drone
	database.DB.Where("id= ?", c.Params("id")).First(&drone)
	if drone.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"massage": "drone not found",
		})
	}
	return c.JSON(drone)
}
