package mylogger

import (
	"runtime"
	"path"
)
// 存放一些公用的工具函数

func getCallerInfo(skip int)(fileName string, line int, funcName string){
	pc,fileName,line,ok := runtime.Caller(skip)
	if !ok{
		return 
	}
	// 从 fileName 中剥离出文件名
	fileName = path.Base(fileName)
	// 根据 pc 获取函数名
	funcName = runtime.FuncForPC(pc).Name()
	return fileName, line, funcName
}