package parse

import (
	"encoding/binary"
	"reflect"
)

// 解码
func Decode(typeName, valueByte []byte) map[string]interface{} {
	typeLen := binary.BigEndian.Uint32(typeName[0:4])
	typeByte := typeName[4:typeLen+4]

	nameByte := typeName[typeLen+4:]

	value, _,_,_ := MapParse[reflect.Map].Decode(typeByte, nameByte, valueByte)
	return value.(map[string]interface{})
}
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
