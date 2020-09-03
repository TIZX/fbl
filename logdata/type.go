package logdata



//等级类型
type Level uint8

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var LevelName [5]string

func init() {
	LevelName[DEBUG] = "DEBUG"
	LevelName[INFO] = "INFO"
	LevelName[WARN] = "WARN"
	LevelName[ERROR] = "ERROR"
	LevelName[FATAL] = "FATAL"
}
