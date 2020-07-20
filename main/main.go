package main

import (
	"github.com/tizx/xvlog"
	"time"
)

func main() {
	xvlog.SetLogger(0)
	xvlog.Logger()
	go func() {
		var i = 0
		for {

			xvlog.Info("数据%s:%d", "测试", i)
			time.Sleep(1 * time.Millisecond)
			i++
		}
	}()

	select {}
}
