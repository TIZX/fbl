package rw

type Write interface {
	WriteGeneral(p []byte) (len int, ret int64, err error)
	WriteIndex(p []byte) (len int, ret int64, err error)
	WriteData(p []byte) (len int, ret int64, err error)
	Flush() error
	Close()
}
type Read interface {
	ReadGeneral(p []byte) (int, error)
	ReadIndex(p []byte) (int, error)
	ReadData(p []byte) (int, error)
}

type ReadWrite interface {
	Write
	Read
}
