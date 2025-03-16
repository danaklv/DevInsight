package api

import (
	"AiTaskManager/internal/config"
	"AiTaskManager/internal/models"
	"github.com/gofiber/fiber/v2"
)

func LoginServiece(c *fiber.Ctx) error {
	//err := c.BodyParser(&models.User{})
	//if err != nil {
	//	return err
	//}

	return nil

}

func RegisterServiece(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err := config.InsertUser(&user)
	if err != nil {

		return err
	}
	return nil
}
