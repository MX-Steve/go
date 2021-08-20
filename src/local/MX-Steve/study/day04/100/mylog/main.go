package mylog

/**
Custom Log Library to record log data.
support:
	1. write log to file
	2. write log to tty
*/

// Log Level: DEBUG,TRACE,INFO,WARNING,ERROR,CRATal
const (
	DEBUG = iota
	TRACE
	INFO
	WARN
	ERROR
	CIRTAL
)

const FileMaxSize = 1024 * 1024 * 20

func getLevelstr(level int) string {
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "INFO"
	case 2:
		return "WARN"
	case 3:
		return "ERROR"
	case 4:
		return "CIRTAL"
	default:
		return "DEBUG"
	}
}
