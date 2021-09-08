package mylogger

import (
	"fmt"
	"os"
	"time"
)

// 往终端写日志
// ConsoleLogger 终端日志结构体
type ConsoleLogger struct {
	level 	Level
}

// NewConsoleLogger 终端结构体的构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger{
	logLevel := parseLogLevel(levelStr)
	cl := &ConsoleLogger{
		level: logLevel,
	}
	return cl
}

// 方法
// log 将公用的记录日志的工具封装成单独的方法
func (c *ConsoleLogger)log(level Level, format string, args ...interface{}){
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志
	// 日志格式 [时间] [文件:行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2016-01-02 15:04:05.000")
	fileName,line,funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",nowStr,fileName,line,funcName,logLevelStr, msg)
	fmt.Fprintln(os.Stdout,logMsg) 
}	

// Debug 方法
func (c *ConsoleLogger)Debug(format string, args ...interface{}){
	c.log(DebugLevel, format, args...)
}

// Info 方法
func (c *ConsoleLogger)Info(format string, args ...interface{}){
	c.log(InfoLevel, format, args...)
}

// Warning 方法
func (c *ConsoleLogger)Warn(format string, args ...interface{}){
	c.log(WarningLevel, format, args...)
}

// Error 方法
func (c *ConsoleLogger)Error(format string, args ...interface{}){
	c.log(ErrorLevel, format, args...)
}

// Fatal 方法
func (c *ConsoleLogger)Fatal(format string, args ...interface{}){
	c.log(FatalLevel, format, args...)
}

// Close 终端标准输出不需要关闭
func (c *ConsoleLogger)Close(){
	
}