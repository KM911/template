package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func StatusCode(app *fiber.App) {
	//  状态码系列集合
	app.Get("/status/:code", func(ctx *fiber.Ctx) error {
		code, err := ctx.ParamsInt("code")
		fmt.Println(code)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		return ctx.SendStatus(code)
	})

	//  通用的可以吗? 当url为 xxx 三位数就自动匹配

}
