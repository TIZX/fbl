package logger

import (
	"fmt"
	"os"
	"path"
	"github.com/tizx/xvlog/logdata"
)

type console struct {}

func NewConsole() Handler {
	return &console{}
}

func (c *console) handle(log *logdata.Log) {
	str := fmt.Sprintf(log.Format(), log.A()...)
	time := log.Time().Format("2006-01-02 15:04:05")
	fmt.Fprintf(os.Stdout,
		"%s [%s] %s %s %d\n",
		time,
		logdata.LevelName[log.Level()],
		str,
		path.Base(log.File()),
		log.Line())
}
