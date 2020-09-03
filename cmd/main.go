package main

import (
	"github.com/tizx/xvlog"
	"github.com/tizx/xvlog/config"
	_ "net/http/pprof"
	"strconv"
	"time"
)

func main() {
	logger := xvlog.NewLogger(&config.Config{
		Logger:     1,
		LogPath:    "./log",
		BufferSize: 4096,
		Processor:  10,
	})
	//for k:=0;k<1000;k++ {
	//	go func() {
	//		goValue := k
	//		i := 0
	//		for {
	//			i++
	//			logger.Error("error", xvlog.H{"value": i, "go": goValue} )
	//		}
	//	}()
	//}
	//select {
	//
	//}
	for i:=0;i<100;i++{
		logger.Error("error"+strconv.Itoa(i), xvlog.H{"value": "dsfasdfsdfsd", "go": "dsfdsafsdfsdfdsf"} )
	}
	time.Sleep(15*time.Second)
}
