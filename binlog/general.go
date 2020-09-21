package binlog

import (
	"encoding/binary"
)

const us uint8 = 0x1E

type general struct {
	ID           uint32
	Length       uint32
	Level        int8   //日志等级
	Message      string // 格式化字符串
	TypeNameByte string
	File         string // 打印的文件
	Line         int    // 打印行号
}

// 长度/ID/level/File/Line/Message/TypeNameByte
func (g *general) Encode() []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint32(res[4:8], uint32(g.ID))
	res = append(res, uint8(g.Level))    // 封装level
	res = append(res, []byte(g.File)...) // 封装文件路径
	res = append(res, us)                // 添加分隔符
	lineByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lineByte, uint32(g.Line))
	res = append(res, lineByte...)          // 添加行数
	res = append(res, []byte(g.Message)...) // 添加描述字段
	res = append(res, us)                   // 添加分隔符
	res = append(res, g.TypeNameByte...)
	binary.BigEndian.PutUint32(res[0:4], uint32(len(res))) // 写入长度
	return res
}

// 长度/ID/level/File/Line/Message/TypeNameByte
func (g *general) Decode(data []byte) {
	g.Length = binary.BigEndian.Uint32(data[0:4])
	g.ID = binary.BigEndian.Uint32(data[4:8])
	g.Level = int8(data[8])
	var i = 9
	for i < len(data) {
		if data[i] == us {
			break
		}
		i++
	}
	g.File = string(data[9:i])
	i++
	g.Line = int(binary.BigEndian.Uint32(data[i : i+4]))
	i += 4
	var j = i
	for j < len(data) {
		if data[j] == us {
			break
		}
		j++
	}
	g.Message = string(data[i:j])
	i = j + 1
	g.TypeNameByte = string(data[i:])
}
