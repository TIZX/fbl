package logdata

import (
	"runtime"
	"time"
)

type Log struct {
	time   time.Time     //日志时间
	level  Level         //日志等级
	format string        // 格式化字符串
	a      []interface{} // 格式化变量
	file   string        // 打印的文件
	line   int           // 打印行号
}

func (l *Log) Level() Level {
	return l.level
}

func NewLog(level Level, format string, a []interface{}) *Log {
	_, file, line, _ := runtime.Caller(2)
	return &Log{
		level:  level,
		time:   time.Now(),
		format: format,
		a:      a,
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

func (l *Log) A() []interface{} {
	return l.a
}

func (l *Log) Format() string {
	return l.format
}

func (l *Log) Time() time.Time {
	return l.time
}
