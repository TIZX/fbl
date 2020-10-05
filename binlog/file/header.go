package file

import (
	"encoding/binary"
	"sync/atomic"
)

const headerLen = 16

type header struct {
	headerByte    []byte
	updateChan    chan bool
	dataOffset    uint64 //位置0
	generalOffset uint64 //位置8
}

func newHeader(data []byte) *header {
	fh := &header{
		headerByte: data,
	}
	fh.dataOffset = binary.BigEndian.Uint64(fh.headerByte[0:8])
	if fh.dataOffset == 0 {
		fh.dataOffset = 16
	}
	fh.generalOffset = binary.BigEndian.Uint64(fh.headerByte[8:16])
	if fh.generalOffset == 0 {
		fh.generalOffset = 16
	}
	return fh
}

func (f *header) dataOffsetAdd(dataOffset uint64) {
	atomic.AddUint64(&f.dataOffset, dataOffset)
}

func (f *header) generalOffsetAdd(generalOffset uint64) {
	atomic.AddUint64(&f.generalOffset, generalOffset)
}

func (f *header) ToByte() []byte {
	binary.BigEndian.PutUint64(f.headerByte[0:8], f.dataOffset)
	binary.BigEndian.PutUint64(f.headerByte[8:16], f.generalOffset)
	return f.headerByte
}
