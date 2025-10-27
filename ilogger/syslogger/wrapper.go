package syslogger

import "github.com/ansencumt/ilog/v2/logger"

func NewILogger(loggerLevel logger.LoggerLevel) logger.Logger {
	return &defaultLogger{
		loggerLevel: loggerLevel,
	}
}
