package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const defaultLogFile = "logger.log"

// Init instantiates the logger with the specified arguments.
// The first `args` parameter is the logger filename (default: logger.log).
//
// IMPORTANT: The logger will log to stderr and to the file.
func Init(serviceName, contextName string, args ...string) Logger {
	var filename string
	if len(args) == 0 || args[0] == "" {
		filename = defaultLogFile
	} else {
		filename = args[0]
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
	l.shared = strings.Join(ctxs, " ")

	// stderr
	l.dbug.stderr = log.New(os.Stderr, LvlDebug.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l.info.stderr = log.New(os.Stderr, LvlInfo.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l.warn.stderr = log.New(os.Stderr, LvlWarn.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l.eror.stderr = log.New(os.Stderr, LvlError.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l.crit.stderr = log.New(os.Stderr, LvlCrit.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile)

	// file
	file, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	f = file

	l.dbug.file = log.New(f, LvlDebug.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
	l.info.file = log.New(f, LvlInfo.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
	l.warn.file = log.New(f, LvlWarn.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
	l.eror.file = log.New(f, LvlError.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
	l.crit.file = log.New(f, LvlCrit.String(), log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
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
	logs := []string{l.shared}
	for i := 0; i < len(ctx); i += 2 {
		logs = append(logs, fmt.Sprintf("%s=\"%v\"", ctx[i], ctx[i+1]))
	}
	logs = append(logs, msg)
	lg := strings.Join(logs, " ")
	switch lvl {
	case LvlDebug:
		l.dbug.stderr.Println(lg)
		l.dbug.file.Println(lg)
	case LvlInfo:
		l.info.stderr.Println(lg)
		l.info.file.Println(lg)
	case LvlWarn:
		l.warn.stderr.Println(lg)
		l.warn.file.Println(lg)
	case LvlError:
		l.eror.stderr.Println(lg)
		l.eror.file.Println(lg)
	case LvlCrit:
		l.crit.stderr.Println(lg)
		l.crit.file.Println(lg)
	}
}

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
