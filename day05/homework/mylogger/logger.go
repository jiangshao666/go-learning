package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Logger interface {
	Debug(msg string, a ...interface{})
	Info(msg string, a ...interface{})
	Warn(msg string, a ...interface{})
	Error(msg string, a ...interface{})
}

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	default:
		err := errors.New("无效的日子级别")
		return UNKNOWN, err
	}
}

func getLogLevelName(logLevel LogLevel) string {
	switch logLevel {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	}
	return "UNKOWN"
}

func getCallInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, line
}
