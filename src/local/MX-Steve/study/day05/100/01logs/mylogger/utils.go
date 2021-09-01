package mylogger

import (
	"path"
	"runtime"
	"strings"
)

// utils: storage some public functions

// getCallerInfo get line file function
func getCallerInfo(skip int) (fileName string, line int, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	fileName = path.Base(fileName)
	// get function name from pc
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return
}

// getLevelStr
func getLevelStr(level Level) string {
	switch level {
	case 0:
		return "Debug"
	case 1:
		return "Info"
	case 2:
		return "Warn"
	case 3:
		return "Error"
	case 4:
		return "Fatal"
	default:
		return "Debug"
	}
}

// get level from user's string
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
