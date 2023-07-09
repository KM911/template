package main

import (
	"demo/config"
	"github.com/spf13/viper"
	"path/filepath"
)

func main() {
	viper.Set("name", "test")
	config.EnableLog(filepath.Join("config", "log.txt"))
	config.Info("info")
	println(viper.GetString("name"))
}
