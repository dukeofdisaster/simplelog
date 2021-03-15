package logger_test

import (
	//"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dukeofdisaster/simplelog/pkg/logger"
)

var badPaths = []struct {
	// input
	in string
	// expected result
	expected []string
}{
	{"/noexist", []string{"read-only", "denied", "wont"}},
	{"/", []string{"read-only", "denied", "wont"}},
	{"/noexist/noexist/test.log", []string{"no such"}},
}

func TestSetLoggerUtc_ShouldReturnErrs(t *testing.T) {
	for _, testcase := range badPaths {
		// all badPaths should generate an error
		actual := logger.SetLoggerUtc(testcase.in)
		if actual == nil {
			t.Errorf("expected SetLoggerUtc to return an error, but actual == nil")
			return
		}
		// test case should contain at least one of the expected strings
		good := 0
		for _, s := range testcase.expected {
			if strings.Contains(actual.Error(), s) {
				good += 1
			}
		}
		if good < 1 {
			t.Errorf("Expected one of these strings: %s, in err: %v", testcase.expected, actual)
		}
	}
}

func TestSetLoggerUtc_InfWithBadPathShouldReturnErr(t *testing.T) {
	for _, testcase := range badPaths {
		err := logger.SetLoggerUtc(testcase.in)
		if err == nil {
			t.Errorf("Expected badpath from setlogger utc")
			err = logger.Inf("DUDE")
			if err == nil {
				t.Errorf("Expected Inf() to return error but it didnt")
			}
		}
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
	err := logger.Err(errors.New("test"))
	if err == nil {
		t.Errorf("Expected call to Err() before SetLoggerUtc() to return error")
	}
}

func TestInf_ShouldContainInfInBrackets(t *testing.T) {
	current_year := fmt.Sprintf("%v", time.Now().Year())
	logPath := "/tmp/simplelog.test"
	err := logger.SetLoggerUtc(logPath)
	if err != nil {
		t.Errorf("unable to use val: %s, as path to stdout, got err: %v", logPath, err)
	}
	logger.Inf("test stdout")
	f, err := os.Open(logPath)
	defer f.Close()
	if err != nil {
		t.Errorf("unable to open stdout path as file")
	}
	buf := make([]byte, 45)
	_, err = io.ReadAtLeast(f, buf, 45)
	if !strings.Contains(string(buf), current_year) || !strings.Contains(string(buf), "[INFO]") {
		t.Errorf("expected bytes not in output: %s", buf)
	}
}
