package binlog

import (
	"bufio"
	"github.com/tizx/fbl/binlog/file"
	"github.com/tizx/fbl/binlog/general"
	"github.com/tizx/fbl/binlog/parse"
	"github.com/tizx/fbl/config"
	"github.com/tizx/fbl/logdata"
)

type Index uint32

type Handle struct {
	dataBuf *bufio.Writer
	receiver receiver
	dataWriteChan      chan []byte
	size     int64 // 文件大小
}

func NewHandle() (*Handle, error) {
	handle := &Handle{}
	var err error
	if config.DefaultConfig.BinlogReceiver == config.LocalFile {
		handle.receiver, err = file.NewLogFile()
		if err != nil {
			return nil, err
		}
		handle.dataBuf = bufio.NewWriter(handle.receiver)
	}
	handle.dataWriteChan = make(chan []byte, 0)
	go handle.startWrite()
	return handle, nil
}

// 解析日志
func (h *Handle) Parse(log *logdata.Log) {
	generalItem := &general.General{
		Level:   log.Level,
		Message: log.Message,
		File:    log.File,
		Line:    log.Line,
	}

	typeNameByte, dataTemp := parse.Encode(log.Fields)

	generalItem.TypeNameByte = string(typeNameByte)

	generalID := h.receiver.GetGeneralID(generalItem)

	data := &logData{}
	data.dataByte = dataTemp
	data.time = uint32(log.Time.Unix())
	data.generalID = uint32(generalID)

	dataByte := data.Encode()

	h.dataWriteChan <- dataByte // 写入写channel
}

func (h *Handle) SyncAndClose() {
	_ = h.dataBuf.Flush()
	h.receiver.SyncAndClose()
}

func (h *Handle)startWrite()  {
	for {
		data, ok := <- h.dataWriteChan
		if !ok {
			break
		}
		h.dataBuf.Write(data)
	}
}