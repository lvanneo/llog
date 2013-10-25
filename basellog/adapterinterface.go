package basellog

//适配器接口
type AdapterInterface interface {
	InitLlogger(channelLength int64)                        //日志适配器初始化
	SetLloggers(configinfo []byte) error                    //通过配置信息配置日志
	SetLlogger(adaptername string, configinfo []byte) error //指定适配器，通过配置信息配置日志
	Debug(val ...interface{})
	Info(val ...interface{})
	Warn(val ...interface{})
	Error(val ...interface{})
	Fatal(val ...interface{})
	Debugf(format string, val ...interface{})
	Infof(format string, val ...interface{})
	Warnf(format string, val ...interface{})
	Errorf(format string, val ...interface{})
	Fatalf(format string, val ...interface{})
}
