package ilog

import (
	"fmt"
)

var ilogger = Log(new(defaultLogger))

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
	// SetLevel(level int)
}

// SetLogger set the Log impl
// l the Log impl. must not be nil
func SetLogger(l Log) {
	if l != nil {
		ilogger = l
		return
	}
}

var loggerLevel = INFO

func GetLevel() int {
	return loggerLevel
}

func GetLevelString() string {
	switch loggerLevel {
	case TRACE:
		return TraceStr
	case DEBUG:
		return DebugStr
	case INFO:
		return InfoStr
	case WARN:
		return WarnStr
	case ERROR:
		return ErrorStr
	case PANIC:
		return PanicStr
	}
	return "unknow"
}

// SetLevel 设置水平
func SetLevel(level int) error {
	if level < TRACE || level > PANIC {
		ilogger.Errorf("level out of index: %d", level)
		return fmt.Errorf("level out of index: 1 to 6, but %d", level)
	}
	loggerLevel = level
	// logger.SetLevel(level)
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
		ilogger.Errorf("level out of index: %d", level)
		return fmt.Errorf("level string is trace, debug, info, warn, panic, but %s", levelStr)
	}
	loggerLevel = level
	// logger.SetLevel(level)
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
	ilogger.Debug(format)
}

// Debugf write the debug msg
func Debugf(format string, args ...any) {
	if !isInLoggerLevel(DEBUG) {
		return
	}
	ilogger.Debugf(format, args...)
}

// Debugln write the debug msg
func Debugln(args ...any) {
	if !isInLoggerLevel(DEBUG) {
		return
	}
	ilogger.Debugln(args...)
}

// Trace write the trace msg
func Trace(format string) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	ilogger.Trace(format)
}

// Tracef write the trace msg
func Tracef(format string, args ...any) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	ilogger.Tracef(format, args...)
}

// Traceln write the trace msg
func Traceln(args ...any) {
	if !isInLoggerLevel(TRACE) {
		return
	}
	ilogger.Traceln(args...)
}

// Info write the info msg
func Info(format string) {
	if !isInLoggerLevel(INFO) {
		return
	}
	ilogger.Info(format)
}

// Infof write the info msg
func Infof(format string, args ...any) {
	if !isInLoggerLevel(INFO) {
		return
	}
	ilogger.Infof(format, args...)
}

// Infoln write the info msg
func Infoln(args ...any) {
	if !isInLoggerLevel(INFO) {
		return
	}
	ilogger.Infoln(args...)
}

// Warn write the warn msg
func Warn(format string) {
	if !isInLoggerLevel(WARN) {
		return
	}
	ilogger.Warn(format)
}

// Warnf write the warn msg
func Warnf(format string, args ...any) {
	if !isInLoggerLevel(WARN) {
		return
	}
	ilogger.Warnf(format, args...)
}

// Warnf write the warn msg
func Warnln(args ...any) {
	if !isInLoggerLevel(WARN) {
		return
	}
	ilogger.Warnln(args...)
}

// Error write the error msg
func Error(format string) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	ilogger.Error(format)
}

// Errorf write the error msg
func Errorf(format string, args ...any) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	ilogger.Errorf(format, args...)
}

// Errorln write the error msg
func Errorln(args ...any) {
	if !isInLoggerLevel(ERROR) {
		return
	}
	ilogger.Errorln(args...)
}

// Panic write the panic msg
func Panic(format string) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	ilogger.Panic(format)
}

// Panicf write the panic msg
func Panicf(format string, args ...any) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	ilogger.Panicf(format, args...)
}

// Panicln write the panic msg
func Panicln(args ...any) {
	if !isInLoggerLevel(PANIC) {
		return
	}
	ilogger.Panicln(args...)
}
