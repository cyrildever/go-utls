package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-stack/stack"
)

const defaultLogFile = "logger.log"

var useFile = true

const (
	NO_FILE = "true"
)

// Init instantiates the logger with the specified arguments.
//
// The first `args` parameter is the logger filename (default: logger.log).
//
// The second `args` parameter indicates whether to get rid of any file logging (default: "false").
// You might want to use the `logger.NO_FILE` constant as this second parameter or the string "true", eg.
//	log := logger.Init("my_package", "MyFunction", "", logger.NO_FILE)
//
// IMPORTANT: The logger will log to stderr and to the file unless stated otherwise.
func Init(serviceName, contextName string, args ...string) Logger {
	var filename string
	if len(args) == 0 || args[0] == "" {
		filename = defaultLogFile
	} else {
		filename = args[0]
	}

	if len(args) > 1 && args[1] == NO_FILE {
		useFile = false
	}

	return New(filename, "service", serviceName, "context", contextName)
}

// InitHandler is generally used in API handlers to display the request ID.
//
// IMPORTANT: The logger will log to stderr and to the file.
func InitHandler(serviceName, contextName, requestID string, args ...string) Logger {
	var filename string
	if len(args) == 0 || args[0] == "" {
		filename = defaultLogFile
	} else {
		filename = args[0]
	}
	return New(filename, "service", serviceName, "context", contextName, "request_id", requestID)
}

// Lvl is a type for predefined log levels.
type Lvl int

// List of predefined log Levels
const (
	LvlCrit Lvl = iota
	LvlError
	LvlWarn
	LvlInfo
	LvlDebug
)

// Returns the name of a Lvl
func (l Lvl) String() string {
	switch l {
	case LvlDebug:
		return "DEBUG "
	case LvlInfo:
		return "INFO  "
	case LvlWarn:
		return "WARN  "
	case LvlError:
		return "ERROR "
	case LvlCrit:
		return "CRIT  "
	default:
		panic("bad level")
	}
}

// Logger writes key/pair values to a handler
type Logger interface {
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Crit(msg string, ctx ...interface{})
}

type logger struct {
	ctx      []interface{}
	filename string
	shared   string

	dbug handler
	info handler
	warn handler
	eror handler
	crit handler
}

type handler struct {
	stderr *log.Logger
	file   *log.Logger
}

var f *os.File

func (l *logger) init() {
	var ctxs []string
	for i := 0; i < len(l.ctx); i += 2 {
		ctxs = append(ctxs, fmt.Sprintf("%s=%v", l.ctx[i], l.ctx[i+1]))
	}
	l.shared = strings.Join(ctxs, "\t")

	// stderr
	l.dbug.stderr = log.New(os.Stderr, LvlDebug.String(), log.Ldate|log.Lmicroseconds)
	l.info.stderr = log.New(os.Stderr, LvlInfo.String(), log.Ldate|log.Lmicroseconds)
	l.warn.stderr = log.New(os.Stderr, LvlWarn.String(), log.Ldate|log.Lmicroseconds)
	l.eror.stderr = log.New(os.Stderr, LvlError.String(), log.Ldate|log.Lmicroseconds)
	l.crit.stderr = log.New(os.Stderr, LvlCrit.String(), log.Ldate|log.Lmicroseconds)

	// file
	if useFile {
		file, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		f = file

		l.dbug.file = log.New(f, LvlDebug.String(), log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
		l.info.file = log.New(f, LvlInfo.String(), log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
		l.warn.file = log.New(f, LvlWarn.String(), log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
		l.eror.file = log.New(f, LvlError.String(), log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
		l.crit.file = log.New(f, LvlCrit.String(), log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
	}
}

func (l *logger) Debug(msg string, ctx ...interface{}) {
	l.write(LvlDebug, msg, ctx)
}

func (l *logger) Info(msg string, ctx ...interface{}) {
	l.write(LvlInfo, msg, ctx)
}

func (l *logger) Warn(msg string, ctx ...interface{}) {
	l.write(LvlWarn, msg, ctx)
}

func (l *logger) Error(msg string, ctx ...interface{}) {
	l.write(LvlError, msg, ctx)
}

func (l *logger) Crit(msg string, ctx ...interface{}) {
	l.write(LvlCrit, msg, ctx)
}

func (l *logger) write(lvl Lvl, msg string, ctx []interface{}) {
	logs := []string{fmt.Sprintf("msg={ %s }", msg)}
	logs = append(logs, l.shared)
	for i := 0; i < len(ctx); i += 2 {
		logs = append(logs, fmt.Sprintf("%s=\"%v\"", ctx[i], ctx[i+1]))
	}
	s := stack.Trace().TrimBelow(stack.Caller(2)).TrimRuntime()
	if len(s) > 0 {
		logs = append(logs, s.String())
	}
	lg := strings.Join(logs, "\t")
	switch lvl {
	case LvlDebug:
		l.dbug.stderr.Println(lg)
		if useFile {
			l.dbug.file.Println(lg)
		}
	case LvlInfo:
		l.info.stderr.Println(lg)
		if useFile {
			l.info.file.Println(lg)
		}
	case LvlWarn:
		l.warn.stderr.Println(lg)
		if useFile {
			l.warn.file.Println(lg)
		}
	case LvlError:
		l.eror.stderr.Println(lg)
		if useFile {
			l.eror.file.Println(lg)
		}
	case LvlCrit:
		l.crit.stderr.Println(lg)
		if useFile {
			l.crit.file.Println(lg)
		}
	}
}

//-- FUNCTIONS

// New ...
func New(filename string, ctx ...interface{}) Logger {
	if len(ctx)%2 != 0 {
		panic("wrong number of context argument")
	}
	child := &logger{
		ctx:      ctx,
		filename: filename,
	}
	child.init()
	return child
}

// GetTimeFromLine ...
func GetTimeFromLine(input string) (t time.Time) {
	tm, err := time.Parse("2006/01/02 15:04:05.999999999-07:00", input[:26]+"000+02:00")
	if err != nil {
		return
	}
	return tm
}
