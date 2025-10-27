package logger

type LoggerLevel int

const (
	// TRACE 1
	TRACE LoggerLevel = iota + 1
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

// Name get the level name
func (l LoggerLevel) Name() LoggerLevelName {
	switch l {
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

// Name get the level name
func (l LoggerLevel) String() string {
	return string(l.Name())
}

// Int get the level int
func (l LoggerLevel) Int() int {
	return int(l)
}

// FromInt get the level from int
func LevelFromInt(level int) LoggerLevel {
	switch level {
	case 1:
		return TRACE
	case 2:
		return DEBUG
	case 3:
		return INFO
	case 4:
		return WARN
	case 5:
		return ERROR
	case 6:
		return PANIC
	}
	return 0
}

func LevelFromString(name string) LoggerLevel {
	switch LoggerLevelName(name) {
	case TraceStr:
		return TRACE
	case DebugStr:
		return DEBUG
	case InfoStr:
		return INFO
	case WarnStr:
		return WARN
	case ErrorStr:
		return ERROR
	case PanicStr:
		return PANIC
	}
	return 0
}

type LoggerLevelName string

const (
	// TRACE 1
	TraceStr LoggerLevelName = "trace"
	// DEBUG 2
	DebugStr LoggerLevelName = "debug"
	// INFO 3
	InfoStr LoggerLevelName = "info"
	// WARN 4
	WarnStr LoggerLevelName = "warn"
	// ERROR 5
	ErrorStr LoggerLevelName = "error"
	// PANIC 6
	PanicStr LoggerLevelName = "painc"
)

// Level get the level from name
func (n LoggerLevelName) Level() LoggerLevel {
	switch n {
	case TraceStr:
		return TRACE
	case DebugStr:
		return DEBUG
	case InfoStr:
		return INFO
	case WarnStr:
		return WARN
	case ErrorStr:
		return ERROR
	case PanicStr:
		return PANIC
	}
	return 0
}

// Logger the interface use in im
// 5 level: debug, trace, info, warn, error, panic
type Logger interface {
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
	SetLevel(level LoggerLevel)
	GetLevel() LoggerLevel
}
