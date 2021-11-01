package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type zeroLevelLogger struct {
	lgr zerolog.Logger
}

func NewZeroLevelLogger() StructLogger {
	sublogger := log.With().CallerWithSkipFrameCount(3).Stack().
		Str("service", "catalog").
		Logger()
	return &zeroLevelLogger{
		lgr: sublogger,
	}
}
func (l zeroLevelLogger) Println(fn, tid string, msg string) {
	l.lgr.Info().Str("function", fn).Str("tid", tid).Msg(msg)
}

func (l zeroLevelLogger) Printf(fn, tid string, format string, args ...interface{}) {
	l.lgr.Info().Str("function", fn).Str("tid", tid).Msgf(format, args...)
}

func (l zeroLevelLogger) Warnln(fn, tid string, msg string) {
	l.lgr.Warn().Str("function", fn).Str("tid", tid).Msg(msg)
}

func (l zeroLevelLogger) Errorln(fn, tid string, msg string) {
	l.lgr.Error().Str("function", fn).Str("tid", tid).Msg(msg)
}
func (l zeroLevelLogger) Errorf(fn, tid string, format string, args ...interface{}) {
	l.lgr.Error().Str("function", fn).Str("tid", tid).Msgf(format, args...)
}

func (l zeroLevelLogger) Print(level LogLevel, fn, tid string, msg string) {
	l.lgr.Log().Str("level", getZeroLevel(level).String()).Str("function", fn).Str("tid", tid).Msg(msg)
}

func getZeroLevel(level LogLevel) zerolog.Level {
	switch level {
	case Info:
		return zerolog.InfoLevel
	case Warn:
		return zerolog.WarnLevel
	case Debug:
		return zerolog.DebugLevel
	case Error:
		return zerolog.ErrorLevel
	case Fatal:
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
