package main

import (
	"github.com/KM911/oslib"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./dist")
	oslib.Run("start http://localhost:3000/index.html")
	app.Listen(":3000")
}
