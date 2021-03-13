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
## TODO
- other time formats
- add tests
    - what happens when we call a utc logging function (Inf(),Dbg(),Wrn()) with non-utc local writer?
- add error logging
    - should probably take an error directly
