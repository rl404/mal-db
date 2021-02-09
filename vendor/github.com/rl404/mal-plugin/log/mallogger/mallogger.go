// Package mallogger provides simple log printing with customizable color
// and level.
package mallogger

import (
	"fmt"
	"log"
	"runtime"
)

const (
	// LevelZero is no log level.
	LevelZero = iota
	// LevelHigh is log level showing Error and Fatal.
	LevelHigh
	// LevelNormal is log level showing Info, Error, and Fatal.
	LevelNormal
	// LevelDebug is log level showing Debug, Info, Warn, Error, and Fatal.
	LevelDebug
	// LevelComplete is log level showing all type of log.
	LevelComplete
)

// Log types.
const (
	TypeTrace = iota
	TypeDebug
	TypeInfo
	TypeWarn
	TypeError
	TypeFatal
)

var icon = []string{"[T]", "[D]", "[I]", "[W]", "[E]", "[F]"}

// Foreground text colors. The output colors may vary on different OS.
// Taken from https://en.wikipedia.org/wiki/ANSI_escape_code.
const (
	Reset         = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	White         = "\033[37m"
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

var colors = []string{Blue, Magenta, Green, BrightYellow, BrightRed, Red}

// Log is log.
type Log struct {
	level    int
	useColor bool
}

// New to create new log.
func New(level int, useColor bool) Log {
	return Log{
		level:    level,
		useColor: useColor,
	}
}

// Print is a wrapper to print different type of log.
func (l Log) Print(logType int, format string, args ...interface{}) {
	switch l.level {
	case LevelZero:
		return
	case LevelHigh:
		if logType != TypeError && logType != TypeFatal {
			return
		}
	case LevelNormal:
		if logType != TypeInfo && logType != TypeError && logType != TypeFatal {
			return
		}
	case LevelDebug:
		if logType == TypeTrace {
			return
		}
	case LevelComplete:
	default:
		return
	}

	log.Printf(l.fmt(logType, format)+Reset, args...)
}

func (l Log) colorize(color string, str string) string {
	if !l.useColor {
		color = ""
	}
	return color + str + Reset
}

func (l Log) fmt(logType int, format string) string {
	return fmt.Sprintf(
		"%s %s",
		l.colorize(colors[logType], icon[logType]),
		l.colorize(colors[logType], format),
	)
}

// Trace to log trace.
func (l Log) Trace(format string, args ...interface{}) {
	l.Print(TypeTrace, format, args...)
}

// Debug to log debug.
func (l Log) Debug(format string, args ...interface{}) {
	l.Print(TypeDebug, format, args...)
}

// Info to log info.
func (l Log) Info(format string, args ...interface{}) {
	l.Print(TypeInfo, format, args...)
}

// Warn to log warning.
func (l Log) Warn(format string, args ...interface{}) {
	l.Print(TypeWarn, format, args...)
}

// Error to log error.
func (l Log) Error(format string, args ...interface{}) {
	l.Print(TypeError, format, args...)
}

// Fatal to log trace.
func (l Log) Fatal(format string, args ...interface{}) {
	l.Print(TypeFatal, format, args...)
	runtime.Goexit()
}
