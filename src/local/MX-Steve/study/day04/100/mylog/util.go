package mylog

import (
	"path"
	"runtime"
)

func getCallerInfo() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	file = path.Base(file)
	return funcName, file, line
}
