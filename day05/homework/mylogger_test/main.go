package main

import (
	"github.com/jiangshao666/day05/homework/mylogger"
)

var logger mylogger.Logger

func main() {
	logger = mylogger.NewConsoleLogger("debug")
	logger = mylogger.NewFileLogger("debug", "./", "debug.log", 10*1024*1024)
	for {
		logger.Debug("这是一个debug日志")
		logger.Info("这是一个Info日志")
		logger.Warn("这是一个Warn日志")
		logger.Error("这是一个Error日志, err:%s 错误级别%d", "未知的错误", 1)
		//time.Sleep(2 * time.Second)
	}
}
