package fbl

import (
	"github.com/tizx/fbl/binlog"
	"github.com/tizx/fbl/logdata"
)

// binLog Logger
type binLog struct {
	handle *binlog.Handle
}

func NewBinLog() Processor {
	b := &binLog{}
	var err error
	b.handle, err = binlog.NewHandle()
	if err != nil {
		panic(err)
	}
	return b
}

func (b *binLog)Process(log *logdata.Log)  {
	b.handle.Parse(log)
}

func (b *binLog)SyncAndClose()  {
	b.handle.SyncAndClose()
}