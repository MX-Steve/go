package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// write log data into file

// FileLogger file log struct
type FileLogger struct {
	level    Level
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
	maxSize  int64
}

// NewFileLogger: FileLogger Generation function
func NewFileLogger(levelStr, filePath, fileName string) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &FileLogger{
		level:    logLevel,
		fileName: fileName,
		filePath: filePath,
		maxSize:  10 * 1024 * 1024,
	}
	fl.initFile()
	return fl
}

// initFile
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	// open normal log file
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open log file %s failed,err: %s", logName, err))
	}
	f.file = fileObj
	// open err log file
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open log file %s failed,err: %s", errLogName, err))
	}
	f.errFile = errFileObj
}

// Debug: Write Debug log file
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info: Write Info log file
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warning: Write Warning log file
func (f *FileLogger) Warning(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error: Write Error log file
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// Fatal: Write Fatal log file
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
}

// write log data into file
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	// log format:[time][file:line][function][log level] log message
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, funcName, getLevelStr(level), msg)
	if f.checkSplit(f.file) {
		f.file = f.SplitLogFile(f.file)
	}

	fmt.Fprintln(f.file, logMsg)
	// Err log into err log file
	if level >= ErrorLevel {
		if f.checkSplit(f.errFile) {
			f.errFile = f.SplitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, logMsg)
	}
}

// check log file size
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat()
	return fileInfo.Size() >= f.maxSize
}

// cut log method
func (f *FileLogger) SplitLogFile(file *os.File) *os.File {
	fileName := file.Name()
	backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Format("20060102150405"))
	file.Close()
	os.Rename(fileName, backupName)
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open log file err"))
	}
	return fileObj
}
