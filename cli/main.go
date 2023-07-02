package main

import (
	"cli/command"
	"cli/config"
	"github.com/KM911/oslib"
	"os"
)

//目前没有任何的扩展能力不是吗

func main() {

	//可以直接显示我们的火焰图不可以吗
	defer oslib.TimerStart().End()
	config.LoadConfig()
	argslens := len(os.Args)
	if argslens < 2 {
		println("need a valid command")
		return
	}
	// 执行功能单一 错误输出不正确 还是利用cli框架吧
	for i := 1; i < argslens; i++ {
		// TODO
		command.Start(os.Args[i])
	}
}
