package xvlog

import (
	"fmt"
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"sync"
)

type logger struct {
	logIndex      uint64
	processorChan []chan *logdata.Log // 每一个processor的chan
	processing sync.WaitGroup // 处理中的
	processor Processor		// 处理器
	config *config.Config
}

func NewLogger(c *config.Config) *logger {
	if c == nil {
		c = config.DefaultConfig
	}
	logger := &logger{
		config: c,
		processorChan: func()[]chan *logdata.Log {
			var processorChan = make([]chan *logdata.Log, c.ProcessorNumber)
			for i:=0;i<c.ProcessorNumber;i++{
				processorChan[i] = make(chan *logdata.Log)
			}
			return processorChan
		}(),
	}
	if c.Processor == 0 {
		logger.processor = NewConsole()
	} else {
		logger.processor = NewBinLog()
	}
	logger.write() // 开启写goroutine

	return logger
}

func (l *logger) receive(log *logdata.Log) {
	processorIndex := l.logIndex % uint64(l.config.ProcessorNumber)
	l.processorChan[processorIndex] <- log
	l.logIndex++
}


func (l *logger) write() {
	l.processing.Add(l.config.ProcessorNumber)
	for i := 0; i < l.config.ProcessorNumber; i++ {
		go func(index int) {
			for {
				log, ok := <- l.processorChan[index]
				if !ok {
					break // 退出循环
				}
				if log == nil {
					fmt.Println("log is nil ")
					continue
				}
				l.processor.Process(log)

			}
			l.processing.Done() // 减少正在运行的处理器
		}(i)
	}
}



func (l *logger) WithFields(fields map[string]interface{}) *Builder {
	return NewBuilder(l).WithFields(fields)
}

func (l *logger) WithLevel(level logdata.Level) *Builder {
	return NewBuilder(l).WithLevel(level)
}

func (l *logger) WithMessage(message string) *Builder {
	return NewBuilder(l).WithMessage(message)
}

func (l *logger) Info(message string) {
	NewBuilder(l).Info(message)
}

func (l *logger) Debug(message string) {
	NewBuilder(l).Debug(message)
}

func (l *logger) Warn(message string) {
	NewBuilder(l).Warn(message)
}

func (l *logger) Error(message string) {
	NewBuilder(l).Error(message)
}

func (l *logger) Fatal(message string) {
	NewBuilder(l).Fatal(message)
}

