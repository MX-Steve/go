package mylog

// 自定义一个日志库，实现日志记录的功能

// 日志分级别
// DEBUG TRACE INFO WARN ERROR CIRTAL

const (
	DEBUG = iota // 只有大于这个级别的日志才记录
	TRACE
	INFO
	WARN
	ERROR
	CRATAL
)

func getLevelStr(level int) string {
	switch level{
	case 0:
		return "DEBUG"
	case 1:
		return "TRACE"
	case 2:
		return "INFO"
	case 3:
		return "ERROR"
	case 4:
		return "WARN"
	case 5:
		return "CRATAL"
	default:
		return "DEBUG"
	}
}
