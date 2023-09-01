package main

import (
	"fmt"
	"github.com/KM911/oslib/adt"
	"github.com/KM911/oslib/fs"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
)

func FileLogger(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

type ErrorMessage struct {
	Type int
	Msg  string
}

func EmitError(n int, msg string) {
	panic(ErrorMessage{
		Type: n,
		Msg:  msg,
	})
}

func Recover(errorHandler func(ErrorMessage)) {
	if errMsg := recover(); errMsg != nil {
		errorHandler(errMsg.(ErrorMessage))
		log.Println(errMsg.(ErrorMessage).Msg)
		log.Println(string(debug.Stack()))
	}
}

func ErrorHandler(err ErrorMessage) {
	switch err.Type {
	case 1:
		fmt.Println("Error type 1")
		fmt.Println(err.Msg)

	default:
		println("Unknown error type")
	}
}

func main() {
	FileLogger(filepath.Join(fs.ExecutePath(), "log.log"))
	defer adt.TimerStart().End()
	defer Recover(ErrorHandler)

}
