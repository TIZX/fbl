package parse

import (
	"fmt"
	"testing"
)

type s struct {
	ID    int
	Name  string
	Value float64
}

// 编码测试
func TestEncodeDecode(t *testing.T) {
	var (
		data map[string]interface{} = make(map[string]interface{})
		//aa map[string]interface{} = make(map[string]interface{})
		typeNameByte []byte = make([]byte, 0)
		valueByte    []byte = make([]byte, 0)
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

	data["Complex64"] = complex(float32(12.12), float32(-12.45))
	data["Complex128"] = complex(-1254.15452, -11232.454645)
	data["String"] = "test strings"

	mapData := make(map[string]interface{})
	mapData["int"] = int(456)
	mapData["int8"] = int8(12)
	mapData["int16"] = int16(655)
	mapData["int32"] = 655365
	mapData["int64"] = 14654646

	mapData["-int"] = -int(456)
	mapData["-int8"] = -int8(12)
	mapData["-int16"] = -int16(655)
	mapData["-int32"] = -655365
	mapData["-int64"] = -14654646

	mapData["float32"] = 1245.4566
	mapData["-float32"] = -1245.4566
	mapData["float64"] = 12485.45661
	mapData["-float64"] = -124545.4566

	mapData["Complex64"] = complex(float32(12.12), float32(-12.45))
	mapData["Complex128"] = complex(-1254.15452, -11232.454645)
	mapData["String"] = "test strings"
	var name = "test name"
	mapData["ptr"] = &name
	mapData["struct"] = &s{
		ID:    -456,
		Name:  "map test name",
		Value: 456.12,
	}
	SliceData := []string{"data21", "data54", "data123", "data456", "4564", "test", "789798"}
	data["Slice"] = SliceData
	data["map"] = mapData

	typeNameByte, valueByte = Encode(data)
	fmt.Println(typeNameByte)
	fmt.Println(valueByte)
	//value := Decode(typeNameByte, valueByte)
	//fmt.Printf("%v\n", value)
	//fmt.Println(value["String"])
}
