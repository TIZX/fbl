package binlog

import (
	"github.com/tizx/fbl/binlog/general"
	"io"
)

type receiver interface {
	GetGeneralID(*general.General) uint32
	io.Writer
	SyncAndClose()
}

