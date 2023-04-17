package config

import "github.com/spf13/viper"

func INIt() {
	viper.SetConfigFile("./config/config.json") // name of config file (without extension)

	// read config file
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}
}
