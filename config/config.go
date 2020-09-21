package config

import "time"

type Config struct {
	Processor        uint8         // 日志写入方式 // 0标准输入(控制台)1文件写入(二进制编码)
	LogPath       string        //日志保存路径
	BufferSize    int           //缓冲区大小
	IsInstant     bool          //是否即时写入---
	FlushDuration time.Duration // 隔多长时间自动调用buffer.Flush
	ProcessorNumber     int
	SplitMethod   int8 // 1:日期分割 、 2:大小分割
}

var DefaultConfig *Config

func init() {
	DefaultConfig = &Config{
		Processor:        0,
		LogPath:       "./log",
		BufferSize:    4096,
		FlushDuration: 60 * time.Second,
		ProcessorNumber:     100,
		SplitMethod:   1,
	}
}
