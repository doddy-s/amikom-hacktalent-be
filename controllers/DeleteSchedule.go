package controllers

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/errors"
	"amikom-hacktalent-be/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteSchedule(c *fiber.Ctx) error {
	schedule := new(models.Schedule)

	result := database.DBConn.Take(&schedule, "schedule_id = ?", c.Query("id"))

	if result.Error != nil {
		return c.Status(404).JSON(errors.NotFound("Schedule with id " + c.Query("id") + " not found"))
	}

	if schedule.UserID != c.Locals("UserId").(uint32) {
		return c.Status(403).JSON(errors.Forbidden("Access denied!"))
	}

	result = database.DBConn.Delete(&schedule)

	if result.Error != nil {
		return c.Status(400).JSON(errors.BadRequestError(result.Error.Error()))
	}

	nah := new(Empty)
	
	return c.Status(200).JSON(newSuccess(nah))
}