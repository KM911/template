package test

import (
	"github.com/KM911/oslib"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func DelShortcutKey(key string) {
	newvalue := viper.GetStringMapStringSlice("shortcut")
	for i, v := range newvalue {
		if v == nil {
			newvalue[i] = []string{}
		}
	}
	delete(newvalue, key)
	viper.Set("shortcut", newvalue)
	viper.WriteConfig()
}
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
		shortcutList[i] = strings.ToLower(oslib.FileName(shortcutList[i]))
	}
	for key, _ := range viper.GetStringMapStringSlice("shortcut") {
		if !oslib.InArray(key, shortcutList) {
			t.Error("shortcut " + key + " not exist")
			DelShortcutKey(key)
		}
	}

}

// 检查快捷键是否重复
func Test_shortcut_is_repeate(t *testing.T) {
	viper.SetConfigFile(filepath.Join(path.Dir(oslib.CmdPath()), "config", "config.json"))
	viper.ReadInConfig()
	//shortcutList := oslib.Dir(filepath.Join(path.Dir(oslib.CmdPath()), "config", "shortcut"))
	for key, _ := range viper.GetStringMapStringSlice("shortcut") {
		for _, value := range viper.GetStringMapStringSlice("shortcut") {
			for _, v := range value {
				// key == v 说明快捷键重复
				if key == v {
					t.Error("shortcut " + key + " is repeate")
					// only remove the v
					//newvalue := value

					break
				}

			}
		}
	}
}
