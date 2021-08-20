package mylog 

import (
	"runtime"
	"path"
)

func getCallerInfo()(fileName, funcName string, line int){
	pc, fileName,line,ok := runtime.Caller(2) // int 是往上级找多少层
	// 根据 pc 拿到当前执行的函数名
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	fileName = path.Base(fileName)
	return fileName, funcName, line
}

