package logger

import "xvlog/logdata"

type Logger interface {
	Write()
	Receive(log *logdata.Log)
	Exit()

}

type Handler interface {
	handle(log *logdata.Log)
}