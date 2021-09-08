package mylogger

import (
	"strings"
)
// my log database

// Level 是一个自定义的类型，代表日志级别
type Level uint16

// 定义具体日志级别常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// Logger 定义一个logger接口
type Logger interface{
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

// getLevelStr 通过 level 获取日志类型
func getLevelStr(level Level) string {
	switch level{
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case ErrorLevel:
		return "ERROR"
	case WarningLevel:
		return "WARN"
	case FatalLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// parseLogLevel 根据用户传入的字符串类型的日志级别，解析出对应的Level
func parseLogLevel(levelStr string) Level {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}