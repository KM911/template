package main

import (
	"fiber_demo/router"

	"github.com/gofiber/fiber/v2"
)

// 我们的这个应该是就是只有api请求接口的服务
func main() {
	app := fiber.New()
	//app.Static("/", "./dist")
	router.InitRouter(app)
	app.Listen(":3000")
}
