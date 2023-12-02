package middlewares

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/errors"
	"amikom-hacktalent-be/models"
	"amikom-hacktalent-be/utils"

	"github.com/gofiber/fiber/v2"
)

func EmailCheck(c *fiber.Ctx) error {
	email := c.Query("email")
	if utils.IsEmptyString(email) {
		return c.Status(400).JSON(errors.BadRequestError("Email is required"))
	}

	if !utils.IsEmail(email) {
		return c.Status(400).JSON(errors.BadRequestError("Invalid email"))
	}

	user := new(models.User)
	result := database.DBConn.Take(&user, "email = ?", c.Query("email"))

	if result.Error != nil {
		return c.Status(404).JSON(errors.NotFound("Email is not found"))
	}

	c.Locals("UserId", user.UserId)
	return c.Next()
}
