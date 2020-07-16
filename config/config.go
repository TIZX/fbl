package config

type c struct {
	Logger int	// 日志写入方式
	LogPath string //日志保存路径
	BufferSize int //缓冲区大小
	IsInstant bool //是否即时写入---
}

var Config c

func init()  {
	Config.Logger = 1
	Config.LogPath = "log"
	Config.BufferSize = 10000
	Config.IsInstant = true
}



