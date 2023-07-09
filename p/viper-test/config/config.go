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
	setDebugEnv()
}

func setBasicEnv() {
	// TODO set basic env
	ExecutePath = fs.ExecutePath()
	ConfigFolderPath = filepath.Join(ExecutePath, "config")
	ConfigFilePath = filepath.Join(ConfigFolderPath, "config.json")
	viper.SetConfigFile(ConfigFilePath)
}

func checkEnv() {
	if !isEnvReady() {
		fixEnv()
		panic("config.json not exist")
	}
	viper.ReadInConfig()
}

func setDebugEnv() {
	if isDebugModel() {
		// TODO set debug env
		viper.Set("debug", true)
	}
}

func isEnvReady() bool {
	// TODO more precise check method
	return fs.IsExit(ConfigFilePath)
}

func fixEnv() {
	// TODO more fix method not just panic
	viper.WriteConfig()
}

func isDebugModel() bool {
	return viper.GetBool("debug")
}
