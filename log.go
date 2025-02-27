package ilog

import (
	"github.com/donnie4w/go-logger/logger"
)

func NewGoLogger(logMonitor *logger.Logging) *GoLoggerWrapper {
	if logMonitor == nil {
		logMonitor = logger.NewLogger()
	}

	return &GoLoggerWrapper{
		logMonitor: logMonitor,
	}
}

type GoLoggerWrapper struct {
	logMonitor *logger.Logging
}

// Debug write the debug msg
func (gl *GoLoggerWrapper) Debug(format string) {
	gl.logMonitor.Debug(format)
}

// Debugf write the debug msg
func (gl *GoLoggerWrapper) Debugf(format string, args ...any) {
	gl.logMonitor.Debugf(format, args...)
}

// Debugln write the debug msg
func (gl *GoLoggerWrapper) Debugln(args ...any) {
	gl.logMonitor.Debug(args...)
}

// Trace write the trace msg
func (gl *GoLoggerWrapper) Trace(format string) {
	gl.logMonitor.Debug(format)
}

// Tracef write the trace msg
func (gl *GoLoggerWrapper) Tracef(format string, args ...any) {
	gl.logMonitor.Debugf(format, args...)
}

// Traceln write the trace msg
func (gl *GoLoggerWrapper) Traceln(args ...any) {
	gl.logMonitor.Debug(args...)
}

// Info write the info msg
func (gl *GoLoggerWrapper) Info(format string) {
	gl.logMonitor.Info(format)
}

// Infof write the info msg
func (gl *GoLoggerWrapper) Infof(format string, args ...any) {
	gl.logMonitor.Infof(format, args...)
}

// Infoln write the info msg
func (gl *GoLoggerWrapper) Infoln(args ...any) {
	gl.logMonitor.Info(args...)
}

// Warn write the warn msg
func (gl *GoLoggerWrapper) Warn(format string) {
	gl.logMonitor.Warn(format)
}

// Warnf write the warn msg
func (gl *GoLoggerWrapper) Warnf(format string, args ...any) {
	gl.logMonitor.Warnf(format, args...)
}

// Warnln write the warn msg
func (gl *GoLoggerWrapper) Warnln(args ...any) {
	gl.logMonitor.Warn(args...)
}

// Error write the error msg
func (gl *GoLoggerWrapper) Error(format string) {
	gl.logMonitor.Error(format)
}

// Errorf write the error msg
func (gl *GoLoggerWrapper) Errorf(format string, args ...any) {
	gl.logMonitor.Errorf(format, args...)
}

// Errorln write the error msg
func (gl *GoLoggerWrapper) Errorln(args ...any) {
	gl.logMonitor.Error(args...)
}

// Panic write the panic msg
func (gl *GoLoggerWrapper) Panic(format string) {
	gl.logMonitor.Fatal(format)
}

// Panicf write the panic msg
func (gl *GoLoggerWrapper) Panicf(format string, args ...any) {
	gl.logMonitor.Fatalf(format, args...)
}

// Panicln write the panic msg
func (gl *GoLoggerWrapper) Panicln(args ...any) {
	gl.logMonitor.Fatal(args...)
}
