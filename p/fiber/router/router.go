package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))
		return c.SendString(string(c.Body()))
	})
	apiGroup := app.Group("/api")
	apiGroup.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Api group Hello, World!")
	})
}
