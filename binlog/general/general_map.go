package general

import "sync"

type Cache struct {
	generalMap map[General]uint32
	generalMapLock sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		generalMap: make(map[General]uint32),
	}
}


func (c *Cache)Get(general *General) (uint32, bool) {
	c.generalMapLock.RLock()
	defer c.generalMapLock.RUnlock()
	id, ok := c.generalMap[*general];
	return id, ok
}

func (c *Cache)Set(general *General, id uint32)  {
	c.generalMapLock.Lock()
	defer c.generalMapLock.Unlock()
	c.generalMap[*general] = id
}