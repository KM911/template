package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	// Any type
	filepath string

	// config value

	Time int `json:"time"`
}

func LoadConfig(filepath string) (config Config) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}
	}
	fmt.Println(string(data))
	// 这里需要将data 进行反序列化
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	config.filepath = filepath
	return config
}

func (config *Config) WriteConfig() {
	// 这里需要将config 进行序列化
	data, err := json.Marshal(&config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	os.WriteFile(config.filepath, data, 0777)
}

func main() {
	config := LoadConfig("config.json")
	//config.filepath = "config.json"
	config.Time = 1000
	config.WriteConfig()

}
