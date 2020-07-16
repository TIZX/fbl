package xvlog

import "xvlog/config"

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
