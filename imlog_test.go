package ilog

import (
	"testing"
)

func TestDefaultLog(t *testing.T) {
	Debug("this is defautl debug msg")
	Trace("this is defautl trace msg")
	Info("this is defautl info msg")
	Warn("this is defautl warn msg")
	Error("this is defautl error msg")
	// Panic("this is defautl panic msg")

	SetLevel(INFO)
	t.Log("set level")

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
	Panicln(format, "panic f")

	SetLevel(INFO)
}
