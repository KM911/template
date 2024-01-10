package main

import (
	"os"
	"time"

	"github.com/KM911/demo/system"
)

func main() {
	// 测试cilp
	// content := "hello world"
	time.Sleep(5 * time.Second)
	system.ExecuteCommand("echo " + os.Args[1] + " | clip")
	// time.Sleep(5 * time.Second)
	// println(filepath.Dir(util.TempDirectory))

}
