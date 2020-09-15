package logdata

import (
	"runtime"
	"time"
)

type Log struct {
	Time    time.Time              //日志时间
	Level   Level                  //日志等级
	Message string                 // 格式化字符串
	Fields  map[string]interface{} // 格式化变量
	File    string                 // 打印的文件
	Line    int                    // 打印行号
}

func NewLog() *Log {
	_, file, line, _ := runtime.Caller(2)
	return &Log{
		Time:   time.Now(),
		Fields: make(map[string]interface{}),
		File:   file,
		Line:   line,
	}
}
