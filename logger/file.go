package logger

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
	"xvlog/config"
	"xvlog/logdata"
)

type File struct {
	path     string       //文件保存路径
	logFiles [5]*os.File  // 五个等级的文件
	logDay   string       // logFiles所属的日志YYYYMMDD
	logLock  sync.RWMutex // logFiles与logDay的读写锁
}

func NewFile() Handler {
	return  &File{
		path: config.Config.LogPath,
	}
}

func (f *File) handle(log *logdata.Log) {

	file := f.getFile(log.Time(), log.Level())

	if file == nil {
		return
	}
	str := fmt.Sprintf(log.Format(), log.A()...)
	time := log.Time().Format("2006-01-02 15:04:05")
	fmt.Fprintf(file,
		"%s [%s] %s %s %d\n",
		time,
		logdata.LevelName[log.Level()],
		str,
		path.Base(log.File()),
		log.Line())

}

// 获取一个文件
func (f *File)getFile(time time.Time, level logdata.Level) *os.File {
	ymd := time.Format("20060102")
	if f.logDay != ymd || f.logFiles[level] == nil {
		if f.logFiles[level] != nil {
			f.logFiles[level].Close()
		}
		file, err := f.createFile(time, level)
		if err != nil {
			fmt.Println("create file error: ", err)
		}
		f.logFiles[level] = file
		f.logDay = ymd
	}
	return f.logFiles[level]
}

// 创建一个日志文件
func (f *File) createFile(time time.Time, level logdata.Level) (file *os.File, err error) {
	fileName := time.Format("20060102") + "_" + logdata.LevelName[level] + ".log"
	filePath := path.Join(f.path,
		strconv.Itoa(time.Year()),
		time.Month().String(),
		strconv.Itoa(time.Day()))
	err = os.MkdirAll(filePath, 0755)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path.Join(filePath, fileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
}
