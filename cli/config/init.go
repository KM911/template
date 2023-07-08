package config

import (
	"github.com/KM911/oslib"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

var (
	PWD         string
	ConfigFile  string
	ShortcutDir string
)

func LoadConfig() {
	SetGlobal()
	CheckEnv()
	CheckDebug()
}

/*
CheckEnv checks if the environment is ready for the cli to run
if not, it will try to fix it
*/
func SetGlobal() {
	PWD = oslib.ExecutePath()
	ConfigFile = filepath.Join(PWD, "config", "config.json")
	ShortcutDir = filepath.Join(PWD, "config", "shortcut")
}
func CheckEnv() {
	viper.SetConfigFile(ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		InitConfig()
		panic(err)
	}
}

func InitConfig() {
	println("config file not found, init config file")
	os.Mkdir(filepath.Join(PWD, "config"), os.ModePerm)
	os.Mkdir(ShortcutDir, os.ModePerm)
	viper.Set("shortcut", map[string]string{})
	// 开始移动文件夹
	PublicDesktop := filepath.Join("C:\\Users\\Public\\Desktop")
	PublicDesktopDir := oslib.Dir(PublicDesktop)
	UserDesktop := filepath.Join("C:\\Users", os.Getenv("username"), "Desktop")
	UserDesktopDir := oslib.Dir(UserDesktop)
	for _, i := range PublicDesktopDir {
		if path.Ext(i) == ".lnk" {
			println("move " + i + " to " + filepath.Join(PWD,
				"config", "shortcut", i))
			os.Rename(filepath.Join(PublicDesktop, i),
				filepath.Join(PWD, "config", "shortcut", i))
		}
	}
	for _, i := range UserDesktopDir {
		if path.Ext(i) == ".lnk" {
			println("move " + i + " to " + filepath.Join(PWD,
				"config", "shortcut", i))
			os.Rename(filepath.Join(UserDesktop, i),
				filepath.Join(PWD, "config", "shortcut", i))
		}
	}
	ShortcutList := oslib.Dir(filepath.Join(PWD, "config", "shortcut"))
	for i := 0; i < len(ShortcutList); i++ {
		viper.Set("shortcut."+oslib.FileName(ShortcutList[i]), []string{})
	}

	viper.Set("debug", false)
	viper.Set("template", "")
	//viper.Set()
	viper.WriteConfig()
}

func CheckDebug() {
	if viper.GetBool("debug") {
		println("debug mode")
	}
}
