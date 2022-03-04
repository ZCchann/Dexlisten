package log

import (
	"fmt"
	"log"
	"os"
)

func Debug(v ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		log.Print("DEBUG ", fmt.Sprintln(v...))
	}
}

func Info(v ...interface{}) {
	log.Print("INFO ", fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	log.Print("ERROR ", fmt.Sprintln(v...))
}