package fbl

import "github.com/tizx/fbl/logdata"

type Processor interface {
	Process(log *logdata.Log)
	SyncAndClose()
}
