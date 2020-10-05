package file

import (
	"encoding/binary"
	"fmt"
	"github.com/tizx/fbl/binlog/general"
	"github.com/tizx/fbl/config"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

type logFile struct {
	header             *header
	file               *os.File
	size               int64 // 文件大小，所占字节数
	general struct{
		generalByte []byte
		generalID uint32
		generalIDLock sync.Mutex
		generalCache 	*general.Cache
	}
	IsCloseSuccessChan chan bool //退出通道
}



func NewLogFile() (*logFile, error) {
	var err error
	file := &logFile{
		header: &header{
			dataOffset:    16,
			generalOffset: 16,
		},
		IsCloseSuccessChan: make(chan bool, 1),
	}
	file.general.generalCache = general.NewCache()

	file.file, err = file.openFile()
	if err != nil {
		return nil, err
	}

	err = file.syncHeader()
	if err != nil {
		return nil, err
	}

	err = file.syncGeneralData()

	if err != nil {
		return nil, err
	}
	return file, nil
}


func (l *logFile) GetGeneralID(general *general.General) uint32 {
	if id, ok := l.general.generalCache.Get(general); ok {
		return id
	}
	l.general.generalIDLock.Lock()
	defer l.general.generalIDLock.Unlock()
	if id, ok := l.general.generalCache.Get(general); ok {
		return id
	}
	l.general.generalID++
	l.general.generalCache.Set(general, l.general.generalID)
	l.general.generalByte = append(l.general.generalByte, general.Encode(l.general.generalID)...)
	return l.general.generalID
}

func (l *logFile) Write(data []byte) (int, error) {
	return l.file.Write(data)
}

// 关闭写goroutine的对应chan后退出循环后调用
func (l *logFile) SyncAndClose() {
	// 同步头
	ret, err := l.file.Seek(0, io.SeekCurrent)
	if err != nil {
		fmt.Println("get file ret error")
	}
	l.header.generalOffset = uint64(ret)
	headerByte := l.header.ToByte()
	l.file.WriteAt(headerByte, 0)

	// 写入general数据
	_, _ = l.file.Write(l.general.generalByte)

	fmt.Println("all data write success")
	l.IsCloseSuccessChan <- true
}
// 根据配置打开对应的日志文件
func (l *logFile)openFile() (*os.File, error) {
	basePath := config.DefaultConfig.LocalFileConfig.LogPath

	fileName := time.Now().Format("2006-01-02") + ".log"
	_ = os.MkdirAll(basePath, os.ModeDir)
	return os.OpenFile(path.Join(basePath, fileName), os.O_RDWR|os.O_CREATE, os.ModePerm)
}
// 同步文件头
func (l *logFile)syncHeader() error {
	headerByte := make([]byte, headerLen)
	_, err := l.file.ReadAt(headerByte, 0)
	if err != nil && err != io.EOF {
		return err
	}
	l.header = newHeader(headerByte)
	_, err = l.file.WriteAt(l.header.ToByte(), 0)
	if err != nil {
		return err
	}
	return nil
}
// 同步General数据
func (l *logFile)syncGeneralData() error {
	l.file.Seek(int64(l.header.generalOffset), io.SeekStart)
	defer l.file.Seek(int64(l.header.generalOffset), io.SeekStart)

	var data = make([]byte, 0)
	var temp = make([]byte, 1024)
	for {
		n, err := l.file.Read(temp)
		if err == nil {
			data = append(data, temp[: n]...)
			continue
		}
		if err == io.EOF {
			break
		}
		return err
	}
	l.general.generalByte = data
	// 解析general byte数据
	var start, end int
	for start < len(data) {
		length := binary.BigEndian.Uint32(data[start:start+4])
		generalItem := &general.General{}
		end = start + int(length)
		id := generalItem.Decode(data[start:end])
		l.general.generalCache.Set(generalItem, id)
		l.general.generalID++
		start = end
	}
	return nil
}