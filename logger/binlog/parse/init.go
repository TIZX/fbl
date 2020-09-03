package parse

import "reflect"

var MapParse map[reflect.Kind]TypeParse

const us uint8 = 0x1E

func init()  {
	MapParse = make(map[reflect.Kind]TypeParse)
	MapParse[reflect.Bool] = Bool{}
	MapParse[reflect.Int] = Int{}
	MapParse[reflect.Int8] = Int8{}
	MapParse[reflect.Int16] = Int16{}
	MapParse[reflect.Int32] = Int32{}
	MapParse[reflect.Int64] = Int64{}
	MapParse[reflect.Uint] = Uint{}
	MapParse[reflect.Uint8] = Uint8{}
	MapParse[reflect.Uint16] = Uint16{}
	MapParse[reflect.Uint32] = Uint32{}
	MapParse[reflect.Uint64] = Uint64{}

	MapParse[reflect.Uintptr] = Uintptr{}
	MapParse[reflect.Float32] = Float32{}
	MapParse[reflect.Float64] = Float64{}
	MapParse[reflect.Complex64] = Complex64{}
	MapParse[reflect.Complex128] = Complex128{}

	MapParse[reflect.Chan] = Empty{}
	MapParse[reflect.Func] = Empty{}

	MapParse[reflect.Interface] = Empty{}
	MapParse[reflect.Map] = Map{}
	MapParse[reflect.Ptr] = Ptr{}
	MapParse[reflect.Slice] = Slice{}
	MapParse[reflect.String] = String{}
	MapParse[reflect.Struct] = Struct{}
	MapParse[reflect.UnsafePointer] = UnsafePointer{}
}