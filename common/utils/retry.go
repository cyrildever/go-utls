package utils

import (
	"os"
	"time"

	"github.com/cyrildever/go-utls/common/logger"
)

// RetryCallback ...
type RetryCallback func(interface{}) bool

// Retry is a function supposed to be called in a goroutine to execute a callback a number of time with a given time between each repetition.
// duration is in Millisecond.
func Retry(numberOfRetry int, duration int, callback RetryCallback, args interface{}, needToExit bool, then RetryCallback, thenArgs interface{}) {
	log := logger.Init("utils", "Retry")
	if numberOfRetry > 0 && duration > 0 {
		ticker := time.NewTicker(time.Duration(duration) * time.Millisecond)
		initNumber := numberOfRetry
		for numberOfRetry > 0 {
			<-ticker.C
			if callback(args) {
				return
			}
			numberOfRetry = numberOfRetry - 1
		}
		if then != nil {
			then(thenArgs)
		}
		if needToExit {
			log.Error("Retry failed", "numberOfRetry", initNumber)
			os.Exit(2)
		}
	} else {
		log.Error("Invalid parameters", "numberOfRetry", numberOfRetry, "duration", duration)
	}
}
