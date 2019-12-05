package retry

import (
	"log"
	"time"
)

// Retry function. fn is the retry function.
// Refer to https://github.com/mnhkahn/gogogo/blob/master/util/retry_util.go
func Retry(retryCnt int, sleep time.Duration, fn func() error) (err error) {
	if err = fn(); err == nil {
		return
	}

	for i := 1; i <= retryCnt; i++ {
		var printLog string
		if i == retryCnt {
			printLog = "Retry func error: %s. retry #%d after %s. retry done !!!"
		} else {
			printLog = "Retry func error: %s. retry #%d after %s."
		}
		log.Printf(printLog, err.Error(), i, time.Duration(i)*sleep)

		if err = fn(); err == nil {
			return
		}
		if s, ok := err.(stop); ok {
			return s.error
		}
		time.Sleep(time.Duration(i) * sleep)
	}
	return
}

type stop struct {
	error
}

// NoRetryError ...
func NoRetryError(err error) stop {
	return stop{err}
}
