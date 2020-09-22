package binlog

import (
	"fmt"
	"github.com/tizx/xvlog/binlog/parse"
	"github.com/tizx/xvlog/config"
	"github.com/tizx/xvlog/logdata"
	"os"
	"path"
	"time"
)

type Index uint32

type Handle struct {
	file *logFile
	size int64 // 文件大小
}

func NewHandle() (*Handle, error) {
	handle := &Handle{}
	filePath := config.DefaultConfig.LogPath
	_ = os.MkdirAll(filePath, os.ModeDir)
	fileName := time.Now().Format("2006-01-02")
	ext := ".xvlog"
	var err error
	handle.file, err = NewLogFile(path.Join(filePath, fileName+ext))
	if err != nil {
		return nil, err
	}
	go handle.file.startWrite()
	return handle, err
}

// 解析日志
func (h *Handle) Parse(log *logdata.Log) {
	general := &general{
		Level:   log.Level,
		Message: log.Message,
		File:    log.File,
		Line:    log.Line,
	}

	typeNameByte, dataTemp := parse.Encode(log.Fields)


	general.TypeNameByte = string(typeNameByte)

	generalID := h.file.generalMap.getGeneralID(general)

	data := &logData{}
	data.dataByte = dataTemp
	data.time = uint32(log.Time.Unix())
	data.generalID = uint32(generalID)

	dataByte := data.Encode()
	fmt.Println(data)
	h.file.dataWriteChan <- dataByte                             // 写入写channel
}

func (h *Handle)SyncAndClose()  {
	close(h.file.dataWriteChan) // 关闭data写chan
	<- h.file.IsCloseSuccessChan //等待全部数据写入成功
}
