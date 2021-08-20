package mylog

import (
	"fmt"
	"os"
	"time"
)

// FileLogger 文件中记录日志的结构体
type FileLogger struct {
	level 			int 
	logFilePath		string
	logFileName		string 
	logFile			*os.File // os包中File指针类型
}

// NewFileLogger 是一个生成文件日志结构体实例的构造函数
func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	flobj := &FileLogger{
		level:level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flobj.initFileLogger()
	return flobj
}

// 专门用来初始化文件句柄
func (f *FileLogger)initFileLogger(){
	filepath := fmt.Sprintf("%s/%s",f.logFilePath,f.logFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed", filepath))
	}
	// 日志切割
	fName := fmt.Sprintf("%s/%s",f.logFilePath,f.logFileName)
	cutLog(fName)
	f.logFile = file
}

// Debug 记录DEBUG级别日志
func (f *FileLogger)Debug(format string , args ...interface{}){
	if f.level > DEBUG { // 如果你设置的日志级别大于当前级别，则不记录
		return
	}
	fileName, funcName, line := getCallerInfo()
	// 往文件中写
	// 日志格式：时间 日志级别 哪个文件 哪一行 哪一个函数 日志信息
	// f.logFile.WriteString(msg)
	// [2019-04-01 10:00:00] [DEBUG] main.go [14] id 为10 的用户一直在尝试破解银行
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s",nowStr, getLevelStr(f.level),fileName,funcName,line, format)
	fmt.Fprintf(f.logFile, format ,  args...)
	fmt.Fprintln(f.logFile) // 添加换号
}

// Info 记录 INFO 级别日志
func (f *FileLogger)Info(format string , args ...interface{}){
	if f.level > INFO {
		return
	}
	fileName, funcName, line := getCallerInfo()
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s",nowStr,getLevelStr(f.level),fileName,funcName,line, format)
	fmt.Fprintf(f.logFile, format , args...)
	fmt.Fprintln(f.logFile)
}

// Error 记录 ERROR 级别日志
func (f *FileLogger)Error(format string , args ...interface{}){
	if f.level > ERROR {
		return 
	}
	fileName, funcName, line := getCallerInfo()
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%s] %s",nowStr,getLevelStr(f.level),fileName,funcName,line,format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile)
}

// Warn 记录 WARN 级别日志
func (f *FileLogger)Warn(format string , args ...interface{}){
	if f.level > WARN {
		return 
	}
	fileName, funcName, line := getCallerInfo()
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%s] %s",nowStr,getLevelStr(f.level),fileName,funcName,line,format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile)
}

// Cratal 记录 CRATAL 级别日志
func (f *FileLogger)Cratal(format string , args ...interface{}){
	if f.level > CRATAL {
		return 
	}
	fileName, funcName, line := getCallerInfo()
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s",nowStr,getLevelStr(f.level),fileName,funcName,line,format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile)
}

