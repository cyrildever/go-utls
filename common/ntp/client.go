package ntp

import (
	"errors"
	"math"
	"time"

	ntpServer "github.com/beevik/ntp"
	"github.com/cyrildever/go-utls/common/logger"
)

const (
	MIN_LIMIT float64 = 100. // in milliseconds
)

var (
	defaultHost string
	limit       float64
	timeServer  func(string) (time.Time, error)
)

// Initialize defines the default NTP host and maximum time leap (in milliseconds) between the current machine and the NTP server time
// NB: Must be initialized before
func Initialize(host string, timeLeapMillis float64) error {
	if host == "" {
		return errors.New("invalid empty NTP host")
	}
	defaultHost = host
	if timeLeapMillis < MIN_LIMIT {
		limit = MIN_LIMIT
	} else {
		limit = timeLeapMillis
	}
	return nil
}

// Time ...
func Time(host string) (time.Time, error) {
	if timeServer == nil {
		if host == "" {
			host = defaultHost
		}
		if host == "" || limit == 0 {
			return time.Now(), errors.New("NTP wasn't initialized")
		}
		local := time.Now()
		if server, err := ntpServer.Time(host); err == nil &&
			math.Abs(float64(local.UnixNano()/1e6)-float64(server.UnixNano()/1e6)) > limit {
			log := logger.Init("ntp", "Time")
			log.Crit("Using remote server because time leap is too high", "local", local.UnixNano(), "server", server.UnixNano(), "diff", math.Abs(float64(local.UnixNano()/1e6)-float64(server.UnixNano()/1e6)), "limit", limit)
			timeServer = ntpServer.Time
		} else {
			timeServer = localTime
		}
	}
	return timeServer(host)
}

func localTime(noHost string) (time.Time, error) {
	return time.Now(), nil
}
