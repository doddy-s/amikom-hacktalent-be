package controllers

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/errors"
	"amikom-hacktalent-be/models"
	"amikom-hacktalent-be/utils"

	"github.com/gofiber/fiber/v2"
)

func GetDetailSchedule(c *fiber.Ctx) error {
	if utils.IsEmptyString(c.Query("day")) {
		return c.Next()
	}

	if !utils.IsDay(c.Query("day")) {
		return c.Status(400).JSON(errors.BadRequestError("Day is invalid"))
	}

	schedule := new([]models.Schedule)
	result := database.DBConn.Where("user_id = ?", c.Locals("UserId").(uint32)).Find(&schedule, "day = ?", c.Query("day"))

	if result.Error != nil {
		return c.Status(400).JSON(errors.BadRequestError("Schedule not found"))
	}

	return c.Status(200).JSON(newSuccess(schedule))
}
