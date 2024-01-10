package app

import (
	"fiber_demo/router"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	// "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func Start() {
	app = fiber.New(
		fiber.Config{
			// 显示请求信息
			// Prefork:       true,
			// CaseSensitive: true,
			// StrictRouting: true,
			ReadTimeout:  5 * time.Second,
			ServerHeader: "Fiber",
			AppName:      "Test App v1.0.1",
		},
	)

	// 显示debug信息

	app.Static("/", "./dist")
	// log.SetLevel(log.LevelInfo)
	cros(app)
	router.InitRouter(app)
	// port := strconv.Itoa(rand.Intn(1000) + 3000)
	// 随机将3000-4000之间的端口号
	fmt.Println("server runing at prot 3000")
	app.Listen(":3000")

	// log.Fatal(app.Listen(":4000"))
}
func cros(app *fiber.App) {
	// app.Use(cors.New())
	// 或者扩展你的配置以进行自定义
	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost:4000",
		// AllowMethods: "POST,OPTIONS",
		// AllowOrigins: "localhost:4000",
	}))
	app.Use(func(c *fiber.Ctx) error {
		fmt.Print(c.Method(), " ：", c.Path(), "\n")
		// fmt.Print("hello")
		c.Set("Server", "Fiber_builde_by_km911")
		return c.Next()
	})
	// app.use

}
func Shutdown() error {
	return app.Shutdown()
}
