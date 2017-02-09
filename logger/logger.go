package logger

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"strings"
)

type Logger interface {
	Debug(...interface{})
	DebugEnabled() bool
	Info(...interface{})
	InfoEnabled() bool
	Warn(...interface{})
	WarnEnabled() bool
	Error(...interface{})
	ErrorEnabled() bool
	SetLogLevel(Level)
}
type Level int

var loggerMap = make(map[string]interface{})

const (
	Debug Level = iota
	Info
	Warn
	Error
)

type FlogoLogger struct {
	loggerName string
	loggerImpl *logrus.Logger
}

const ()

func init() {

}

type FlogoFormatter struct {
	loggerName string
}

func (f *FlogoFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	logEntry := fmt.Sprintf("%s %-6s [%s] - %s\n", entry.Time.Format("2006-01-02 15:04:05.000000"), getLevel(entry.Level), f.loggerName, strings.TrimPrefix(strings.TrimSuffix(entry.Message, "]"), "["))
	return []byte(logEntry), nil
}

func getLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.ErrorLevel:
		return "ERROR"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.PanicLevel:
		return "PANIC"
	case logrus.FatalLevel:
		return "FATAL"
	}

	return "UNKNOWN"
}

// Debugf logs message at Debug level.
func (logger *FlogoLogger) Debug(args ...interface{}) {
	logger.loggerImpl.Debug(args)
}

// DebugEnabled checks if Debug level is enabled.
func (logger *FlogoLogger) DebugEnabled() bool {
	return logger.loggerImpl.Level >= logrus.DebugLevel
}

// Infof logs message at Info level.
func (logger *FlogoLogger) Info(args ...interface{}) {
	logger.loggerImpl.Info(args)
}

// InfoEnabled checks if Info level is enabled.
func (logger *FlogoLogger) InfoEnabled() bool {
	return logger.loggerImpl.Level >= logrus.InfoLevel
}

// Warnf logs message at Warning level.
func (logger *FlogoLogger) Warn(args ...interface{}) {
	logger.loggerImpl.Warn(args)
}

// WarnEnabled checks if Warning level is enabled.
func (logger *FlogoLogger) WarnEnabled() bool {
	return logger.loggerImpl.Level >= logrus.WarnLevel
}

// Errorf logs message at Error level.
func (logger *FlogoLogger) Error(args ...interface{}) {
	logger.loggerImpl.Error(args)
}

// ErrorEnabled checks if Error level is enabled.
func (logger *FlogoLogger) ErrorEnabled() bool {
	return logger.loggerImpl.Level >= logrus.ErrorLevel
}

//SetLog Level
func (logger *FlogoLogger) SetLogLevel(logLevel Level) {
	switch logLevel {
	case Debug:
		logger.loggerImpl.Level = logrus.DebugLevel
	case Info:
		logger.loggerImpl.Level = logrus.InfoLevel
	case Error:
		logger.loggerImpl.Level = logrus.ErrorLevel
	case Warn:
		logger.loggerImpl.Level = logrus.WarnLevel
	default:
		logger.loggerImpl.Level = logrus.ErrorLevel
	}
}

func GetLogger(name string) Logger {
	logger := loggerMap[name]
	if logger == nil {
		logImpl := logrus.New()
		logImpl.Formatter = &FlogoFormatter{
			loggerName: name,
		}
		logger = &FlogoLogger{
			loggerName: name,
			loggerImpl: logImpl,
		}
		loggerMap[name] = logger
	}
	return logger.(Logger)
}
