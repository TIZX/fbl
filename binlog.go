package xvlog

import (
	"github.com/tizx/xvlog/binlog"
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"sync"
)

// binLog Logger
type binLog struct {
	logIndex      uint64
	processorChan []chan *logdata.Log
}

func NewBinLog() *binLog {
	b := &binLog{}
	processorChan := make([]chan *logdata.Log, config.DefaultConfig.Processor)

	for i := 0; i < config.DefaultConfig.Processor; i++ {
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
	var close sync.Once
	for i := 0; i < config.DefaultConfig.Processor; i++ {
		wait.Add(1)
		go func() {
			index := i
			wait.Done()
			defer func() {
				close.Do(func() {
					binlog.Write.Flush()
					binlog.Write.Close()
				})
			}()
			for {
				log := <-b.processorChan[index]
				structure := binlog.NewStructure(log)
				structure.Parse()
			}
		}()
		wait.Wait()
	}
}
