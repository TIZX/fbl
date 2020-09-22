package fbl

import (
	"github.com/tizx/fbl/logdata"
)

type Builder struct {
	log    *logdata.Log
	logger *logger
}

func NewBuilder(logger *logger) *Builder {
	return &Builder{
		log:    logdata.NewLog(),
		logger: logger,
	}
}

func (b *Builder) WithFields(fields map[string]interface{}) *Builder {
	b.log.Fields = fields
	return b
}

func (b *Builder) WithLevel(level logdata.Level) *Builder {
	b.log.Level = level
	return b
}

func (b *Builder) WithMessage(message string) *Builder {
	b.log.Message = message
	return b
}

func (b *Builder) Info(message string) {
	b.WithLevel(logdata.INFO)
	b.WithMessage(message)
	b.logger.receive(b.log)
}

func (b *Builder) Debug(message string) {
	b.WithLevel(logdata.DEBUG)
	b.WithMessage(message)
	b.logger.receive(b.log)
}

func (b *Builder) Warn(message string) {
	b.WithLevel(logdata.WARN)
	b.WithMessage(message)
	b.logger.receive(b.log)
}

func (b *Builder) Error(message string) {
	b.WithLevel(logdata.ERROR)
	b.WithMessage(message)
	b.logger.receive(b.log)
}

func (b *Builder) Fatal(message string) {
	b.WithLevel(logdata.FATAL)
	b.WithMessage(message)
	b.logger.receive(b.log)
}

func (b *Builder) Send() {
	b.logger.receive(b.log)
}
