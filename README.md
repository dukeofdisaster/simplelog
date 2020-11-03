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

## TODO
- use ISO timestamps
- make a better way to log instead of passing the logger each time
