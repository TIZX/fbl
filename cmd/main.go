package main

import (
	log "github.com/tizx/fbl"
	"github.com/tizx/fbl/config"
	_ "net/http/pprof"
)

func main() {
	logger := log.NewLogger(&config.Config{
		Processor:     1,
		LogPath:    "./log",
		BufferSize: 4096,
		ProcessorNumber:  10,
	})

	for i := 0; i < 100; i++ {
		logger.WithFields(log.Map{
			"value": i,
		}).Info("test info")
	}
	defer logger.SyncAndClose()
}
