package logdata

import (
	"runtime"
	"time"
)

type Log struct {
	time   time.Time     //日志时间
	level  Level         //日志等级
	field string        // 格式化字符串
	data      map[string]interface{} // 格式化变量
	file   string        // 打印的文件
	line   int           // 打印行号
}

func (l *Log) Level() Level {
	return l.level
}

func NewLog(level Level, field string, h map[string]interface{}) *Log {
	_, file, line, _ := runtime.Caller(2)
	return &Log{
		level:  level,
		time:   time.Now(),
		field: field,
		data:      h,
		file:   file,
		line:   line,
	}
}

func (l *Log) Line() int {
	return l.line
}

func (l *Log) File() string {
	return l.file
}

func (l *Log) Data() map[string]interface{} {
	return l.data
}

func (l *Log) Field() string {
	return l.field
}

func (l *Log) Time() time.Time {
	return l.time
}
