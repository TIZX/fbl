package binlog

import "encoding/binary"

// log data
type logData struct {
	length uint32
	generalID uint32
	time uint32
	dataByte []byte
}

func (d *logData)Encode() []byte {
	res := make([]byte, 12)
	d.length = 12+uint32(len(d.dataByte))
	binary.BigEndian.PutUint32(res[0:4], d.length)
	binary.BigEndian.PutUint32(res[4:8], d.generalID)
	binary.BigEndian.PutUint32(res[8:12], d.time)
	res = append(res, d.dataByte...)
	return res
}
