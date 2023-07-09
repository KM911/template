package config

import (
	"log"
	"os"
)

func EnableLog(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func Info(info string) {
	log.Println("INFO: " + info)
}

func Error(err string) {
	log.Println("ERROR: " + err)
}
