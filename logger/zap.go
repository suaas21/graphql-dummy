package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func NewZapLogger(level zapcore.Level) Logger {
	return &zapLogger{newZapLogger(level)}
}

func (l zapLogger) Print(v ...interface{}) {
	defer l.sugaredLogger.Sync()
	l.sugaredLogger.Info(v)
}
func (l zapLogger) Println(v ...interface{}) {
	defer l.sugaredLogger.Sync()
	l.sugaredLogger.Info(v)
}
func (l zapLogger) Printf(format string, args ...interface{}) {
	defer l.sugaredLogger.Sync()
	l.sugaredLogger.Infof(format, args)
}

func (l zapLogger) Errorf(format string, args ...interface{}) {
	defer l.sugaredLogger.Sync()
	l.sugaredLogger.Errorf(format, args)
}

type LogLevel string

const (
	Info  LogLevel = "Info"
	Warn  LogLevel = "Warn"
	Debug LogLevel = "Debug"
	Error LogLevel = "Error"
	Fatal LogLevel = "Fatal"
)

type zapLevelLogger struct{}

func NewZapLevelLogger() LaveledLogger {
	return &zapLevelLogger{}
}
func (l zapLevelLogger) Println(v ...interface{}) {
	lgr := newZapLogger(zapcore.InfoLevel)
	defer lgr.Sync()
	lgr.Info(v)
}

func (l zapLevelLogger) Warnln(v ...interface{}) {
	lgr := newZapLogger(zapcore.WarnLevel)
	defer lgr.Sync()
	lgr.Warn(v)
}

func (l zapLevelLogger) Errorln(v ...interface{}) {
	lgr := newZapLogger(zapcore.ErrorLevel)
	defer lgr.Sync()
	lgr.Error(v)
}

func (l zapLevelLogger) Print(level LogLevel, v ...interface{}) {
	lgr := newZapLogger(getZapLevel(level))
	defer lgr.Sync()
	lgr.Info(v)
}
func (l zapLevelLogger) Printf(format string, v ...interface{}) {
	lgr := newZapLogger(zapcore.InfoLevel)
	defer lgr.Sync()
	lgr.Infof(format, v)
}

func (l zapLevelLogger) Errorf(level LogLevel, format string, v ...interface{}) {
	lgr := newZapLogger(getZapLevel(level))
	defer lgr.Sync()
	lgr.Errorf(format, v)
}

func getZapLevel(level LogLevel) zapcore.Level {
	switch level {
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Debug:
		return zapcore.DebugLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.TimeKey = "time"
	return zapcore.NewJSONEncoder(cfg)
}

func newZapLogger(level zapcore.Level) *zap.SugaredLogger {
	core := zapcore.NewCore(getEncoder(), zapcore.Lock(os.Stdout), level)
	logger := zap.New(core, zap.AddCallerSkip(1), zap.AddCaller()).Sugar()
	return logger
}
