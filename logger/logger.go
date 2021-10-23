package logger

import "go.uber.org/zap/zapcore"

// Logger is an interface to print logs
type Logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	Errorf(format string, args ...interface{})
}

type LaveledLogger interface {
	Println(v ...interface{})
	Warnln(v ...interface{})
	Errorln(v ...interface{})
	Printf(format string, args ...interface{})
	Print(level LogLevel, v ...interface{})
	Errorf(level LogLevel, format string, args ...interface{})
}

type StructLogger interface {
	Println(fn, tid string, msg string)
	Printf(fn, tid string, format string, args ...interface{})
	Warnln(fn, tid string, msg string)
	Errorln(fn, tid string, msg string)
	Errorf(fn, tid string, format string, args ...interface{})
	Print(level LogLevel, fn, tid string, msg string)
}

type LogStructure struct {
	TID  string `zap:"tid"`
	Line int    `zap:"line"`
	Msg  string `zap:"msg"`
}

var (
	DefaultOutLogger Logger
	DefaultErrLogger Logger

	DefaultOutLevelLogger LaveledLogger

	DefaultOutStructLogger StructLogger
)

func init() {
	DefaultOutLogger = NewZapLogger(zapcore.InfoLevel)
	DefaultErrLogger = NewZapLogger(zapcore.ErrorLevel)

	DefaultOutLevelLogger = NewZapLevelLogger()
	DefaultOutStructLogger = NewZeroLevelLogger()
}
