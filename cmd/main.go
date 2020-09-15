package main

import (
	log "github.com/tizx/xvlog"
	"github.com/tizx/xvlog/config"
	_ "net/http/pprof"
	"time"
)

func main() {
	logger := log.NewLogger(&config.Config{
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
	for i := 0; i < 100; i++ {
		logger.WithFields(log.Map{
			"value": i,
		}).Info("test info")
	}
	time.Sleep(15 * time.Second)
}
