package rw

import "sync"

type lock struct {
	IndexLock sync.Mutex
	DataLock sync.Mutex
	GeneralLock sync.Mutex
}

var WriteLock *lock

func init()  {
	WriteLock = &lock{}
}