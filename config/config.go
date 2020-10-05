package config

import "time"

const (
	LocalFile int = iota + 1
	RPCServer
)
const (
	SizeSplit int8 = iota + 1
	DateSplit
)

type LocalFileConfig struct {
	LogPath         string        //日志保存路径
	SplitMethod     int8 // 1:日期分割 、 2:大小分割
}

type RPCServerConfig struct {

}


type Config struct {
	Processor       uint8         // 日志写入方式 // 0标准输入(控制台)1文件写入(二进制编码)
	BufferSize      int           //缓冲区大小
	IsInstant       bool          //是否即时写入---
	FlushDuration   time.Duration // 隔多长时间自动调用buffer.Flush
	ProcessorNumber int  // 处理器数量
	BinlogReceiver int // 日志的写入方式
	LocalFileConfig *LocalFileConfig
	RPCServerConfig *RPCServerConfig
}

var DefaultConfig *Config

func init() {
	DefaultConfig = &Config{
		Processor:       0,
		BufferSize:      4096,
		FlushDuration:   60 * time.Second,
		ProcessorNumber: 100,
		BinlogReceiver:  LocalFile,
		LocalFileConfig: &LocalFileConfig{
			LogPath:         "./log",
			SplitMethod:     DateSplit,
		},
		RPCServerConfig: nil,
	}
}
