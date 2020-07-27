package logger

import (
	"encoding/json"
	"fmt"
	"github.com/tizx/xvlog/logdata"
	"os"
	"path"
)

type console struct {}

func NewConsole() Handler {
	return &console{}
}

func (c *console) handle(log *logdata.Log) {
	data,_ := json.Marshal(log.Data())
	str := fmt.Sprint(log.Field(),"  ", string(data))
	time := log.Time().Format("2006-01-02 15:04:05")
	fmt.Fprintf(os.Stdout,
		"%s [%s] %s %s %d\n",
		time,
		logdata.LevelName[log.Level()],
		str,
		path.Base(log.File()),
		log.Line())
}
