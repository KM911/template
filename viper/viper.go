package main

import (
	"github.com/spf13/viper"
	"reflect"
	"viper/config"
)

func CheckType(v interface{}) {
	println("type of v is ", reflect.TypeOf(v))
}

func InterfaceToString(v interface{}) string {
	return reflect.ValueOf(v).String()
}
func InterfaceToInt(v interface{}) int {
	return int(reflect.ValueOf(v).Int())
}
func InterfaceToFloat64(v interface{}) float64 {
	return float64(reflect.ValueOf(v).Int())
}

func main() {
	// viper demo
	// set config file path
	config.INIt()
	//
	//"大小写不敏感"
	//Vscodpath := viper.GetString("VscodPath")
	//println(Vscodpath)

	// 就是不是很方便不是吗就是说 其实还是不错的喝喝
	shortcut := viper.GetStringMap("shortcut")
	for key, value := range shortcut {
		println(key)
		for _, value := range value.([]interface{}) {
			println(InterfaceToString(value))
		}
	}

	// write config file

	// 这样就可以写会去了 保存成功了就是说
	viper.Set("shortcuttest", map[string]interface{}{
		"vscode":  []interface{}{"vscode", "vsc"},
		"chrome":  []interface{}{"chrome", "c", "google"},
		"pycharm": []interface{}{"pycharm", "p"},
	})

}
