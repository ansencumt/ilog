package ilog

import (
	"fmt"

	"github.com/ansencumt/ilog/v2/ilogger/syslogger"
	"github.com/ansencumt/ilog/v2/logger"
)

var ilogger logger.Logger = syslogger.NewILogger(logger.INFO)

// SetLogger set the Log impl
// l the Log impl. must not be nil
func SetLogger(l logger.Logger) {
	if l != nil {
		ilogger = l
		return
	}
}

// GetLogger get the Log impl
func GetLogger() logger.Logger {
	return ilogger
}

// GetLevel 获取水平
func GetLevel() logger.LoggerLevel {
	return ilogger.GetLevel()
}

// GetLevelString 获取水平字符串
func GetLevelString() string {
	return ilogger.GetLevel().String()
}

// SetLevel 设置水平
func SetLevel(level logger.LoggerLevel) error {
	if level < logger.TRACE || level > logger.PANIC {
		ilogger.Errorf("level out of index: %d", level)
		return fmt.Errorf("level out of index: 1 to 6, but %d", level)
	}
	ilogger.SetLevel(level)
	return nil
}

// SetLevel 设置水平
func SetLevelWithString(levelStr string) error {
	ilogger.SetLevel(logger.LevelFromString(levelStr))
	return nil
}

// Debug write the debug msg
func Debug(format string) {
	ilogger.Debug(format)
}

// Debugf write the debug msg
func Debugf(format string, args ...any) {
	ilogger.Debugf(format, args...)
}

// Debugln write the debug msg
func Debugln(args ...any) {
	ilogger.Debugln(args...)
}

// Trace write the trace msg
func Trace(format string) {
	ilogger.Trace(format)
}

// Tracef write the trace msg
func Tracef(format string, args ...any) {
	ilogger.Tracef(format, args...)
}

// Traceln write the trace msg
func Traceln(args ...any) {
	ilogger.Traceln(args...)
}

// Info write the info msg
func Info(format string) {
	ilogger.Info(format)
}

// Infof write the info msg
func Infof(format string, args ...any) {
	ilogger.Infof(format, args...)
}

// Infoln write the info msg
func Infoln(args ...any) {
	ilogger.Infoln(args...)
}

// Warn write the warn msg
func Warn(format string) {
	ilogger.Warn(format)
}

// Warnf write the warn msg
func Warnf(format string, args ...any) {
	ilogger.Warnf(format, args...)
}

// Warnf write the warn msg
func Warnln(args ...any) {
	ilogger.Warnln(args...)
}

// Error write the error msg
func Error(format string) {
	ilogger.Error(format)
}

// Errorf write the error msg
func Errorf(format string, args ...any) {
	ilogger.Errorf(format, args...)
}

// Errorln write the error msg
func Errorln(args ...any) {
	ilogger.Errorln(args...)
}

// Panic write the panic msg
func Panic(format string) {
	ilogger.Panic(format)
}

// Panicf write the panic msg
func Panicf(format string, args ...any) {
	ilogger.Panicf(format, args...)
}

// Panicln write the panic msg
func Panicln(args ...any) {
	ilogger.Panicln(args...)
}

var djsoner logger.IJson = &logger.DefaultJson{}

func JSON(obj any) string {
	return djsoner.JSON(obj)
}
