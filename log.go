package xvlog

import (
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"github.com/tizx/xvlog/logger"
)

type log struct {
	logger logger.Logger
	config *config.Config
}


func NewLogger(c *config.Config) *log {
	if c == nil {
		c = config.DefaultConfig
	}
	log := &log{
		config: c,
	}

	if log.config.Logger == 0 {
		log.logger = logger.NewConsole()
	}else{
		log.logger = logger.NewBinLog()
	}
	go log.logger.Write() // 开启写goroutine
	return log
}

func (l *log)Debug(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.DEBUG, field, h)
	l.logger.Receive(log)
}
func (l *log)Info(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.INFO, field, h)

	l.logger.Receive(log)
}
func (l *log)Warn(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.WARN, field, h)

	l.logger.Receive(log)
}
func (l *log)Error(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.ERROR, field, h)

	l.logger.Receive(log)
}
func (l *log)tyFatal(field string, h map[string]interface{}) {
	log := logdata.NewLog(logdata.FATAL, field, h)

	l.logger.Receive(log)
}
