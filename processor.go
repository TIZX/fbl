package xvlog

import "github.com/tizx/xvlog/logdata"

type Processor interface {
	Process(log *logdata.Log)
	SyncAndClose()
}