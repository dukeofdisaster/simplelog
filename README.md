# simplelog
I just needed a dumb logger without any fancy bells and whistles. 

## Sample Usage

```go

import (
	logger "github.com/dukeofdisaster/simplelog/pkg/logger"
)
var MY_LOG  = "/tmp/testlog.log"
func main() {
	l := logger.GetLogger(MY_LOG)
	logger.Info(l,"Hello from import")
	logger.Warn(l,"WARNING FROM IMPORT")
	logger.Debug(l,"DEBUG FROM IMPORT")

}
```

- the above gives the following log file

```
2020/11/03 16:16:45 INFO: Hello from import
2020/11/03 16:16:45 WARN: WARNING FROM IMPORT
2020/11/03 16:16:45 DBUG: DEBUG FROM IMPORT
```

## Sample without passing around a log.Logger object
```go

import (
	logger "github.com/dukeofdisaster/simplelog/pkg/logger"
)
var MY_LOG  = "/tmp/testlog.log"
func main() {
	logger.SetLogger(MY_LOG)
	logger.Infos("Hello from import")
	logger.Warns("WARNING FROM IMPORT")
	logger.Debugs("DEBUG FROM IMPORT")
}

```

## Sample with IOS8601 Timestamps
```go
package main

import (
	"fmt"
	logger "github.com/dukeofdisaster/simplelog/pkg/logger"
)

func main() {
	logger.SetLoggerUtc("/var/log/wiplogs/curator.log")
	logger.Inf("TEST LOG")
	logger.Wrn("warning message")
	logger.Dbg("debug message")
}
```
- the above gives the following log file
```
2021-03-11T01:44:12.47Z [INFO] TEST LOG
2021-03-11T01:44:12.47Z [WARN] warning message
2021-03-11T01:44:12.47Z [DEBUG] debug message
```

## Testing
Current testing implemented
```
user@box:~/gitstuff/simplelog/tests$
/home/user/gitstuff/simplelog/tests
user@box:~/gitstuff/simplelog/tests$ go test -v 
=== RUN   TestSetLoggerUtc_ShouldReturnErrs
--- PASS: TestSetLoggerUtc_ShouldReturnErrs (0.00s)
=== RUN   TestSetLoggerUtc_InfWithBadPathShouldReturnErr
--- PASS: TestSetLoggerUtc_InfWithBadPathShouldReturnErr (0.00s)
=== RUN   TestInf_ValidInputWithNoSetLoggerUtcShouldReturnErr
--- PASS: TestInf_ValidInputWithNoSetLoggerUtcShouldReturnErr (0.00s)
=== RUN   TestWrn_ValdidInputWithNoSetLoggerUtcShouldReturnErr
--- PASS: TestWrn_ValdidInputWithNoSetLoggerUtcShouldReturnErr (0.00s)
=== RUN   TestDbg_ValdidInputWithNoSetLoggerUtcShouldReturnErr
--- PASS: TestDbg_ValdidInputWithNoSetLoggerUtcShouldReturnErr (0.00s)
=== RUN   TestErr_ValdidInputWithNoSetLoggerUtcShouldReturnErr
--- PASS: TestErr_ValdidInputWithNoSetLoggerUtcShouldReturnErr (0.00s)
=== RUN   TestInf_ShouldContainInfInBrackets
--- PASS: TestInf_ShouldContainInfInBrackets (0.00s)
PASS
ok  	/home/user/gitstuff/simplelog/tests	0.003s
```

## TODO
- DITCH SHORTFILE?
    - after using this in a different project I realized the shortfile behavior
    is undesirable when using shortfile; with logger.Inf()/Wrn()/Ddbg()/Err()
    the line number that shows up is the the line number from the function
    declaration in logger.go
    - if the SetLoggerUtc() is used, and you call log.Println() from the
    std lib, the ISO8601 timestamp is still used and the line # is then
    the correct one, i.e. the location of the call in main.go or wherever.
    - so it seems the Inf()/Wrn()/Dbg()/Err() functions are pointless at this
    point in time if the desire is to have line numbers relevant to your
    actual sorce code. 
    - The only way forward seems like maybe setLoggerUtcStd() and then
    keeping the functions, but then checking the flags, return error if
    calling Inf()/Warn()/Dbg()/Err() 
    - see example below
```
current:
--- some src file with line numbers ---
23 logger.SetLoggerUtc("some/file/path")
24 // here we expect line 25 to appear in the log
25 loggern.Inf("test")
26 log.Println("[INFO] test")
---


gives:
2021-03-19T17:30:13.072Z logger.go:81: [INFO] test
2021-03-19T17:30:13.721Z main.go:16: [INFO] test
```
    - so any use of shortfile without direct calls to log.Println() (if you expect the
    line number in the log to be relevant to the flow in main, 
    say after some err or important process) is pointless
    - the same is true even if you use a local wrapper; the line number in the log
    could be 10s or 100s of lines away from where its relevant
```
example:
--- some src file with line numbers ---
30 func myWarningWrapper(s string) {
31    log.Println(s)
32 }
        --- snip ---
78 // somewhere in main
29 myWarningWrapper("hello world")


gives:
2021-03-19T17:30:13.721Z main.go:31: [INFO] test
---
```
- other time formats
- cosider making utc logging functions panic on err; err := logger.Inf() is an annoying pattern,
and there'es potential to mask the bug, however std log doesn't lib
doesn't return any errors, so could go either way.

