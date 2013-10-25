package basellog

//日志接口
type LlogInterface interface {
	InitLog(configInfo []byte) (level int, err error) //日志初始化
	WriteLog(level int, msg string) (err error)       //写日志
	CloseLog()                                        //关闭日志
}
