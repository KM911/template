package app

import (
	"fiber_demo/router"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()
	//app.Static("/", "./dist")
	router.InitRouter(app)
	app.Listen(":3000")
}
