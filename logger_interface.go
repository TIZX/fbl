package xvlog

import "github.com/tizx/xvlog/logdata"

type Logger interface {
	Write()
	Receive(log *logdata.Log) //接收日志条目
}
