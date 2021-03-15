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
- other time formats
- cosider making utc logging functions panic on err; err := logger.Inf() is an annoying pattern,
and there'es potential to mask the bug, however std log doesn't lib
doesn't return any errors, so could go either way.
