package binlog

import (
	"encoding/binary"
	"github.com/tizx/xvlog/binlog/parse"
	"github.com/tizx/xvlog/binlog/rw"
	"github.com/tizx/xvlog/logdata"
)

type Index uint32

var IndexTotal Index = 1

// 公共数据位置
var MapIndex = make(map[string]Index)

const us uint8 = 0x1E

type Structure struct {
	general []byte // 常规公共数据 // 二次封装 /4 长度 /

	indexByte [24]byte // 8 general位置 / 8 data位置 / 4 data长度 / 4 time时间
	data      []byte   // data
	log       *logdata.Log
}

func NewStructure(log *logdata.Log) *Structure {
	return &Structure{
		general: make([]byte, 0),
		data:    make([]byte, 0),
		log:     log}
}

func (s *Structure) Parse() {
	s.makeGeneral()

	_, ret, err := s.writeGeneral()
	binary.BigEndian.PutUint64(s.indexByte[:8], uint64(ret))

	dataLen, ret, err := s.writeData()
	if err != nil {
		panic("write data error: " + err.Error())
	}
	binary.BigEndian.PutUint64(s.indexByte[8:16], uint64(ret))
	binary.BigEndian.PutUint32(s.indexByte[16:20], uint32(dataLen))
	binary.BigEndian.PutUint32(s.indexByte[20:], uint32(s.log.Time.Second()))
	s.writeIndex()
}

func (s *Structure) makeGeneral() {
	s.general = append(s.general, '0', '0', '0', '0') // 四个字节占位

	s.general = append(s.general, uint8(s.log.Level))    // 封装level
	s.general = append(s.general, []byte(s.log.File)...) // 封装文件路径
	s.general = append(s.general, us)                    // 添加分隔符
	lineByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lineByte, uint32(s.log.Line))
	s.general = append(s.general, lineByte...)              // 添加行数
	s.general = append(s.general, []byte(s.log.Message)...) // 添加描述字段
	s.general = append(s.general, us)                       // 添加分隔符

	typeNameByte := make([]byte, 0)
	typeNameByte, s.data = parse.Encode(s.log.Fields)

	s.general = append(s.general, typeNameByte...)

	binary.BigEndian.PutUint32(s.general[:4], uint32(len(s.general)-4)) // 写入不包括长度字段后的数据长度
}

func (s *Structure) writeGeneral() (length int, ret int64, err error) {
	ret, ok := generalGet(s.general)
	if !ok {
		// 没有缓存到这个字段 // 加锁
		rw.WriteLock.GeneralLock.Lock()
		// 获取再判断
		ret, ok := generalGet(s.general)
		if !ok {
			var err error
			_, ret, err = Write.WriteGeneral(s.general)
			if err != nil {
				panic("write general error: " + err.Error())
			}
		}
		// 缓存这个字段
		generalPut(s.general, ret)
		rw.WriteLock.GeneralLock.Unlock()
	}
	return len(s.general), ret, err
}

func (s *Structure) writeIndex() (len int, ret int64, err error) {
	rw.WriteLock.IndexLock.Lock()
	defer rw.WriteLock.IndexLock.Unlock()
	return Write.WriteIndex(s.indexByte[:])
}
func (s *Structure) writeData() (len int, ret int64, err error) {
	rw.WriteLock.DataLock.Lock()
	defer rw.WriteLock.DataLock.Unlock()
	return Write.WriteData(s.data)
}
