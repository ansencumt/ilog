package ilog

import (
	"fmt"
)

var logger = Log(new(defaultLogger))

const (
	// TRACE 1
	TRACE = iota + 1
	// DEBUG 2
	DEBUG
	// INFO 3
	INFO
	// WARN 4
	WARN
	// ERROR 5
	ERROR
	// PANIC 6
	PANIC
)

const (
	// TRACE 1
	TraceStr = "trace"
	// DEBUG 2
	DebugStr = "debug"
	// INFO 3
	InfoStr = "info"
	// WARN 4
	WarnStr = "warn"
	// ERROR 5
	ErrorStr = "error"
	// PANIC 6
	PanicStr = "painc"
)

// Log the interface use in im
// 5 level: debug, trace, info, warn, error, panic
type Log interface {
	Debug(format string)
	Debugf(format string, args ...any)
	Debugln(args ...any)
	Trace(format string)
	Tracef(format string, args ...any)
	Traceln(args ...any)
	Info(format string)
	Infof(format string, args ...any)
	Infoln(args ...any)
	Warn(format string)
	Warnf(format string, args ...any)
	Warnln(args ...any)
	Error(format string)
	Errorf(format string, args ...any)
	Errorln(args ...any)
	Panic(format string)
	Panicf(format string, args ...any)
	Panicln(args ...any)
	SetLevel(level int)
}

// SetLogger set the Log impl
// l the Log impl. must not be nil
func SetLogger(l Log) {
	if l != nil {
		logger = l
		return
	}
}

var loggerLevel = INFO

// SetLevel 设置水平
func SetLevel(level int) error {
	if level < TRACE || level > PANIC {
		logger.Errorf("level out of index: %d", level)
		return fmt.Errorf("level out of index: 1 to 6, but %d", level)
	}
	loggerLevel = level
	logger.SetLevel(level)
	return nil
}

// SetLevel 设置水平
func SetLevelWithString(levelStr string) error {
	level := ERROR

	switch levelStr {
	case TraceStr:
		level = TRACE
		break
	case DebugStr:
		level = DEBUG
		break
	case InfoStr:
		level = INFO
		break
	case WarnStr:
		level = WARN
		break
	case PanicStr:
		level = PANIC
		break
	default:
		logger.Errorf("level out of index: %d", level)
		return fmt.Errorf("level string is trace, debug, info, warn, panic, but %s", levelStr)
	}
	loggerLevel = level
	logger.SetLevel(level)
	return nil
}

func isInLoggerLevel(level int) bool {
	if level < loggerLevel {
		return false
	}
	return true
}

// Debug write the debug msg
func Debug(format string) {
	if !isInLoggerLevel(DEBUG) {
		return
	}
	logger.Debug(format)
}

// Debugf write the debug msg
func Debugf(format string, args ...any) {
	if !isInLoggerLevel(DEBUG) {
		return
	}
	logger.Debugf(format, args...)
}

// Debugln write the debug msg
func Debugln(args ...any) {
	if !isInLoggerLevel(DEBUG) {
		return
	}
	logger.Debugln(args...)
}

// Trace write the trace msg
func Trace(format string) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	logger.Trace(format)
}

// Tracef write the trace msg
func Tracef(format string, args ...any) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	logger.Tracef(format, args...)
}

// Traceln write the trace msg
func Traceln(args ...any) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	logger.Traceln(args...)
}

// Info write the info msg
func Info(format string) {
	if !isInLoggerLevel(INFO) {
		return
	}
	logger.Info(format)
}

// Infof write the info msg
func Infof(format string, args ...any) {
	if !isInLoggerLevel(INFO) {
		return
	}
	logger.Infof(format, args...)
}

// Infoln write the info msg
func Infoln(args ...any) {
	if !isInLoggerLevel(INFO) {
		return
	}
	logger.Infoln(args...)
}

// Warn write the warn msg
func Warn(format string) {
	if !isInLoggerLevel(WARN) {
		return
	}
	logger.Warn(format)
}

// Warnf write the warn msg
func Warnf(format string, args ...any) {
	if !isInLoggerLevel(WARN) {
		return
	}
	logger.Warnf(format, args...)
}

// Warnf write the warn msg
func Warnln(args ...any) {
	if !isInLoggerLevel(WARN) {
		return
	}
	logger.Warnln(args...)
}

// Error write the error msg
func Error(format string) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	logger.Error(format)
}

// Errorf write the error msg
func Errorf(format string, args ...any) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	logger.Errorf(format, args...)
}

// Errorln write the error msg
func Errorln(args ...any) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	logger.Errorln(args...)
}

// Panic write the panic msg
func Panic(format string) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	logger.Panic(format)
}

// Panicf write the panic msg
func Panicf(format string, args ...any) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	logger.Panicf(format, args...)
}

// Panicln write the panic msg
func Panicln(args ...any) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	logger.Panicln(args...)
}
