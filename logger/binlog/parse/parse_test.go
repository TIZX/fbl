package parse

import (
	"fmt"
	"testing"
)

type s struct {
	ID int
	Name string
	Value float64
}

// 编码测试
func TestEncodeDecode(t *testing.T) {
	var (
		data map[string]interface{} = make(map[string]interface{})
		//aa map[string]interface{} = make(map[string]interface{})
		typeNameByte []byte = make([]byte,0)
		valueByte []byte  = make([]byte,0)
	)
	data["bool"] = true
	data["int"] = int(456)
	data["int8"] = int8(12)
	data["int16"] = int16(655)
	data["int32"] = 655365
	data["int64"] = 14654646

	data["-int"] = -int(456)
	data["-int8"] = -int8(12)
	data["-int16"] = -int16(655)
	data["-int32"] = -655365
	data["-int64"] = -14654646

	data["float32"] = 1245.4566
	data["-float32"] = -1245.4566
	data["float64"] = 12485.45661
	data["-float64"] = -124545.4566

	data["Complex64"] = complex(float32(12.12),float32(-12.45))
	data["Complex128"] = complex(-1254.15452,-11232.454645)
	data["String"] = "test strings"

	mapdata := make(map[string]interface{})
	mapdata["bool"] = true
	mapdata["int"] = int(456)
	mapdata["int8"] = int8(12)
	mapdata["int16"] = int16(655)
	mapdata["int32"] = 655365
	mapdata["int64"] = 14654646

	mapdata["-int"] = -int(456)
	mapdata["-int8"] = -int8(12)
	mapdata["-int16"] = -int16(655)
	mapdata["-int32"] = -655365
	mapdata["-int64"] = -14654646

	mapdata["float32"] = 1245.4566
	mapdata["-float32"] = -1245.4566
	mapdata["float64"] = 12485.45661
	mapdata["-float64"] = -124545.4566

	mapdata["Complex64"] = complex(float32(12.12),float32(-12.45))
	mapdata["Complex128"] = complex(-1254.15452,-11232.454645)
	mapdata["String"] = "test strings"
	var name = "dsfadsfads"
	mapdata["ptr"] = &name
	mapdata["struct"] = &s{
		ID:    -456,
		Name:  "dsfsdf",
		Value: 456.12,
	}
	SliceData := []string{"data21","data54","data123","dsffa","dsfa","dsfa","sdfads"}
	data["Slice"] = SliceData
	data["map"] = mapdata

	typeNameByte, valueByte = Encode(data)
	fmt.Println(typeNameByte)
	fmt.Println(valueByte)
	value := Decode(typeNameByte, valueByte)
	fmt.Printf("%v\n", value)
	fmt.Println(value["String"])
}

