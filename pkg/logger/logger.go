package logger

import (
    "fmt"
    "io"
    "log"
    "os"
    "time"
)

type LocalWriter struct {
    DestLog string
}

var ISO_8601_FMT = "2006-01-02T15:04:05.999Z"
var DEFAULT = "/var/log/simplelogger.log"
var CURRENT_LOGGER *log.Logger = nil
var UTC_LOGGER *LocalWriter = nil
var ERR_NIL_LOGGER = "Attempted to log to nil logger; did you run logger.SetLogger() ?"
var ERR_NIL_UTC_LOGGER = "Attempted to log with nil LocalWriterr; did you run logger.SetLoggerUtc() ?"

// define our own Write method for LocalWriter so we can give our own timestamp
func (writer *LocalWriter) Write(bytes []byte) (int,error) {
    f,err := os.OpenFile(writer.DestLog,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    ws := io.WriteString
    msg := string(bytes)
    return ws(f,fmt.Sprintf(time.Now().UTC().Format(ISO_8601_FMT) + " " + msg))
}

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
        panic(err)
    }
    CURRENT_LOGGER = log.New(f,"",log.LstdFlags)
}

func SetLoggerUtc(filepath string) {
   UTC_LOGGER = new(LocalWriter)
   UTC_LOGGER.DestLog = filepath
}

func Inf(msg string) {
    if UTC_LOGGER != nil {
        log.SetFlags(0)
        log.SetOutput(UTC_LOGGER)
        log.Println("[INFO] "+msg)
    } else {
        panic(ERR_NIL_UTC_LOGGER)
    } 
}

func Wrn(msg string) {
    if UTC_LOGGER != nil {
        log.SetFlags(0)
        log.SetOutput(UTC_LOGGER)
        log.Println("[WARN] "+msg)
    } else {
        panic(ERR_NIL_UTC_LOGGER)
    } 
}

func Dbg(msg string) {
    if UTC_LOGGER != nil {
        log.SetFlags(0)
        log.SetOutput(UTC_LOGGER)
        log.Println("[DEBUG] "+msg)
    } else {
        panic(ERR_NIL_UTC_LOGGER)
    } 
}
/* Log with global logger after calling SetLogger */
func Infos(msg string) {
    if CURRENT_LOGGER !=  nil {
        CURRENT_LOGGER.Println("INFO: "+msg)
    } else {
        panic(ERR_NIL_LOGGER)
    }
}
func Warns(msg string) {
    if CURRENT_LOGGER != nil {
        CURRENT_LOGGER.Println("WARN: "+msg)
    } else {
        panic(ERR_NIL_LOGGER)
    }
}
func Debugs(msg string) {
    if CURRENT_LOGGER != nil {
        CURRENT_LOGGER.Println("DEBUG: "+msg)
    } else {
        panic(ERR_NIL_LOGGER)
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

