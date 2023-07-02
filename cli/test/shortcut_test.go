package test

import (
	"github.com/KM911/oslib"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"testing"
)

func Test_ShortcutExist(t *testing.T) {
	viper.SetConfigFile(filepath.Join(path.Dir(oslib.CmdPath()), "config", "config.json"))
	viper.ReadInConfig()
	shortcutList := oslib.Dir(filepath.Join(path.Dir(oslib.CmdPath()), "config", "shortcut"))
	for _, i := range shortcutList {
		// 去掉后缀
		shortcut := oslib.FileName(i)
		value := viper.Get("shortcut." + shortcut)
		if value == nil {
			t.Error("shortcut " + shortcut + " not found")
			viper.Set("shortcut."+shortcut, []string{})
		} else {
			//println("shortcut " + shortcut + " found")
		}
	}
	viper.WriteConfig()
}

func Test_ShortcutNotExist(t *testing.T) {
	viper.SetConfigFile(filepath.Join(path.Dir(oslib.CmdPath()), "config", "config.json"))
	viper.ReadInConfig()
	shortcutList := oslib.Dir(filepath.Join(path.Dir(oslib.CmdPath()), "config", "shortcut"))
	for i := range shortcutList {
		shortcutList[i] = oslib.FileName(shortcutList[i])
	}
	for key, _ := range viper.GetStringMapStringSlice("shortcut") {
		if !oslib.InArray(key, shortcutList) {
			t.Error("shortcut " + key + " not exist")
			// remove the key
			viper.Set("shortcut."+key, nil)
		}
	}

	//viper.WriteConfig()
}

func Test_shortcut_is_repeate(t *testing.T) {
	viper.SetConfigFile(filepath.Join(path.Dir(oslib.CmdPath()), "config", "config.json"))
	viper.ReadInConfig()
	//shortcutList := oslib.Dir(filepath.Join(path.Dir(oslib.CmdPath()), "config", "shortcut"))
	for key, _ := range viper.GetStringMapStringSlice("shortcut") {
		for _, value := range viper.GetStringMapStringSlice("shortcut") {
			for _, v := range value {
				if key == v {
					t.Error("shortcut " + key + " is repeate")
					panic("shortcut " + key + " is repeate")
					//newvalue := append(value[:index], value[index+1:]...)
					//viper.Set("shortcut."+key, newvalue)
					break
				}
			}
		}
	}
}
