package binlog

import (
	"github.com/tizx/xvlog/binlog/rw"
	"github.com/tizx/xvlog/config"
	"path"
	"time"
)

var Write rw.Write

func init() {
	fileName := path.Join(config.DefaultConfig.LogPath, time.Now().Format("20060102"))
	Write = rw.NewWrite(fileName)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			<-ticker.C
			Write.Flush()
		}
	}()
}
