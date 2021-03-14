package logger_test

import (
	"errors"
	"fmt"
	"github.com/dukeofdisaster/simplelog/pkg/logger"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

var badPaths = []struct {
	// input
	in string
	// expected result
	expected []string
}{
	{"/noexist", []string{"read-only", "denied", "Shouldn't"}},
	{"/", []string{"read-only", "denied", "Shouldn't"}},
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
	stdoutPath := "/tmp/simplelog.test"
	err := logger.SetLoggerUtc(stdoutPath)
	if err != nil {
		t.Errorf("unable to use val: %s, as path to stdout, got err: %v", stdoutPath, err)
	}
	logger.Inf("test stdout")
	f, err := os.Open(stdoutPath)
	defer f.Close()
	if err != nil {
		t.Errorf("unable to open stdout path as file")
	}
	buf_four := make([]byte, 4)
	_, err = io.ReadAtLeast(f, buf_four, 4)
	if string(buf_four) != current_year {
		t.Errorf("expected bytes not in output: %s", buf_four)
	}
}
