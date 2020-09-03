package rw

import (
	"bufio"
	"github.com/tizx/xvlog/config"
	"os"
	"path"
)

// 写入文件的数据的偏移量
type offset struct {
	indexOffset int64
	dataOffset int64
	generalOffset int64
}

type buffer struct {
	indexBuf *bufio.Writer
	dataBuf *bufio.Writer
	generalBuf *bufio.Writer
}

type file struct {
	buffer *buffer
	offset *offset

	index *os.File
	data *os.File
	describe *os.File
	general *os.File
}

func getFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if os.IsNotExist(err) {
		file, err = os.Create(fileName)
		if err != nil {
			panic("create "+ fileName +" file error: " + err.Error())
		}
	}
	return file
}

func NewWrite(fileName string) *file {
	err := os.MkdirAll(fileName, 0755)
	if err != nil {
		panic("mkdir "+ fileName + "error: "+err.Error())
	}
	describeFileName := path.Join(fileName, "Describe")
	indexFileName := path.Join(fileName, "Index")
	dataFileName := path.Join(fileName, "Data")
	generalFileName := path.Join(fileName, "General")
	file := &file{
		index: getFile(indexFileName),
		data: getFile(dataFileName),
		describe: getFile(describeFileName),
		general: getFile(generalFileName),
	}

	file.buffer = &buffer{
		indexBuf:   bufio.NewWriterSize(file.index, config.DefaultConfig.BufferSize),
		dataBuf:    bufio.NewWriterSize(file.data, config.DefaultConfig.BufferSize),
		generalBuf: bufio.NewWriterSize(file.general,config.DefaultConfig.BufferSize),
	}
	indexRet, err := file.index.Seek(0, 0)
	if err != nil {
		panic("get file ret error :" + err.Error())
	}
	dataRet, err := file.data.Seek(0,0)
	if err != nil {
		panic("get file ret error :" + err.Error())
	}
	generalRet, err := file.general.Seek(0,0)
	if err != nil {
		panic("get file ret error :" + err.Error())
	}
	file.offset = &offset{
		indexOffset:   indexRet,
		dataOffset:    dataRet,
		generalOffset: generalRet,
	}
	file.initDescribe()
	return file
}
// 初始化描述文件---同步
func (f *file)initDescribe()  {
	
}

func (f *file)WriteGeneral(p []byte) (length int, ret int64, err error) {
	ret = f.offset.generalOffset
	length, err = f.buffer.generalBuf.Write(p)
	//length, err = f.general.Write(p)
	f.offset.generalOffset = f.offset.generalOffset + int64(length)
	return length, ret, err
}

func (f *file)WriteIndex(p []byte) (length int, ret int64, err error) {
	ret = f.offset.indexOffset
	//length, err = f.index.Write(p)
	length, err = f.buffer.indexBuf.Write(p)
	f.offset.indexOffset = f.offset.indexOffset + int64(length)
	return length, ret, err
}

func (f *file)WriteData(p []byte) (length int, ret int64, err error) {
	ret = f.offset.dataOffset
	//length, err = f.data.Write(p)
	length, err = f.buffer.dataBuf.Write(p)
	f.offset.dataOffset = f.offset.dataOffset + int64(length)
	return length, ret, err
}

func (f *file)Flush() error {
	var err error
	if f.buffer.dataBuf.Buffered() > 0 {
		WriteLock.DataLock.Lock()
		err = f.buffer.dataBuf.Flush()
		WriteLock.DataLock.Unlock()
	}

	if f.buffer.indexBuf.Buffered() > 0 {
		WriteLock.IndexLock.Lock()
		err = f.buffer.indexBuf.Flush()
		WriteLock.IndexLock.Unlock()
	}
	if f.buffer.generalBuf.Buffered() > 0{
		WriteLock.GeneralLock.Lock()
		err = f.buffer.generalBuf.Flush()
		WriteLock.GeneralLock.Unlock()
	}
	return err
}

func (f *file)Close()  {
	_ = f.data.Close()
	_ = f.index.Close()
	_ = f.general.Close()
	_ = f.describe.Close()
}