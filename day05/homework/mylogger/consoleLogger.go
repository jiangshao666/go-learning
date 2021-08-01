package mylogger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	logLevel LogLevel
}

func NewConsoleLogger(logLevel string) *ConsoleLogger {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{
		logLevel: level,
	}
}

func (c *ConsoleLogger) log(level LogLevel, format string, a ...interface{}) {
	if level >= c.logLevel {
		now := time.Now()
		sname := getLogLevelName(level)
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getCallInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:03:04"),
			sname, fileName, funcName, lineNo, msg)
	}
}

func (c *ConsoleLogger) Debug(msg string, a ...interface{}) {
	c.log(DEBUG, msg, a...)
}

func (c *ConsoleLogger) Info(msg string, a ...interface{}) {
	c.log(INFO, msg, a...)
}

func (c *ConsoleLogger) Warn(msg string, a ...interface{}) {
	c.log(WARN, msg, a...)
}

func (c *ConsoleLogger) Error(msg string, a ...interface{}) {
	c.log(ERROR, msg, a...)
}
