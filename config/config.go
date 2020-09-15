package config

import "time"

type Config struct {
	Logger        uint8         // 日志写入方式 // 0标准输入(控制台)1文件写入(二进制编码)
	LogPath       string        //日志保存路径
	BufferSize    int           //缓冲区大小
	IsInstant     bool          //是否即时写入---
	FlushDuration time.Duration // 隔多长时间自动调用buffer.Flush
	Processor     int
}

var DefaultConfig *Config

func init() {
	DefaultConfig = &Config{
		Logger:        0,
		LogPath:       "./log",
		BufferSize:    4096,
		FlushDuration: 60 * time.Second,
		Processor:     100,
	}
}
