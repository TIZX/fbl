package fbl

// 退出前同步日志
func (l *logger)SyncAndClose()  {
	for _,v :=  range l.processorChan{
		close(v)
	}
	l.processing.Wait() // 等待所有处理器关闭
	l.processor.SyncAndClose()
}
