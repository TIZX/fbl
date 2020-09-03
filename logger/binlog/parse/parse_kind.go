package parse

import (
	"encoding/binary"
	"math"
	"reflect"
)

type Bool struct {}

func (Bool) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 1)
	if value.(bool) {
		valueByte[0] = 1
	}else{
		valueByte[0] = 0
	}
	return append(typeByte,uint8(reflect.Bool)),nameByte, valueByte
}
func (Bool)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	if valueByte[0] == 1 {
		return true, typeByte[1:], nameByte, valueByte[1:]
	}
	return false, typeByte[1:], nameByte, valueByte[1:]
}

type Int struct {}

func (Int) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, uint32(value.(int)))
	return append(typeByte, uint8(reflect.Int)),nameByte, valueByte
}
func (Int)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int(binary.BigEndian.Uint32(valueByte[0:4])), typeByte[1:], nameByte, valueByte[4:]
}


type Int8 struct {}

func (Int8) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = append(valueByte, byte(value.(int8)))
	return append(typeByte, uint8(reflect.Int8)),nameByte, valueByte
}
func (Int8)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int8(valueByte[0]), typeByte[1:], nameByte,valueByte[1:]
}

type Int16 struct {}

func (Int16) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 2)
	binary.BigEndian.PutUint16(valueByte, uint16(value.(int16)))
	return append(typeByte, uint8(reflect.Int16)),nameByte, valueByte
}
func (Int16)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int16(binary.BigEndian.Uint16(valueByte[0:2])), typeByte[1:], nameByte,valueByte[2:]
}


type Int32 struct {}

func (Int32) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, uint32(value.(int32)))
	return append(typeByte, uint8(reflect.Int32)),nameByte, valueByte
}
func (Int32)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int32(binary.BigEndian.Uint32(valueByte[0:4])), typeByte[1:], nameByte, valueByte[4:]
}


type Int64 struct {}

func (Int64) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 8)
	binary.BigEndian.PutUint64(valueByte, uint64(value.(int64)))
	return append(typeByte,uint8(reflect.Int64)),nameByte, valueByte
}

func (Int64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return int64(binary.BigEndian.Uint64(valueByte[0:8])),typeByte[1:], nameByte, valueByte[8:]
}

type Uint struct {}

func (Uint) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, uint32(value.(uint)))
	return append(typeByte, uint8(reflect.Uint)),nameByte, valueByte
}
func (Uint)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return uint(binary.BigEndian.Uint32(valueByte[0:4])), typeByte[1:], nameByte, valueByte[4:]
}

type Uint8 struct {}

func (Uint8) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = append(valueByte,value.(uint8))
	return append(typeByte,uint8(reflect.Uint8)),nameByte, valueByte
}
func (Uint8)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return uint8(valueByte[0]), typeByte[1:], nameByte,valueByte[1:]
}

type Uint16 struct {}

func (Uint16) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 2)
	binary.BigEndian.PutUint16(valueByte, value.(uint16))
	return append(typeByte,uint8(reflect.Uint16)),nameByte, valueByte
}
func (Uint16)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return binary.BigEndian.Uint16(valueByte[0:2]), typeByte[1:], nameByte,valueByte[2:]
}



type Uint32 struct {}

func (Uint32) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, value.(uint32))
	return append(typeByte, uint8(reflect.Uint32)),nameByte,  valueByte
}
func (Uint32)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return binary.BigEndian.Uint32(valueByte[0:4]), typeByte[1:], nameByte, valueByte[4:]
}

type Uint64 struct {}

func (Uint64) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 8)
	binary.BigEndian.PutUint64(valueByte, value.(uint64))
	return append(typeByte, uint8(reflect.Uint64)), nameByte, valueByte
}
func (Uint64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return uint64(binary.BigEndian.Uint64(valueByte[0:8])), typeByte[1:], nameByte, valueByte[8:]
}


type Float32 struct {}

func (Float32) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueByte, math.Float32bits(value.(float32)))
	return append(typeByte, uint8(reflect.Float32)),nameByte, valueByte
}
func (Float32)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return math.Float32frombits(binary.BigEndian.Uint32(valueByte[0:4])),
		typeByte[1:],
		nameByte,
		valueByte[4:]
}

type Float64 struct {}

func (Float64) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	valueByte = make([]byte, 8)
	binary.BigEndian.PutUint64(valueByte, math.Float64bits(value.(float64)))
	return append(typeByte, uint8(reflect.Float64)),nameByte, valueByte
}
func (Float64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return math.Float64frombits(binary.BigEndian.Uint64(valueByte[0:8])),
		typeByte[1:],
		nameByte,
		valueByte[8:]
}

type Complex64 struct {}

func (Complex64) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	realValue := make([]byte,4)
	imagValue := make([]byte, 4)
	binary.BigEndian.PutUint32(realValue, math.Float32bits(real(value.(complex64))))
	binary.BigEndian.PutUint32(imagValue, math.Float32bits(imag(value.(complex64))))
	valueByte = append(valueByte, realValue...)
	valueByte = append(valueByte, imagValue...)
	return append(typeByte, uint8(reflect.Complex64)),nameByte, valueByte
}

func (Complex64)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	real := math.Float32frombits(binary.BigEndian.Uint32(valueByte[0:4]))
	imag := math.Float32frombits(binary.BigEndian.Uint32(valueByte[4:8]))
	return complex(real,imag), typeByte[1:],nameByte, valueByte[8:]
}


type Complex128 struct {}

func (Complex128) Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	realValue := make([]byte,8)
	imagValue := make([]byte, 8)
	binary.BigEndian.PutUint64(realValue, math.Float64bits(real(value.(complex128))))
	binary.BigEndian.PutUint64(imagValue, math.Float64bits(imag(value.(complex128))))
	valueByte = append(valueByte, realValue...)
	valueByte = append(valueByte, imagValue...)
	return append(typeByte, uint8(reflect.Complex128)),nameByte, valueByte
}
func (Complex128)Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte) {
	real := math.Float64frombits(binary.BigEndian.Uint64(valueByte[0:8]))
	imag := math.Float64frombits(binary.BigEndian.Uint64(valueByte[8:16]))
	return complex(real,imag), typeByte[1:],nameByte, valueByte[16:]
}

type String struct {}

func (String)Encode(value interface{}) (typeByte, nameByte,valueByte []byte) {
	str := append([]byte(value.(string)), us)
	return append(typeByte, uint8(reflect.String)),nameByte, str
}
func (String) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	var (
		i int
		v byte
	)
	for i, v = range valueByte {
		if v == us {
			break
		}
	}
	return string(valueByte[:i]), typeByte[1:], nameByte, valueByte[i+1:]
}



type Slice struct {}

func (Slice) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := reflect.ValueOf(arg)
	sliceLen := value.Len()
	if sliceLen == 0{
		return
	}
	// 写入切片类型的typeByte
	typeByte = append(typeByte, uint8(reflect.Slice))

	// 写入切片长度到valueByte
	sliceLenByte := make([]byte, 4)
	binary.BigEndian.PutUint32(sliceLenByte, uint32(sliceLen))
	valueByte = append(valueByte, sliceLenByte...)

	for i:=0;i<sliceLen;i++ {
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
func (Slice)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte)  {
	len := binary.BigEndian.Uint32(valueByte[:4])
	valueByte = valueByte[4:]
	resValue := make([]interface{}, len)

	itemType := reflect.Kind(typeByte[1])
	for i := uint32(0); i<len; i++{
		resValue[i], _, nameByte,valueByte = MapParse[itemType].Decode(typeByte[1:],nameByte, valueByte)
	}

	return resValue, typeByte[2:], nameByte, valueByte
}

type Uintptr struct {}

func (Uintptr)Encode(arg interface{}) (typeByte, nameByte, valueByte []byte)  {
	value := reflect.Indirect(reflect.ValueOf(arg))
	return MapParse[value.Kind()].Encode(value.Interface())
}
func (Uintptr)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}



type UnsafePointer struct {}

func (UnsafePointer)Encode(arg interface{}) (typeByte, nameByte, valueByte []byte)  {
	value := reflect.Indirect(reflect.ValueOf(arg))
	return MapParse[value.Kind()].Encode(value.Interface())
}
func (UnsafePointer)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}

type Ptr struct {}

func (Ptr) Encode(arg interface{}) (typeByte, nameByte, valueByte []byte) {
	value := reflect.Indirect(reflect.ValueOf(arg))
	return MapParse[value.Kind()].Encode(value.Interface())
}
func (Ptr)Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}


type Struct struct {}

func (Struct) Encode(arg interface{}) (typeByte ,nameByte, valueByte []byte) {
	value := reflect.ValueOf(arg)
	//get struct type
	reflectType := value.Type()
	typeByte = append(typeByte, uint8(reflect.Struct))

	numField := value.NumField()
	var numFieldByte []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(numFieldByte, uint32(numField))
	typeByte = append(typeByte, numFieldByte...)
	for i:=0;i<numField; i++ {
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
func (Struct) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte){
	numField := binary.BigEndian.Uint32(typeByte[1:5])
	var resValue = make(map[string]interface{})
	typeByte = typeByte[5:]
	var (
		k int
		v byte
		value interface{}
	)
	for i:=uint32(0);i<numField;i++ {
		for k,v = range nameByte {
			if v == us {
				break
			}
		}
		name := string(nameByte[:k])
		value, typeByte, nameByte, valueByte = MapParse[reflect.Kind(typeByte[0])].Decode(typeByte, nameByte[k+1:], valueByte)
		resValue[name] = value
	}
	return resValue, typeByte, nameByte, valueByte
}


type Map struct {}

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
		iterTypeByte, iterNameByte , iterValueByte := MapParse[kind(value)].Encode(value)
		typeByte = append(typeByte, iterTypeByte...)
		nameByte = append(nameByte, iterNameByte...)
		valueByte = append(valueByte, iterValueByte...)
	}
	return typeByte, nameByte, valueByte
}
func (Map) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	typeByte = typeByte[1:]
	// map长度
	len := binary.BigEndian.Uint32(typeByte[:4])
	typeByte = typeByte[4:]
	var (
		k int
		v byte
		i uint32
		value interface{}
		resValue map[string]interface{} = make(map[string]interface{})
	)
	for i = 0;i<len;i++ {
		for k,v = range nameByte {
			if v == us {
				break
			}
		}
		name := string(nameByte[:k])
		value, typeByte, nameByte, valueByte = MapParse[reflect.Kind(typeByte[0])].Decode(typeByte, nameByte[k+1:], valueByte)
		resValue[name] = value
	}
	return resValue, typeByte, nameByte, valueByte
}




type Empty struct {}
func (Empty) Decode(typeByte, nameByte, valueByte []byte) (interface{}, []byte, []byte, []byte) {
	return nil,nil,nil,nil
}

func (Empty) Encode(value interface{}) (typeByte, nameByte, valueByte []byte) {
	return
}