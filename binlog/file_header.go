package binlog

import (
	"encoding/binary"
	"sync/atomic"
)

type fileHeader struct {
	headerByte []byte
	updateChan chan bool
	dataOffset    uint64 //位置0
	generalOffset uint64 //位置8
}

func newFileHeader() *fileHeader {
	header := &fileHeader{}
	header.headerByte = make([]byte, 16)
	return header
}

func (f *fileHeader) dataOffsetAdd(dataOffset uint64) {
	atomic.AddUint64(&f.dataOffset, dataOffset)
}


func (f *fileHeader) generalOffsetAdd(generalOffset uint64) {
	atomic.AddUint64(&f.generalOffset, generalOffset)
}

func (f *fileHeader)ParseByte()  {
	f.dataOffset = binary.BigEndian.Uint64(f.headerByte[0:8])
	if f.dataOffset == 0 {
		f.dataOffset = 16
	}
	f.generalOffset = binary.BigEndian.Uint64(f.headerByte[8:16])
	if f.generalOffset == 0 {
		f.generalOffset = 16
	}
}

func (f *fileHeader)ToByte() []byte {
	binary.BigEndian.PutUint64(f.headerByte[0:8], f.dataOffset)
	binary.BigEndian.PutUint64(f.headerByte[8:16], f.generalOffset)
	return f.headerByte
}
