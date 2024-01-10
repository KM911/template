package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world 👋!")
	})

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, Hello 👋!")
	})

	app.Post("/payload", func(ctx *fiber.Ctx) error {
		fmt.Println("ctx", string(ctx.Body()))

		return ctx.SendString("Hello, World 👋!")
	})
	app.Get("/payload", func(ctx *fiber.Ctx) error {
		// ctx.BodyParser()

		fmt.Println("ctx", string(ctx.Body()))

		return ctx.SendString("Hello, World  Get post👋!" + string(ctx.Body()))
	})

	app.Post("/file", func(ctx *fiber.Ctx) error {
		// 获取form 表单里的数据
		res := ctx.GetReqHeaders()
		fmt.Println("res headers ")
		for k, v := range res {
			fmt.Println("k", k, "v", v)
		}
		file, err := ctx.FormFile("name")
		if err != nil {
			fmt.Println("err", err)
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		fmt.Println("file", file.Filename)
		fmt.Println("file header ", file.Header)
		fmt.Println("file size ", file.Size)
		// 如何获取文件的内容 利用 ctx.Body() 读取文件内容 错误的
		ctx.SaveFile(file, file.Filename)

		// 如果是多个file有可能吗
		// fd, _ := os.OpenFile(file.Filename, os.O_RDWR|os.O_CREATE, 0666)
		// fd.Write(ctx.Body())
		// fd.Close()
		return ctx.SendString("Hello, World 👋!")
	})
	app.Get("/download", func(ctx *fiber.Ctx) error {
		//  返回文件供服务器下载
		// ctx.Response().Header.Set("Content-Type", "application/octet-stream")
		// ctx.Response().Header.Set("Content-Disposition", "attachment; filename=video.mp4")
		// 这都是sendfile给我们提供的 我们还是不够细节啊
		return ctx.SendFile("./1.mp4")

	})
	app.Get("/forbidden", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusForbidden)
	})

	//  会自动跳转到hello
	app.Get("/redirect", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/hello")
	})
	// 提交一个大文件让我看看 服务器是如何处理的 不然的话 根本就不知道还是不可以的
	app.Get("/upload", func(ctx *fiber.Ctx) error {
		// 接收从form发来的文件
		file, _ := ctx.FormFile("name")
		fmt.Println("file", file.Filename)
		return ctx.SendString("Hello, World 👋!")
	})

	// app.Get()

	// 其实proxy是什么呢? 其实非常简单我觉得 就是

	// StatusCode(app)

	//
}
