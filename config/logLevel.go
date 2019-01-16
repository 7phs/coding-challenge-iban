package config

import (
	"fmt"
	"strings"
)

const (
	LogLevelError   LogLevel = "error"
	LogLevelWarning LogLevel = "warning"
	LogLevelInfo    LogLevel = "info"
	LogLevelDebug   LogLevel = "debug"
	LogLevelUnknown LogLevel = "unknown"
)

type LogLevel string

func NewLogLevel(str string) LogLevel {
	switch strings.ToLower(str) {
	case "error":
		return LogLevelError
	case "warning":
		return LogLevelWarning
	case "debug":
		return LogLevelDebug
	case "info":
		return LogLevelInfo
	default:
		return LogLevelUnknown
	}
}

func (o LogLevel) String() string {
	switch o {
	case LogLevelError:
		return "error"
	case LogLevelWarning:
		return "warning"
	case LogLevelDebug:
		return "debug"
	case LogLevelInfo:
		return "info"
	case LogLevelUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("[INVALID: %s]", string(o))
	}
}
