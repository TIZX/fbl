package binlog

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

type generalMap struct {
	cacheFile *os.File
	cacheFileBuf *bufio.Writer
	generalMap         map[general]int
	generalMapLock     sync.RWMutex
	generalCounter     int        // general计数器-计算id
	generalCounterLock sync.Mutex //计数器锁
	length             int
}

func newGeneralMap(data []byte) *generalMap {
	gMap := &generalMap{
		generalMap: make(map[general]int),
	}

	var err error
	gMap.cacheFile, err = os.OpenFile("", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, os.ModePerm)
	if err != nil {
		// 打开缓存文件失败
	}
	gMap.cacheFileBuf = bufio.NewWriter(gMap.cacheFile)
	gMap.cacheFileBuf.Write(data)
	gMap.cacheFileBuf.Flush()

	gMap.length = len(data)
	for i := 0; i < gMap.length; {
		itemLen := binary.BigEndian.Uint32(data[i : i+4])
		itemByte := data[i : i+int(itemLen)]
		item := &general{}
		item.Decode(itemByte)
		//temp.Length = itemLen
		i = i + int(itemLen)
		gMap.generalMap[*item] = int(item.ID)
		gMap.counter()
	}
	return gMap
}

func (gMap *generalMap) putGeneral(g *general) int {
	gMap.generalMapLock.Lock()
	defer gMap.generalMapLock.Unlock()
	ID := gMap.counter()
	gMap.generalMap[*g] = ID
	g.ID = uint32(ID)
	gMap.cacheFileBuf.Write(g.Encode())
	return ID
}

func (gMap *generalMap) getGeneralID(g *general) int {
	gMap.generalMapLock.RLock()
	ID, ok := gMap.generalMap[*g]
	gMap.generalMapLock.RUnlock()
	if ok {
		return ID
	}
	return gMap.putGeneral(g)
}

func (gMap *generalMap) counter() int {
	gMap.generalCounterLock.Lock()
	defer gMap.generalCounterLock.Unlock()
	gMap.generalCounter = gMap.generalCounter + 1
	return gMap.generalCounter
}

func (gMap *generalMap) toByte() []byte {
	data := make([]byte, 0)
	for general, _ := range gMap.generalMap {
		data = append(data, general.Encode()...)
	}
	return data
}
