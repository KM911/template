package config

import (
	"github.com/KM911/oslib/fs"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	ExecutePath      string
	ConfigFolderPath string
	ConfigFilePath   string
)

func LoadConfig() {
	setBasicEnv()
	checkEnv()
}

func setBasicEnv() {
	ExecutePath = fs.ExecutePath()
	ConfigFolderPath = filepath.Join(ExecutePath, "config")
	ConfigFilePath = filepath.Join(ConfigFolderPath, "config.json")
	viper.SetConfigFile(ConfigFilePath)
	// set default value for config
}

func checkEnv() {
	// Check if the file exists
	if !fs.IsExit(ConfigFilePath) {
		// TODO more fix method not panic
		fixEnv()
		panic("config.json not exist")
	}
	viper.ReadInConfig()
}

func fixEnv() {
	viper.WriteConfig()
}
