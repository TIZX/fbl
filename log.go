package xvlog

import (
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
)

type log struct {
	logger Logger
	config config.Config
}

func NewLogger(c *config.Config) *log {
	if c == nil {
		c = config.DefaultConfig
	}
	logger := &log{
		config: config.Config{},
	}

	if c.Logger == 0 {
		logger.logger = NewConsole()
	} else {
		logger.logger = NewBinLog()
	}
	go logger.logger.Write() // 开启写goroutine
	return logger
}

func (l *log) WithFields(fields map[string]interface{}) *Builder {
	return NewBuilder(l.logger).WithFields(fields)
}

func (l *log) WithLevel(level logdata.Level) *Builder {
	return NewBuilder(l.logger).WithLevel(level)
}

func (l *log) WithMessage(message string) *Builder {
	return NewBuilder(l.logger).WithMessage(message)
}

func (l *log) Info(message string) {
	NewBuilder(l.logger).Info(message)
}

func (l *log) Debug(message string) {
	NewBuilder(l.logger).Debug(message)
}

func (l *log) Warn(message string) {
	NewBuilder(l.logger).Warn(message)
}

func (l *log) Error(message string) {
	NewBuilder(l.logger).Error(message)
}

func (l *log) Fatal(message string) {
	NewBuilder(l.logger).Fatal(message)
}
