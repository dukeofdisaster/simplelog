package logger

import (
	"fmt"
	"os"
	"log"
)
// GLOBALS
var DEFAULT = "/var/log/simplelogger.log"
func GetLogger(filepath string) (*log.Logger) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE |os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return(log.New(f,"",log.LstdFlags))
}

/* LOGGING FUNCTINOALITY */
func Info(someLogger *log.Logger,msg string) {
	someLogger.Println("INFO: "+msg)
}
func Warn(someLogger *log.Logger, msg string) {
	someLogger.Println("WARN: "+msg)
}
func Debug(someLogger *log.Logger, msg string) {
	someLogger.Println("DBUG: "+msg)
}

