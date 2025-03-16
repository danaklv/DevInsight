package handlers

import (
	"AiTaskManager/internal/api"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		err := api.LoginServiece(c)
		if err != nil {
			log.Fatal("Error in LoginServiece: ", err)

		}
		return err
	})

	app.Post("/register", func(ctx *fiber.Ctx) error {
		err := api.RegisterServiece(ctx)
		if err != nil {
			log.Fatal("Error in RegisterServiece: ", err)
		}
		return err

	})

}
