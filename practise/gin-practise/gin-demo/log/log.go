package log

import (
	"log"
	"os"
)

const (
	afileLog = "access.log"
	fileLog  = "gin.log"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file := fileLog
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	Info = log.New(logFile, "[INFO]\t", log.LstdFlags)
	Warning = log.New(logFile, "[WARNING]\t", log.LstdFlags)
	Error = log.New(logFile, "[ERROR]\t", log.LstdFlags)
	return
}

func AccessLog() *os.File {
	file := afileLog
	logfile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return logfile
}
