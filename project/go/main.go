package main

import (
	"demo/internal"
	"demo/util"
	"path/filepath"
)

func main() {
	util.FileLogger(filepath.Join(util.ExecutePath(), "log.log"))
	defer util.TimerStart().End()
	defer util.Recover(util.ErrorHandler)
	internal.MainProcess()
}
