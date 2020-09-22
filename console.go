package fbl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tizx/fbl/config"
	"github.com/tizx/fbl/logdata"
	"os"
	"sync"
)

type console struct {
	logBuffer  *bufio.Writer
	bufferLock sync.Mutex
}

func NewConsole() Processor {
	return &console{
		logBuffer: bufio.NewWriterSize(os.Stdout, config.DefaultConfig.BufferSize),
	}
}

func (c *console) Process(log *logdata.Log) {
	field, err := json.Marshal(log.Fields)
	if err != nil {
		field = []byte("marshal error " + err.Error())
	}
	logByte := fmt.Sprintf("%s [%s] %s %s %s %d\n", log.Time.Format("2006-01-02 15:04:05"), logdata.LevelName[log.Level], log.Message, field, log.File, log.Line)

	c.bufferLock.Lock()
	_, _ = c.logBuffer.Write([]byte(logByte))
	c.bufferLock.Unlock()

}
func (c *console)SyncAndClose()  {

}