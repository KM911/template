package command

import (
	"cli/config"
	"github.com/KM911/oslib"
	"github.com/spf13/viper"
	"path/filepath"
)

func Start(cmd string) {
	ShortcutMap := viper.GetStringMapStringSlice("shortcut")
	target := ""
	for key, value := range ShortcutMap {
		if key == cmd {
			target = key
		}
		for _, v := range value {
			if v == cmd {
				target = key
			}
		}
	}
	if target == "" {
		println("command not found")
	} else {
		oslib.Run("start " + filepath.Join(config.ShortcutDir, target))
	}
}
