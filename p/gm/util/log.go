package util

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gookit/color"
)

// use color to beautify the log

//OFF、FATAL、ERROR、WARN、INFO、DEBUG、TRACE、 ALL。
// Success green
// Debug white
// Info ray
// Warning Yellow
// FATAL Red

var ()

// default_logger
func FileLogger(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	// 将os.Stdout和logFile都设置为输出目的地
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	//
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)

}

// ERROR: error prompt message
// WARNING: warning prompt message
// INFO: info prompt message
// NOTE: note prompt message

// [tag] : msg
func LogErorr(msg string) {
	log.Println(color.BgRed.Render("[Error]"), ": ", color.Error.Render(msg))
}

func Error(msg string) {
	fmt.Println(color.BgRed.Render("[Error]"), ": ", color.Error.Render(msg))
}

func ErrorMessage(error string, msg string) {
	fmt.Println(color.BgRed.Render(error), ": ", color.Error.Render(msg))
}

func LogWarning(msg string) {
	log.Println(color.BgYellow.Render("[Warning]"), ": ", color.Warn.Render(msg))
}

func Warning(msg string) {
	fmt.Println(color.BgYellow.Render("[Warning]"), ": ", color.Warn.Render(msg))

}

func WarningMessage(warning string, msg string) {
	fmt.Println(color.BgYellow.Render(warning), ": ", color.Warn.Render(msg))
}

func LogInfo(msg string) {
	log.Println(color.BgBlue.Render("[Info]"), ": ", color.Info.Render(msg))
}

func Info(msg string) {
	fmt.Println(color.BgBlue.Render("[Info]"), ": ", color.Info.Render(msg))
}

func InfoMessage(info string, msg string) {
	fmt.Println(color.BgBlue.Render(info), ": ", color.Info.Render(msg))
}

func LogNote(msg string) {
	log.Println(color.BgHiBlue.Render("[Note]"), ": ", color.Note.Render(msg))
}

func Note(msg string) {
	NoteMessage("Note", msg)
}
func NoteMessage(note string, msg string) {
	fmt.Println(color.BgHiBlue.Render(note), ": ", color.Note.Render(msg))
}
func LogSuccess(msg string) {
	log.Println(color.BgGreen.Render("[Success]"), ": ", color.Success.Render(msg))
}

func Success(msg string) {
	fmt.Println(color.BgGreen.Render("[Success]"), ": ", color.Success.Render(msg))
}

func SuccessMessage(success string, msg string) {
	fmt.Println(color.BgGreen.Render(success), ": ", color.Success.Render(msg))
}

func LogExample() {
	LogErorr("error")
	LogWarning("warning")
	LogInfo("info")
	LogNote("note")
	LogSuccess("success")
}

func Example() {
	Error("error")
	Warning("warning")
	Info("info")
	Note("note")
	Success("success")
}
