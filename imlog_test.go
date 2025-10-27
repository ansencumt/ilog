package ilog

import (
	"testing"

	"github.com/ansencumt/ilog/v2/ilogger/paylogger"
	"github.com/ansencumt/ilog/v2/logger"
)

func TestDefaultLog(t *testing.T) {
	t.Log("set level to default")
	Debug("this is defautl debug msg")
	Trace("this is defautl trace msg")
	Info("this is defautl info msg")
	Warn("this is defautl warn msg")
	Error("this is defautl error msg")
	// Panic("this is defautl panic msg")

	SetLevel(logger.TRACE)
	t.Log("*** set level to trace ***")

	format := "this is defautl format %s msg"
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.DEBUG)
	t.Log("*** set level to debug ***")

	format = "this is defautl format %s msg"
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.INFO)
	t.Log("*** set level to info ***")
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.WARN)
	t.Log("*** set level to warn ***")
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.ERROR)
	t.Log("*** set level to error ***")
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")
}

func TestGoLogger(t *testing.T) {
	gl := paylogger.NewILogger(nil)
	SetLogger(gl)

	t.Log("set level to default")
	Debug("this is defautl debug msg")
	Trace("this is defautl trace msg")
	Info("this is defautl info msg")
	Warn("this is defautl warn msg")
	Error("this is defautl error msg")
	// Panic("this is defautl panic msg")

	SetLevel(logger.TRACE)
	t.Log("*** set level to trace ***")

	format := "this is defautl format %s msg"
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.DEBUG)
	t.Log("*** set level to debug ***")

	format = "this is defautl format %s msg"
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.INFO)
	t.Log("*** set level to info ***")
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.WARN)
	t.Log("*** set level to warn ***")
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")

	SetLevel(logger.ERROR)
	t.Log("*** set level to error ***")
	Debugf(format, "debug f")
	Tracef(format, "trace f")
	Infof(format, "info f")
	Warnf(format, "warn f")
	Errorf(format, "error f")
	// Panicf(format, "panic f")

	format = "this is defautl format msg from ln"
	Debugln(format, "debug f")
	Traceln(format, "trace f")
	Infoln(format, "info f")
	Warnln(format, "warn f")
	Errorln(format, "error f")
	// Panicln(format, "panic f")
}
