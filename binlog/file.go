package binlog

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
)

type logFile struct {
	header        *fileHeader
	generalMap    *generalMap
	file          *os.File
	dataWriteChan chan []byte
	size          int64 // 文件大小，所占字节数

	bufLock            sync.Mutex
	buf                *bufio.Writer
	IsCloseSuccessChan chan bool //退出通道
}

func NewLogFile(filePath string) (*logFile, error) {
	var err error
	file := &logFile{
		header: &fileHeader{
			dataOffset:    16,
			generalOffset: 16,
		},
		dataWriteChan: make(chan []byte),
		IsCloseSuccessChan: make(chan bool, 1),
	}
	_ = os.MkdirAll(path.Dir(filePath), os.ModeDir)
	file.file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	file.buf = bufio.NewWriter(file.file)
	file.readHeader()

	data, err := file.readGeneralAll()

	if err != nil {
		return nil, err
	}

	file.generalMap = newGeneralMap(data)

	return file, nil
}

func (l *logFile) readHeader() {
	l.header = newFileHeader()
	var err error
	_, err = l.file.ReadAt(l.header.headerByte, 0)
	l.header.ParseByte()
	if err != nil && err == io.EOF {
		headerByte := l.header.ToByte()
		l.file.WriteAt(headerByte, 0)
	}
}

// 读取General全部数据
func (l *logFile) readGeneralAll() ([]byte, error) {
	l.file.Seek(int64(l.header.generalOffset), io.SeekStart)
	defer l.file.Seek(int64(l.header.generalOffset), io.SeekStart)

	var data = make([]byte, 0)
	var temp = make([]byte, 1024)
	for {
		n, err := l.file.Read(temp)
		data = append(data, temp[:n]...)
		if err != nil && err == io.EOF {
			break
		}
	}
	return data, nil
}

// 开启一个写goroutine ,写入数据
func (l *logFile) startWrite() {
	for {
		data, ok := <-l.dataWriteChan
		if !ok {
			// 关闭了写数据的chan
			break
		}
		n, _ := l.buf.Write(data)
		l.header.generalOffsetAdd(uint64(n))
	}
	l.close()
}

// 关闭写goroutine的对应chan后退出循环后调用
func (l *logFile) close() {
	//处理数据
	headerByte := l.header.ToByte()

	l.file.WriteAt(headerByte, 0) //写入头

	l.buf.Write(l.generalMap.toByte())
	l.buf.Flush()
	fmt.Println("all data write success")
	l.IsCloseSuccessChan <- true
}
