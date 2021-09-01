package mylog

import (
	"fmt"
	"io"
	"os"
	"time"
)

// FileLogger: write log data to file
type FileLogger struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File
}

// NewFileLogger: generate struct
func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	flObj := &FileLogger{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flObj.initFileLogger()
	return flObj
}

// initalize file
func (f *FileLogger) initFileLogger() {
	filepath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file: %s failed", filepath))
	}
	f.logFile = file
}

// Debug Level
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.logFormat(DEBUG, format, args...)
}

// Info Level
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.logFormat(INFO, format, args...)
}

// Warn Level
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.logFormat(WARN, format, args...)
}

// Error Level
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.logFormat(ERROR, format, args...)
}

// Cirtal Level
func (f *FileLogger) Cirtal(format string, args ...interface{}) {
	f.logFormat(CIRTAL, format, args...)
}

func (f *FileLogger) logFormat(level int, format string, args ...interface{}) {
	f.CutLogByTime()
	// f.CutLogBySize()
	if f.level > level {
		return
	}
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	funcName, file, line := getCallerInfo()
	format = fmt.Sprintf("%s [%s] [%s %s %d] %s", nowStr, getLevelstr(f.level), file, funcName, line, format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile)
}

func (f *FileLogger) CutLogBySize() {
	f1, _ := f.logFile.Stat()
	if f1.Size() > FileMaxSize {
		now := time.Now().Format("20060102150405")
		log := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
		log1, _ := os.OpenFile(log, os.O_RDONLY, 0644)
		newLog := fmt.Sprintf("%s/%s.%s", f.logFilePath, f.logFileName, now)
		newLog1, _ := os.OpenFile(newLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		io.Copy(newLog1, log1)
		os.OpenFile(log, os.O_TRUNC, 0644)
	}
}

func (f *FileLogger) CutLogByTime() {
	now := time.Now()
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	if now.Minute() == 22 && now.Second() == 0 {
		log := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
		log1, _ := os.OpenFile(log, os.O_RDONLY, 0644)
		newLog := fmt.Sprintf("%s/%s.%s", f.logFilePath, f.logFileName, now)
		newLog1, _ := os.OpenFile(newLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		io.Copy(newLog1, log1)
		os.OpenFile(log, os.O_TRUNC, 0644)
	}
}
