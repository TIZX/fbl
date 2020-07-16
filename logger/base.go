package logger

import (
	"fmt"
	"xvlog/logdata"
)

type baseLog struct {
	logChan chan *logdata.Log
	exit    chan bool
	handler Handler
}

func NewLog(handler Handler) Logger {
	return &baseLog{
		logChan: make(chan *logdata.Log, 10),
		exit:    make(chan bool),
		handler: handler,
	}
}

func (b *baseLog) Receive(log *logdata.Log) {
	select {
		case b.logChan <- log:
		default:
			fmt.Println("丢弃")
	}
}

func (b *baseLog) Exit() {
	b.exit <- true
}

func (b *baseLog) Write() {
	for {
		select {
		case log := <- b.logChan:
			b.handler.handle(log)
		case <-b.exit:
			return
		}
	}
}
