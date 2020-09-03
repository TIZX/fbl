package logger

import (
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"github.com/tizx/xvlog/logger/binlog"
	"sync"
)

type binLog struct {
	logIndex uint64
	processorChan []chan *logdata.Log
}

func NewBinLog() *binLog {
	b := &binLog{}
	processorChan := make([]chan *logdata.Log, config.DefaultConfig.Processor)

	for i:=0; i<config.DefaultConfig.Processor; i++ {
		processorChan[i] = make(chan *logdata.Log)
	}
	b.processorChan = processorChan
	return b
}

func (b *binLog) Receive(log *logdata.Log) {
	processorIndex := b.logIndex % uint64(config.DefaultConfig.Processor)
	b.processorChan[processorIndex] <- log
	b.logIndex++
}

func (b *binLog) Write() {
	// WaitGroup 防止创建goroutine后i还没复制给index就进行了i++
	var wait sync.WaitGroup
	for i:=0; i < config.DefaultConfig.Processor; i++ {
		wait.Add(1)
		go func() {
			index := i
			wait.Done()
			for {
				log := <- b.processorChan[index]
				structure := binlog.NewStructure(log)
				structure.Parse()
			}
		}()
		wait.Wait()
	}
}

