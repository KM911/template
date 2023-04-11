package router

import (
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	apiGroup := app.Group("/api")
	apiGroup.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Api group Hello, World!")
	})
}
