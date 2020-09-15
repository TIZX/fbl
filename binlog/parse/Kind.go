package parse

import "reflect"

func kind(value interface{}) reflect.Kind {
	switch value.(type) {
	case bool:
		return reflect.Bool
	case int:
		return reflect.Int
	case int8:
		return reflect.Int8
	case int16:
		return reflect.Int16
	case int32:
		return reflect.Int32
	case int64:
		return reflect.Int64
	case uint:
		return reflect.Uint
	case uint8:
		return reflect.Uint8
	case uint16:
		return reflect.Uint16
	case uint32:
		return reflect.Uint32
	case uint64:
		return reflect.Uint64
	case float32:
		return reflect.Float32
	case float64:
		return reflect.Float64
	case complex64:
		return reflect.Complex64
	case complex128:
		return reflect.Complex128
	case string:
		return reflect.String
	default:
		return reflect.TypeOf(value).Kind()
	}
}
