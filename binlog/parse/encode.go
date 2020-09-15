package parse

import (
	"encoding/binary"
	"math"
	"reflect"
)

func (Bool) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 1)
	if value.(bool) {
		valueByte[0] = 1
	} else {
		valueByte[0] = 0
	}
	return append(typeByte, uint8(reflect.Bool)), nameByte, valueByte
}

func (Int) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, uint32(value.(int)))
	return append(typeByte, uint8(reflect.Int)), nameByte, valueByte
}

func (Int8) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = append(valueByte, byte(value.(int8)))
	return append(typeByte, uint8(reflect.Int8)), nameByte, valueByte
}

func (Int16) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 2)
	binary.BigEndian.PutUint16(valueByte, uint16(value.(int16)))
	return append(typeByte, uint8(reflect.Int16)), nameByte, valueByte
}

func (Int32) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, uint32(value.(int32)))
	return append(typeByte, uint8(reflect.Int32)), nameByte, valueByte
}

func (Int64) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 8)
	binary.BigEndian.PutUint64(valueByte, uint64(value.(int64)))
	return append(typeByte, uint8(reflect.Int64)), nameByte, valueByte
}

func (Uint) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, uint32(value.(uint)))
	return append(typeByte, uint8(reflect.Uint)), nameByte, valueByte
}

func (Uint8) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = append(valueByte, value.(uint8))
	return append(typeByte, uint8(reflect.Uint8)), nameByte, valueByte
}

func (Uint16) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 2)
	binary.BigEndian.PutUint16(valueByte, value.(uint16))
	return append(typeByte, uint8(reflect.Uint16)), nameByte, valueByte
}

func (Uint32) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, value.(uint32))
	return append(typeByte, uint8(reflect.Uint32)), nameByte, valueByte
}

func (Uint64) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 8)
	binary.BigEndian.PutUint64(valueByte, value.(uint64))
	return append(typeByte, uint8(reflect.Uint64)), nameByte, valueByte
}

func (Float32) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, math.Float32bits(value.(float32)))
	return append(typeByte, uint8(reflect.Float32)), nameByte, valueByte
}

func (Float64) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	valueByte = make([]byte, 8)
	binary.BigEndian.PutUint64(valueByte, math.Float64bits(value.(float64)))
	return append(typeByte, uint8(reflect.Float64)), nameByte, valueByte
}

func (Complex64) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	realValue := make([]byte, 4)
	imagValue := make([]byte, 4)
	binary.BigEndian.PutUint32(realValue, math.Float32bits(real(value.(complex64))))
	binary.BigEndian.PutUint32(imagValue, math.Float32bits(imag(value.(complex64))))
	valueByte = append(valueByte, realValue...)
	valueByte = append(valueByte, imagValue...)
	return append(typeByte, uint8(reflect.Complex64)), nameByte, valueByte
}

func (Complex128) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	realValue := make([]byte, 8)
	imagValue := make([]byte, 8)
	binary.BigEndian.PutUint64(realValue, math.Float64bits(real(value.(complex128))))
	binary.BigEndian.PutUint64(imagValue, math.Float64bits(imag(value.(complex128))))
	valueByte = append(valueByte, realValue...)
	valueByte = append(valueByte, imagValue...)
	return append(typeByte, uint8(reflect.Complex128)), nameByte, valueByte
}

func (String) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	str := append([]byte(value.(string)), us)
	return append(typeByte, uint8(reflect.String)), nameByte, str
}

func (Slice) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := reflect.ValueOf(arg)
	sliceLen := value.Len()
	if sliceLen == 0 {
		return
	}
	// 写入切片类型的typeByte
	typeByte = append(typeByte, uint8(reflect.Slice))

	// 写入切片长度到valueByte
	sliceLenByte := make([]byte, 4)
	binary.BigEndian.PutUint32(sliceLenByte, uint32(sliceLen))
	valueByte = append(valueByte, sliceLenByte...)

	for i := 0; i < sliceLen; i++ {
		indexValue := value.Index(i)
		indexTypeByte, indexNameByte, indexValueByte := MapParse[indexValue.Kind()].Encode(indexValue.Interface())
		if i == 0 {
			typeByte = append(typeByte, indexTypeByte...)
		}
		nameByte = append(nameByte, indexNameByte...)
		valueByte = append(valueByte, indexValueByte...)
	}
	return typeByte, nameByte, valueByte

}

func (Uintptr) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := reflect.Indirect(reflect.ValueOf(arg))
	return MapParse[value.Kind()].Encode(value.Interface())
}

func (UnsafePointer) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := reflect.Indirect(reflect.ValueOf(arg))
	return MapParse[value.Kind()].Encode(value.Interface())
}

func (Ptr) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := reflect.Indirect(reflect.ValueOf(arg))
	return MapParse[value.Kind()].Encode(value.Interface())
}

func (Struct) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := reflect.ValueOf(arg)
	//get struct type
	reflectType := value.Type()
	typeByte = append(typeByte, uint8(reflect.Struct))

	numField := value.NumField()
	var numFieldByte []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(numFieldByte, uint32(numField))
	typeByte = append(typeByte, numFieldByte...)
	for i := 0; i < numField; i++ {
		field := reflect.Indirect(value.Field(i))
		fieldTypeByte, fieldNameByte, fieldValueByte := MapParse[field.Kind()].Encode(field.Interface())
		typeByte = append(typeByte, fieldTypeByte...)
		fieldName := reflectType.Field(i).Name
		nameByte = append(nameByte, []byte(fieldName)...)
		nameByte = append(nameByte, us)
		nameByte = append(nameByte, fieldNameByte...)
		valueByte = append(valueByte, fieldValueByte...)
	}
	return typeByte, nameByte, valueByte
}

func (Map) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := arg.(map[string]interface{})
	typeByte = append(typeByte, uint8(reflect.Map))
	// 写入map长度
	lenByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lenByte, uint32(len(value)))
	typeByte = append(typeByte, lenByte...)

	var sortKey []string = make([]string, len(value))
	var i = 0
	for k := range value {
		sortKey[i] = k
		i++
	}

	for _, key := range sortKey {
		value := value[key]
		nameByte = append(nameByte, []byte(key)...)
		nameByte = append(nameByte, us)
		iterTypeByte, iterNameByte, iterValueByte := MapParse[kind(value)].Encode(value)
		typeByte = append(typeByte, iterTypeByte...)
		nameByte = append(nameByte, iterNameByte...)
		valueByte = append(valueByte, iterValueByte...)
	}
	return typeByte, nameByte, valueByte
}

func (Empty) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	return
}
