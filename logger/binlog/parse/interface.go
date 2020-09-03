package parse

type TypeParse interface {
	Encode(value interface{})(typeByte, nameByte,valueByte []byte)
	Decode(typeByte, nameByte,valueByte []byte) (interface{}, []byte, []byte, []byte)
}

