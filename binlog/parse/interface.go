package parse

type Decoder interface {
	Encode(value interface{}) (typeByte, nameByte, valueByte []byte)
}
