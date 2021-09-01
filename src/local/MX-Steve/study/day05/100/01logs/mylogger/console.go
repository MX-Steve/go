package mylogger

import (
	"fmt"
	"os"
	"time"
)

// ConsoleLogger write data into console
type ConsoleLogger struct {
	level Level
}

// NewFileLogger: FileLogger Generation function
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &ConsoleLogger{
		level: logLevel,
	}
	return fl
}

// Debug: Write Debug log file
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

// Info: Write Info log file
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

// Warning: Write Warning log file
func (c *ConsoleLogger) Warning(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

// Error: Write Error log file
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

// Fatal: Write Fatal log file
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}
func (c *ConsoleLogger) Close() {

}

// write log data into file
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	// log format:[time][file:line][function][log level] log message
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, funcName, getLevelStr(level), msg)
	fmt.Fprintln(os.Stdout, logMsg)
}
