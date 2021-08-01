package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	logLevel    LogLevel
	filePath    string
	fileName    string
	logFile     *os.File
	errFile     *os.File
	maxFileSize uint32
}

func NewFileLogger(logLevel string, filePath string, fileName string, size uint32) *FileLogger {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		panic(err)
	}
	fullPath := path.Join(filePath, fileName)
	fileObj, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}

	fullPath = fullPath + ".err"
	fileErrObj, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		logLevel:    level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: size,
		logFile:     fileObj,
		errFile:     fileErrObj,
	}
}

func (fl *FileLogger) checkFileSize(fileObj *os.File) bool {
	stat, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("checkFileSize failed, err:%s", err)
		return false
	}
	return stat.Size() >= int64(fl.maxFileSize)
}

func (fl *FileLogger) spitFile(fileObj *os.File) *os.File {
	curName := fileObj.Name()
	curPath := path.Join(fl.filePath, curName)
	now := time.Now()
	backFilePath := curPath + "_" + now.Format("20060102030405")

	fileObj.Close()
	os.Rename(curPath, backFilePath)

	fileObj, err := os.OpenFile(curPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	return fileObj
}

func (fl *FileLogger) log(level LogLevel, format string, a ...interface{}) {
	if level >= fl.logLevel {
		now := time.Now()
		sname := getLogLevelName(level)
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getCallInfo(3)
		if fl.checkFileSize(fl.logFile) {
			fl.logFile = fl.spitFile(fl.logFile)
		}
		fmt.Fprintf(fl.logFile, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:03:04"),
			sname, fileName, funcName, lineNo, msg)
		if level >= ERROR {
			if fl.checkFileSize(fl.errFile) {
				fl.errFile = fl.spitFile(fl.errFile)
			}
			fmt.Fprintf(fl.errFile, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:03:04"),
				sname, fileName, funcName, lineNo, msg)
		}
	}
}

func (fl *FileLogger) Debug(msg string, a ...interface{}) {
	fl.log(DEBUG, msg, a...)
}

func (fl *FileLogger) Info(msg string, a ...interface{}) {
	fl.log(INFO, msg, a...)
}

func (fl *FileLogger) Warn(msg string, a ...interface{}) {
	fl.log(WARN, msg, a...)
}

func (fl *FileLogger) Error(msg string, a ...interface{}) {
	fl.log(ERROR, msg, a...)
}
