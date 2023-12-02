package controllers

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/models"

	"github.com/gofiber/fiber/v2"
)

type Week struct {
	Monday    []models.Schedule `json:"monday"`
	Tuesday   []models.Schedule `json:"tuesday"`
	Wednesday []models.Schedule `json:"wednesday"`
	Thursday  []models.Schedule `json:"thursday"`
	Friday    []models.Schedule `json:"friday"`
}

func GetListSchedule(c *fiber.Ctx) error {
	schedules := new([][]models.Schedule)

	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday"}
	for i := 0; i < 5; i++ {
		schedule := new([]models.Schedule)
		database.DBConn.Where("user_id = ?", c.Locals("UserId").(uint32)).Find(&schedule, "day = ?", days[i])
		*schedules = append(*schedules, *schedule)
	}

	return c.Status(200).JSON(newSuccess(Week{
		Monday:    (*schedules)[0],
		Tuesday:   (*schedules)[1],
		Wednesday: (*schedules)[2],
		Thursday:  (*schedules)[3],
		Friday:    (*schedules)[4],
	}))
}