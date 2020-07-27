package xvlog

import (
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"github.com/tizx/xvlog/logger"
)

var loggerValue logger.Logger

type H = map[string]interface{}


func Logger()  {
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

func Debug(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.DEBUG, field, h)
	loggerValue.Receive(log)
}
func Info(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.INFO, field, h)

	loggerValue.Receive(log)
}
func Warn(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.WARN, field, h)
	loggerValue.Receive(log)
}
func Error(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.ERROR, field, h)
	loggerValue.Receive(log)
}
func Fatal(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.FATAL, field, h)
	loggerValue.Receive(log)
}
