package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"os"
	"sync"
)


type console struct {
	logIndex uint64
	logBuffer *bufio.Writer
	bufferLock sync.Mutex
	processorChan []chan *logdata.Log
}

func NewConsole() Logger {
	c := &console{
		logBuffer: bufio.NewWriterSize(os.Stdout, config.DefaultConfig.BufferSize),
	}
	processorChan := make([]chan *logdata.Log, config.DefaultConfig.Processor)

	for i:=0; i<config.DefaultConfig.Processor; i++ {
		processorChan[i] = make(chan *logdata.Log)
	}
	c.processorChan = processorChan
	return c
}

func (c *console) Receive(log *logdata.Log) {
	processorIndex := c.logIndex % uint64(config.DefaultConfig.Processor)
	c.processorChan[processorIndex] <- log
	c.logIndex++
}


func (c *console) Write() {
	// WaitGroup 防止创建goroutine后i还没复制给index就进行了i++
	var wait sync.WaitGroup
	for i:=0; i < config.DefaultConfig.Processor; i++ {
		wait.Add(1)
		go func() {
			index := i
			wait.Done()
			for {
				log := <- c.processorChan[index]
				logByte := c.processor(log)
				c.bufferLock.Lock()
				c.logBuffer.Write(logByte)
				c.bufferLock.Unlock()
			}
		}()
		wait.Wait()
	}

}
func (c *console)processor(log *logdata.Log) []byte {
	data, err := json.Marshal(log.Data())
	if err != nil {
		data = []byte("marshal error " + err.Error())
	}
	str := fmt.Sprintf("%s [%s] %s %s %s %d\n", log.Time().Format("2006-01-02 15:04:05"), logdata.LevelName[log.Level()], log.Field(), data, log.File(), log.Line())
	return []byte(str)
}
