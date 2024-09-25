package ilog

import (
	"fmt"
	"log"
)

var (
	prefixDebug = " [debug] "
	prefixTrace = " [trace] "
	prefixInfo  = " [info] "
	prefixWarn  = " [warn] "
	prefixError = " [error] "
	prefixPanic = " [panic] "
)

type defaultLogger struct {
}

// Debug write the debug msg
func (logger *defaultLogger) Debug(format string) {
	log.Println(prefixDebug, format)
}

// Debugf write the debug msg
func (logger *defaultLogger) Debugf(format string, args ...any) {
	log.Printf("%s%s", prefixDebug, fmt.Sprintf(format, args...))
}

// Debugln write the debug msg
func (logger *defaultLogger) Debugln(args ...any) {
	args = append([]any{prefixDebug}, args...)
	log.Println(args...)
}

// Trace write the trace msg
func (logger *defaultLogger) Trace(format string) {
	log.Println(prefixTrace, format)
}

// Tracef write the trace msg
func (logger *defaultLogger) Tracef(format string, args ...any) {
	log.Printf("%s%s", prefixTrace, fmt.Sprintf(format, args...))
}

// Traceln write the trace msg
func (logger *defaultLogger) Traceln(args ...any) {
	args = append([]any{prefixTrace}, args...)
	log.Println(args...)
}

// Info write the info msg
func (logger *defaultLogger) Info(format string) {
	log.Println(prefixInfo, format)
}

// Infof write the info msg
func (logger *defaultLogger) Infof(format string, args ...any) {
	log.Printf("%s%s", prefixInfo, fmt.Sprintf(format, args...))
}

// Infoln write the info msg
func (logger *defaultLogger) Infoln(args ...any) {
	args = append([]any{prefixInfo}, args...)
	log.Println(args...)
}

// Warn write the warn msg
func (logger *defaultLogger) Warn(format string) {
	log.Println(prefixWarn, format)
}

// Warnf write the warn msg
func (logger *defaultLogger) Warnf(format string, args ...any) {
	log.Printf("%s%s", prefixWarn, fmt.Sprintf(format, args...))
}

// Warnln write the warn msg
func (logger *defaultLogger) Warnln(args ...any) {
	args = append([]any{prefixWarn}, args...)
	log.Println(args...)
}

// Error write the error msg
func (logger *defaultLogger) Error(format string) {
	log.Println(prefixError, format)
}

// Errorf write the error msg
func (logger *defaultLogger) Errorf(format string, args ...any) {
	log.Printf("%s%s", prefixError, fmt.Sprintf(format, args...))
}

// Errorln write the error msg
func (logger *defaultLogger) Errorln(args ...any) {
	args = append([]any{prefixError}, args...)
	log.Println(args...)
}

// Panic write the panic msg
func (logger *defaultLogger) Panic(format string) {
	log.Fatal(prefixPanic, format)
}

// Panicf write the panic msg
func (logger *defaultLogger) Panicf(format string, args ...any) {
	log.Fatalf("%s%s", prefixPanic, fmt.Sprintf(format, args...))
}

// Panicln write the panic msg
func (logger *defaultLogger) Panicln(args ...any) {
	args = append([]any{prefixPanic}, args...)
	log.Fatalln(args...)
}

// SetLevel 配置logger的大小
func (logger *defaultLogger) SetLevel(level int) {

}
