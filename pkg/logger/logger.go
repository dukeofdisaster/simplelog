package logger

import (
	"fmt"
	"os"
	"log"
)
// GLOBALS
var DEFAULT = "/var/log/simplelogger.log"
var CURRENT_LOGGER *log.Logger = nil

func GetLogger(filepath string) (*log.Logger) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE |os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return(log.New(f,"",log.LstdFlags))
}

func SetLogger(filepath string) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE |os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
    CURRENT_LOGGER = log.New(f,"",log.LstdFlags)
}

/* Log with global logger after calling SetLogger */
func Infos(msg string) {
    if CURRENT_LOGGER !=  nil {
        CURRENT_LOGGER.Println("INFO: "+msg)
    } else {
        panic("Attempted to log to nil logger; did you run logger.SetLogger() ?")
    }
}
func Warns(msg string) {
    if CURRENT_LOGGER != nil {
        CURRENT_LOGGER.Println("WARN: "+msg)
    } else {
        panic("Attempted to log to nil logger; did you run loger.SetLogger() ?")
    }
}
func Debugs(msg string) {
    if CURRENT_LOGGER != nil {
        CURRENT_LOGGER.Println("DEBUG: "+msg)
    } else {
        panic("Attempted to log to nil logger; did you run logger.SetLogger() ?")
    }
}

/* Log with pointer */
func Info(someLogger *log.Logger,msg string) {
	someLogger.Println("INFO: "+msg)
}
func Warn(someLogger *log.Logger, msg string) {
	someLogger.Println("WARN: "+msg)
}
func Debug(someLogger *log.Logger, msg string) {
	someLogger.Println("DBUG: "+msg)
}

