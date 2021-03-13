package logger

import (
    "fmt"
    "errors"
    "io"
    "log"
    "os"
    "path"
    "syscall"
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
        return 1, err
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

func SetLogger(filepath string) error {
    f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE |os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    CURRENT_LOGGER = log.New(f,"",log.LstdFlags)
    return nil
}
/*
    Amounts to SetOutput, which doesn't return err in std lib
*/
func SetLoggerUtc(filepath string) error {
    dir := path.Dir(filepath)
    err := syscall.Access(dir, syscall.O_RDWR)
    if err != nil {
        return err
    }
    err = syscall.Access(filepath, syscall.O_RDWR)
    if err != nil {
        return err
    }
    UTC_LOGGER = new(LocalWriter)
    UTC_LOGGER.DestLog = filepath
    log.SetFlags(log.Lshortfile)
    log.SetOutput(UTC_LOGGER)
    return nil
}


func Inf(msg string) error {
    if UTC_LOGGER != nil {
        log.Println("[INFO] "+args...)
        return nil
    } else {
        return errors.New(ERR_NIL_UTC_LOGGER)
    }
}

func Wrn(msg string) error {
    if UTC_LOGGER != nil {
        log.Println("[WARN] "+msg)
        return nil
    } else {
        return errors.New(ERR_NIL_UTC_LOGGER)
    }
}

func Dbg(msg string) error {
    if UTC_LOGGER != nil {
        log.Println("[DEBUG] "+msg)
        return nil
    } else {
        return errors.New(ERR_NIL_UTC_LOGGER)
    }
}

func Err(e error) error {
    if UTC_LOGGER != nil {
        log.Println("[ERROR] "+e.Error())
        return nil
    } else {
        return errors.New(ERR_NIL_UTC_LOGGER)
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

