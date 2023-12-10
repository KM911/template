package config

import (
	"encoding/json"
	"os"

	"github.com/KM911/demo/util"
)

const (
	JsonConfigurationFile = "ap.json"
)

type JsonConfiguration struct {
	Name         string
	Target       string
	BuildCommand string
	BuildOutput  string
	ExecuteBoxs  []ExecuteBox
}

type ExecuteBox struct {
	Args     []string
	Env      []string
	Dir      string
	ExitCode int
}

var (
	DefaultJsonConfiguration = JsonConfiguration{
		Name:         "sapkin",
		Target:       TargetPath(),
		BuildCommand: "go build -o sapkin",
		BuildOutput:  "sapkin",
		ExecuteBoxs: []ExecuteBox{
			{
				Args:     []string{},
				Env:      []string{},
				Dir:      "",
				ExitCode: 0,
			},
			{
				Args:     []string{"-f", "config.toml"},
				Env:      []string{},
				Dir:      "",
				ExitCode: 0,
			},
		},
	}

	UserJsonConfiguration = JsonConfiguration{}
)

func TargetPath() string {
	// temp文件夹下无法进行go build
	return "T:\\sapkin"
}

func CreateDefaultJson() {
	// TODO : Create a default json file
	byte_json, err := json.Marshal(DefaultJsonConfiguration)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(JsonConfigurationFile, byte_json, 0777)
	if err != nil {
		panic(err)
	}
}

func LoadJsonConfiguration() {
	// TODO : Load the json file
	if !util.IsExist(JsonConfigurationFile) {
		CreateDefaultJson()
		println("Create a default json file")
		os.Exit(1)
	}

	fileByte, err := os.ReadFile(JsonConfigurationFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileByte, &UserJsonConfiguration)
	if err != nil {
		panic(err)
	}
}
