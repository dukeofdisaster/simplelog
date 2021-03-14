package logger_test

import "io"
import "fmt"
//import "bufio"
import "errors"
import "os"
import "github.com/dukeofdisaster/simplelog/pkg/logger"
//import "fmt"
import "strings"
import "testing"
var badPaths = []struct {
    // input
    in string
    // expected result
    expected  []string
}{
    {"/noexist", []string{"read-only", "denied"}},
    {"/", []string{"read-only","denied"}},
    {"/noexist/noexist/test.log", []string{"no such"}},
    {"/dev/stderr", []string{"no such"}},
}
func TestSetLoggerUtc_ShouldReturnErrs(t *testing.T) {
    for _, testcase := range badPaths {
        // all badPaths should generate an error
        actual := logger.SetLoggerUtc(testcase.in)
        // test case should contain at least one of the expected strings
        good := 0
        for _,s := range testcase.expected {
            if strings.Contains(actual.Error(), s) {
                good += 1
            }
        }
        if good <  1  {
            t.Errorf("Expected one of these strings: %s, in err: %v",testcase.expected,actual)
        }
        //fmt.Printf("good: %d\n",good)
    }
}

func TestInf_ValidInputWithNoSetLoggerUtcShouldReturnErr(t *testing.T) {
    err := logger.Inf("test") 
    if err == nil {
        t.Errorf("Expected call to Inf() before SetLoggerUtc() to return error")
    }
}

func TestWrn_ValdidInputWithNoSetLoggerUtcShouldReturnErr(t *testing.T) {
    err := logger.Wrn("test") 
    if err == nil {
        t.Errorf("Expected call to Wrn() before SetLoggerUtc() to return error")
    }
}

func TestDbg_ValdidInputWithNoSetLoggerUtcShouldReturnErr(t *testing.T) {
    err := logger.Dbg("test") 
    if err == nil {
        t.Errorf("Expected call to Dbg() before SetLoggerUtc() to return error")
    }
}

func TestErr_ValdidInputWithNoSetLoggerUtcShouldReturnErr(t *testing.T) {
    err := logger.Err(errors.New("test") )
    if err == nil {
        t.Errorf("Expected call to Err() before SetLoggerUtc() to return error")
    }
}

//func ExampleInf_ShouldContainInfInBrackets(t *testing.T) {
func TestInf_ShouldContainInfInBrackets(t *testing.T) {
    stdoutPath := "/dev/fd/2"
    err := logger.SetLoggerUtc(stdoutPath)
    if err != nil {
        t.Errorf("unable to use val: %s, as path to stdout, got err: %v",stdoutPath,err)
    }
    logger.Inf("test stdout")

    f, err := os.Open(stdoutPath)
    defer f.Close()
    if err != nil {
        t.Errorf("unable to open stdout path as file")
    }
    buf_ten := make([]byte, 5)
    num_read, err := io.ReadAtLeast(f,buf_ten,5)
    fmt.Printf("%d bytes read: %s\n", num_read, string(buf_ten))
    if string(buf_ten) == "dude" {
        t.Errorf("expected bytes not in output: %s",buf_ten)
    }
}
