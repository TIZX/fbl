package binlog

import "sync"

var generalMap map[string]int64 = make(map[string]int64)

var generalMapLock sync.RWMutex

func generalGet(key []byte) (int64, bool) {
	generalMapLock.RLock()
	defer generalMapLock.RUnlock()
	value, ok := generalMap[string(key)]
	return value, ok

}

func generalPut(key []byte, value int64) {
	generalMapLock.Lock()
	defer generalMapLock.Unlock()
	generalMap[string(key)] = value
}
