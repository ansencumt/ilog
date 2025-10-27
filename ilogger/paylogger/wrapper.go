package paylogger

import (
	imlogger "github.com/ansencumt/ilog/v2/logger"

	"github.com/donnie4w/go-logger/logger"
)

func NewILogger(logMonitor *logger.Logging) imlogger.Logger {
	if logMonitor == nil {
		logMonitor = logger.NewLogger()
	}

	return &ILogger{
		logMonitor: logMonitor,
	}
}

type ILogger struct {
	logMonitor *logger.Logging
}

// Debug write the debug msg
func (gl *ILogger) Debug(format string) {
	gl.logMonitor.Debug(format)
}

// Debugf write the debug msg
func (gl *ILogger) Debugf(format string, args ...any) {
	gl.logMonitor.Debugf(format, args...)
}

// Debugln write the debug msg
func (gl *ILogger) Debugln(args ...any) {
	gl.logMonitor.Debug(args...)
}

// Trace write the trace msg
func (gl *ILogger) Trace(format string) {
	gl.logMonitor.Debug(format)
}

// Tracef write the trace msg
func (gl *ILogger) Tracef(format string, args ...any) {
	gl.logMonitor.Debugf(format, args...)
}

// Traceln write the trace msg
func (gl *ILogger) Traceln(args ...any) {
	gl.logMonitor.Debug(args...)
}

// Info write the info msg
func (gl *ILogger) Info(format string) {
	gl.logMonitor.Info(format)
}

// Infof write the info msg
func (gl *ILogger) Infof(format string, args ...any) {
	gl.logMonitor.Infof(format, args...)
}

// Infoln write the info msg
func (gl *ILogger) Infoln(args ...any) {
	gl.logMonitor.Info(args...)
}

// Warn write the warn msg
func (gl *ILogger) Warn(format string) {
	gl.logMonitor.Warn(format)
}

// Warnf write the warn msg
func (gl *ILogger) Warnf(format string, args ...any) {
	gl.logMonitor.Warnf(format, args...)
}

// Warnln write the warn msg
func (gl *ILogger) Warnln(args ...any) {
	gl.logMonitor.Warn(args...)
}

// Error write the error msg
func (gl *ILogger) Error(format string) {
	gl.logMonitor.Error(format)
}

// Errorf write the error msg
func (gl *ILogger) Errorf(format string, args ...any) {
	gl.logMonitor.Errorf(format, args...)
}

// Errorln write the error msg
func (gl *ILogger) Errorln(args ...any) {
	gl.logMonitor.Error(args...)
}

// Panic write the panic msg
func (gl *ILogger) Panic(format string) {
	gl.logMonitor.Fatal(format)
}

// Panicf write the panic msg
func (gl *ILogger) Panicf(format string, args ...any) {
	gl.logMonitor.Fatalf(format, args...)
}

// Panicln write the panic msg
func (gl *ILogger) Panicln(args ...any) {
	gl.logMonitor.Fatal(args...)
}

// SetLevel 配置logger的大小
func (gl *ILogger) SetLevel(level imlogger.LoggerLevel) {
	gl.logMonitor.SetLevel(convertLevel(level))
}

// GetLevel 获取logger的大小
func (gl *ILogger) GetLevel() imlogger.LoggerLevel {
	return imlogger.TRACE
}

func convertLevel(level imlogger.LoggerLevel) logger.LEVELTYPE {
	switch level {
	case imlogger.DEBUG:
		return logger.LEVEL_DEBUG
	case imlogger.INFO:
		return logger.LEVEL_INFO
	case imlogger.WARN:
		return logger.LEVEL_WARN
	case imlogger.ERROR:
		return logger.LEVEL_ERROR
	case imlogger.PANIC:
		return logger.LEVEL_FATAL
	default:
		return logger.LEVEL_ALL
	}
}

func convertLogLevel(level logger.LEVELTYPE) imlogger.LoggerLevel {
	switch level {
	case logger.LEVEL_DEBUG:
		return imlogger.DEBUG
	case logger.LEVEL_INFO:
		return imlogger.INFO
	case logger.LEVEL_WARN:
		return imlogger.WARN
	case logger.LEVEL_ERROR:
		return imlogger.ERROR
	case logger.LEVEL_FATAL:
		return imlogger.PANIC
	default:
		return imlogger.TRACE
	}
}
