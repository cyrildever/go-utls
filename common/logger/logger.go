package logger

import (
	"os"
	"path/filepath"

	"github.com/inconshreveable/log15"
)

const defaultLogFile = "logger.log"

// Init instantiates the logger with the specified arguments.
// The first `args` parameter is the logger filename (default: logger.log).
//
// IMPORTANT: The logger will log to stderr and to the file.
func Init(serviceName, contextName string, args ...string) log15.Logger {
	if len(args) != 1 || args[0] == "" {
		if len(args) == 0 {
			args = append(args, defaultLogFile)
		} else {
			args[0] = defaultLogFile
		}
	}
	logger := log15.New("service", serviceName, "context", contextName)
	return setHandlers(logger, args[0])
}

// InitHandler is generally used in API handlers to display the request ID.
//
// IMPORTANT: The logger will log to stderr and to the file.
func InitHandler(serviceName, contextName, requestID string, args ...string) log15.Logger {
	if len(args) != 1 || args[0] == "" {
		if len(args) == 0 {
			args = append(args, defaultLogFile)
		} else {
			args[0] = defaultLogFile
		}
	}
	logger := log15.New("service", serviceName, "context", contextName, "request_id", requestID)
	return setHandlers(logger, args[0])
}

func setHandlers(logger log15.Logger, filename string) log15.Logger {
	var handlers []log15.Handler

	// stdErr
	stdErrHandler := log15.CallerStackHandler("%+v", log15.StderrHandler)
	handlers = append(handlers, stdErrHandler)

	// file
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fileHandler := log15.Must.FileHandler(pwd+"/"+filename, log15.LogfmtFormat())
	handlers = append(handlers, fileHandler)

	logger.SetHandler(log15.MultiHandler(handlers...))
	return logger
}
