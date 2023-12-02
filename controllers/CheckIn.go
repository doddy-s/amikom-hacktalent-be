package controllers

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/errors"
	"amikom-hacktalent-be/models"
	"amikom-hacktalent-be/utils"

	"github.com/gofiber/fiber/v2"
)

func CheckIn(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if utils.IsEmptyString(user.Email) {
		return c.Status(400).JSON(errors.BadRequestError("Email is required"))
	}

	if !utils.IsEmail(user.Email) {
		return c.Status(400).JSON(errors.BadRequestError("Invalid email"))
	}

	result := database.DBConn.FirstOrCreate(&user, models.User{Email: user.Email})

	if result.Error != nil {
		return c.Status(400).JSON(result.Error)
	}

	return c.Status(200).JSON(newSuccess(user))
}