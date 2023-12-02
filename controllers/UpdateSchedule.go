package controllers

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/errors"
	"amikom-hacktalent-be/models"
	"amikom-hacktalent-be/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateSchedule(c *fiber.Ctx) error {
	schedule := new(models.Schedule)
	temp := new(models.Schedule)

	if err := c.BodyParser(temp); err != nil {
		return c.Status(400).JSON(errors.BadRequestError(err.Error()))
	}

	if utils.IsEmptyString(temp.Title) {
		return c.Status(400).JSON(errors.BadRequestError("Title is required"))
	}

	result := database.DBConn.Take(&schedule, "schedule_id = ?", c.Query("id"))

	if result.Error != nil {
		return c.Status(404).JSON(errors.NotFound("Schedule with id " + c.Query("id") + " not found"))
	}

	if schedule.UserID != c.Locals("UserId").(uint32) {
		return c.Status(403).JSON(errors.Forbidden("Access denied!"))
	}

	schedule.Title = temp.Title
	result = database.DBConn.Save(&schedule)

	if result.Error != nil {
		return c.Status(400).JSON(errors.BadRequestError(result.Error.Error()))
	}

	return c.Status(200).JSON(newSuccess(schedule))
}