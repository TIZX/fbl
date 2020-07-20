package xvlog

import (
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"github.com/tizx/xvlog/logger"
)

var loggerValue logger.Logger

func Logger() {
	var handle logger.Handler
	switch config.Config.Logger {
		case 1:
			handle = logger.NewFile()
		default:
			handle = logger.NewConsole()
	}

	loggerValue = logger.NewLog(handle)

	go loggerValue.Write()
}

func Debug(format string, a ...interface{}) {
	log := logdata.NewLog(logdata.DEBUG, format, a)
	loggerValue.Receive(log)
}
func Info(format string, a ...interface{}) {
	log := logdata.NewLog(logdata.INFO, format, a)

	loggerValue.Receive(log)
}
func Warn(format string, a ...interface{}) {
	log := logdata.NewLog(logdata.WARN, format, a)
	loggerValue.Receive(log)
}
func Error(format string, a ...interface{}) {
	log := logdata.NewLog(logdata.ERROR, format, a)
	loggerValue.Receive(log)
}
func Fatal(format string, a ...interface{}) {
	log := logdata.NewLog(logdata.FATAL, format, a)
	loggerValue.Receive(log)
}
