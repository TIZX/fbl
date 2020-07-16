package main

import (
	"time"
	"xvlog"
)

func main() {
	xvlog.SetLogger(1)
	xvlog.Logger()
	//var i = 0
	//xvlog.Info("数据%s:%d", "测试", i)
	go func() {
		var i = 0
		for {
			time.Sleep(1 * time.Millisecond)
			xvlog.Info("数据%s:%d", "测试", i)
			i++
		}
	}()

	select {}
}
