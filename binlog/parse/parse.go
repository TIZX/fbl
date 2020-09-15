package parse

import (
	"encoding/binary"
	"reflect"
)

// 编码
func Encode(value map[string]interface{}) (typeName, valueByte []byte) {
	typeByte, nameByte, valueByte := MapParse[reflect.Map].Encode(value)
	typeLen := make([]byte, 4)
	binary.BigEndian.PutUint32(typeLen, uint32(len(typeByte)))
	typeName = append(typeName, typeLen...)
	typeName = append(typeName, typeByte...)
	typeName = append(typeName, nameByte...)
	return typeName, valueByte
}
