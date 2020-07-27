package xvlog

import "github.com/tizx/xvlog/config"

// 可配置

func SetLogger(v int)  {
	config.Config.Logger = v
}
func SetLogPath(v string)  {
	config.Config.LogPath = v
}
func SetBufferSize(v int)  {
	config.Config.BufferSize = v
}
func SetIsInstant(v bool)  {
	config.Config.IsInstant = v
}
