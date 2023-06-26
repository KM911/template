package main

import (
	"cli/command"
	"cli/config"
	"github.com/KM911/oslib"
	"os"
)

func main() {
	// 可以直接显示我们的火焰图不可以吗?

	defer oslib.TimerStart().End()
	config.LoadConfig()
	argslens := len(os.Args)
	if argslens < 2 {
		println("need a valid command")
		return
	}
	for i := 1; i < argslens; i++ {
		// TODO
		command.Start(os.Args[i])
	}
}
