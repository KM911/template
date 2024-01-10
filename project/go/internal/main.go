package internal

import (
	"demo/util"
	"os"
)

func MainProcess() {
	file, err := os.Open("test.file")
	if err != nil {
		util.EmitError(1, "open file error")
	}
	defer file.Close()
}
