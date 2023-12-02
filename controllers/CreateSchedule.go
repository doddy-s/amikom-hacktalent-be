package controllers

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/errors"
	"amikom-hacktalent-be/models"
	"amikom-hacktalent-be/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateSchedule(c *fiber.Ctx) error {
	fmt.Println(c.Locals("UserId"))
	
	schedule := new(models.Schedule)
	schedule.UserID = c.Locals("UserId").(uint32)


	if err := c.BodyParser(schedule); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if utils.IsEmptyString(schedule.Title) {
		return c.Status(400).JSON(errors.BadRequestError("Title is required"))
	}

	if utils.IsEmptyString(schedule.Day) {
		return c.Status(400).JSON(errors.BadRequestError("Day is required"))
	}

	if !utils.IsDay(schedule.Day) {
		return c.Status(400).JSON(errors.BadRequestError("Invalid day"))
	}

	result := database.DBConn.Create(&schedule)

	if result.Error != nil {
		return c.Status(400).JSON(result.Error)
	}

	return c.Status(200).JSON(newSuccess(schedule))
}