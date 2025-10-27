package syslogger

import (
	"fmt"
	"log"

	"github.com/ansencumt/ilog/v2/logger"
)

const (
	prefixDebug = " [debug] "
	prefixTrace = " [trace] "
	prefixInfo  = " [info] "
	prefixWarn  = " [warn] "
	prefixError = " [error] "
	prefixPanic = " [panic] "
)

type defaultLogger struct {
	loggerLevel logger.LoggerLevel
}

func (logger *defaultLogger) isNotInLoggerLevel(level logger.LoggerLevel) bool {
	if level < logger.loggerLevel {
		return false
	}
	return true
}

// Trace write the trace msg
func (dl *defaultLogger) Trace(format string) {
	if !dl.isNotInLoggerLevel(logger.TRACE) {
		return
	}
	log.Println(prefixTrace, format)
}

// Tracef write the trace msg
func (dl *defaultLogger) Tracef(format string, args ...any) {
	if !dl.isNotInLoggerLevel(logger.TRACE) {
		return
	}
	log.Printf("%s%s", prefixTrace, fmt.Sprintf(format, args...))
}

// Traceln write the trace msg
func (dl *defaultLogger) Traceln(args ...any) {
	if !dl.isNotInLoggerLevel(logger.TRACE) {
		return
	}
	args = append([]any{prefixTrace}, args...)
	log.Println(args...)
}

// Debug write the debug msg
func (dl *defaultLogger) Debug(format string) {
	if !dl.isNotInLoggerLevel(logger.DEBUG) {
		return
	}
	log.Println(prefixDebug, format)
}

// Debugf write the debug msg
func (dl *defaultLogger) Debugf(format string, args ...any) {
	if !dl.isNotInLoggerLevel(logger.DEBUG) {
		return
	}

	log.Printf("%s%s", prefixDebug, fmt.Sprintf(format, args...))
}

// Debugln write the debug msg
func (dl *defaultLogger) Debugln(args ...any) {
	if !dl.isNotInLoggerLevel(logger.DEBUG) {
		return
	}

	args = append([]any{prefixDebug}, args...)
	log.Println(args...)
}

// Info write the info msg
func (dl *defaultLogger) Info(format string) {
	if !dl.isNotInLoggerLevel(logger.INFO) {
		return
	}
	log.Println(prefixInfo, format)
}

// Infof write the info msg
func (dl *defaultLogger) Infof(format string, args ...any) {
	if !dl.isNotInLoggerLevel(logger.INFO) {
		return
	}
	log.Printf("%s%s", prefixInfo, fmt.Sprintf(format, args...))
}

// Infoln write the info msg
func (dl *defaultLogger) Infoln(args ...any) {
	if !dl.isNotInLoggerLevel(logger.INFO) {
		return
	}
	args = append([]any{prefixInfo}, args...)
	log.Println(args...)
}

// Warn write the warn msg
func (dl *defaultLogger) Warn(format string) {
	if !dl.isNotInLoggerLevel(logger.WARN) {
		return
	}
	log.Println(prefixWarn, format)
}

// Warnf write the warn msg
func (dl *defaultLogger) Warnf(format string, args ...any) {
	if !dl.isNotInLoggerLevel(logger.WARN) {
		return
	}
	log.Printf("%s%s", prefixWarn, fmt.Sprintf(format, args...))
}

// Warnln write the warn msg
func (dl *defaultLogger) Warnln(args ...any) {
	if !dl.isNotInLoggerLevel(logger.WARN) {
		return
	}
	args = append([]any{prefixWarn}, args...)
	log.Println(args...)
}

// Error write the error msg
func (dl *defaultLogger) Error(format string) {
	if !dl.isNotInLoggerLevel(logger.ERROR) {
		return
	}
	log.Println(prefixError, format)
}

// Errorf write the error msg
func (dl *defaultLogger) Errorf(format string, args ...any) {
	if !dl.isNotInLoggerLevel(logger.ERROR) {
		return
	}
	log.Printf("%s%s", prefixError, fmt.Sprintf(format, args...))
}

// Errorln write the error msg
func (dl *defaultLogger) Errorln(args ...any) {
	if !dl.isNotInLoggerLevel(logger.ERROR) {
		return
	}
	args = append([]any{prefixError}, args...)
	log.Println(args...)
}

// Panic write the panic msg
func (dl *defaultLogger) Panic(format string) {
	if !dl.isNotInLoggerLevel(logger.PANIC) {
		return
	}
	log.Fatal(prefixPanic, format)
}

// Panicf write the panic msg
func (dl *defaultLogger) Panicf(format string, args ...any) {
	if !dl.isNotInLoggerLevel(logger.PANIC) {
		return
	}
	log.Fatalf("%s%s", prefixPanic, fmt.Sprintf(format, args...))
}

// Panicln write the panic msg
func (dl *defaultLogger) Panicln(args ...any) {
	if !dl.isNotInLoggerLevel(logger.PANIC) {
		return
	}
	args = append([]any{prefixPanic}, args...)
	log.Fatalln(args...)
}

// SetLevel 配置logger的大小
func (dl *defaultLogger) SetLevel(level logger.LoggerLevel) {
	dl.loggerLevel = level
}

// GetLevel 获取logger的大小
func (dl *defaultLogger) GetLevel() logger.LoggerLevel {
	return dl.loggerLevel
}
