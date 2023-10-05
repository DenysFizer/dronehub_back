package Controller

import (
	"dron_hub_back_/database"
	"dron_hub_back_/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
func Sorted(c *fiber.Ctx) error {
	p := Params{}
	if err := c.QueryParser(&p); err != nil {
		return err
	}
	//if p.Battary == "" && p.Height == "" && p.Fpv == "" && p.Vizor == "" && p.Range == "" && p.Maxspeed == "" {
	//	return c.JSON(fiber.Map{
	//		"massage": "bad params",
	//	})
	//}
	Drones := []models.Drone{}
	queryDB := database.DB.Table("drones")
	goodparams := make(map[string]interface{})
	if p.Height != "" {
		var err error
		goodparams["Height"], err = strconv.Atoi(p.Height)
		if err != nil {
			return err
		}
		queryDB.Where("height >= ?", goodparams["Height"])
	}
	if p.Maxspeed != "" {
		var err error
		goodparams["Maxspeed"], err = strconv.Atoi(p.Maxspeed)
		if err != nil {
			return err
		}
		queryDB.Where("maxspeed >= ?", goodparams["Maxspeed"])
	}
	if p.Range != "" {
		var err error
		goodparams["Range"], err = strconv.Atoi(p.Range)
		if err != nil {
			return err
		}
		queryDB.Where("range >= ?", goodparams["Range"])
	}
	if p.Battary != "" {
		var err error
		goodparams["Battary"], err = strconv.Atoi(p.Battary)
		if err != nil {
			return err
		}
		queryDB.Where("batterycap >= ?", goodparams["Battary"])
	}
	if p.Vizor != "" {
		var err error
		goodparams["Vizor"], err = strconv.ParseBool(p.Vizor)
		if err != nil {
			return err
		}
		queryDB.Where("vizor = ?", goodparams["Vizor"])
	}
	if p.Fpv != "" {
		var err error
		goodparams["Fpv"], err = strconv.ParseBool(p.Fpv)
		if err != nil {
			return err
		}
		queryDB.Where("fpv = ?", goodparams["Fpv"])
	}
	//database.DB.Table("drones").Where("vizor = ? AND fpv = ?", true, false).Find(&Drones)
	queryDB.Find(&Drones)
	//return c.JSON(fiber.Map{
	//	"drones": Drones,
	//	"p":      p,
	//
	//	//"map": goodparams,
	//})
	//return c.JSON(goodparams)
	return c.JSON(Drones)
}

type Params struct {
	Height   string `query:"height"`
	Maxspeed string `query:"maxspeed"`
	Range    string `query:"range"`
	Battary  string `query:"battary"`
	Vizor    string `query:"vizor"`
	Fpv      string `query:"fpv"`
}

func SmartSort(c *fiber.Ctx) error {
	p := Params{}
	if err := c.QueryParser(&p); err != nil {
		return err
	}
	Drones := []models.Drone{}
	for i := 0; i < 7; i++ {
		queryDB := database.DB.Table("drones")
		goodparams := make(map[string]interface{})
		minusheight := 100 + 100*i
		minuspeed := 50 + 50*i
		minusrange := 200 + 200*i
		minusbattary := 75 + 75*i
		if p.Height != "" {
			var err error
			var newheight int
			newheight, err = strconv.Atoi(p.Height)
			goodparams["Height"] = newheight - minusheight
			if err != nil {
				return err
			}
			goodparams["Height"] = queryDB.Where("height >= ?", goodparams["Height"])
		}
		if p.Maxspeed != "" {
			var err error
			var newmaxspeed int
			newmaxspeed, err = strconv.Atoi(p.Maxspeed)
			goodparams["Maxspeed"] = newmaxspeed - minuspeed
			if err != nil {
				return err
			}
			queryDB.Where("maxspeed >= ?", goodparams["Maxspeed"])
		}
		if p.Range != "" {
			var err error
			var newrange int
			newrange, err = strconv.Atoi(p.Range)
			goodparams["Range"] = newrange - minusrange
			if err != nil {
				return err
			}
			queryDB.Where("range >= ?", goodparams["Range"])
		}
		if p.Battary != "" {
			var err error
			var newbattary int
			newbattary, err = strconv.Atoi(p.Battary)
			goodparams["Battary"] = newbattary - minusbattary
			if err != nil {
				return err
			}
			queryDB.Where("batterycap >= ?", goodparams["Battary"])
		}

		queryDB.Find(&Drones)
		if len(Drones) != 0 {
			break
		}
	}

	return c.JSON(Drones)
}
func SetComment(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println("Error during conversion")
		return err
	}
	if err != nil {
		return err
	}
	var id int
	var err2 error
	id, err2 = strconv.Atoi(data["droneid"])

	if err2 != nil {
		fmt.Println("Error during conversion")
		return err2
	}
	comment := models.Comments{
		DroneId:  uint(id),
		UserName: data["username"],
		Comtext:  data["comtext"],
	}
	database.DB.Table("comments").Create(&comment)
	return c.JSON(fiber.Map{
		"massage": "success",
		"comment": comment,
	})
}
func GetComments(c *fiber.Ctx) error {
	comments := []models.Comments{}
	database.DB.Table("comments").Where("drone_id= ?", c.Params("id")).Find(&comments)
	return c.JSON(comments)
}
