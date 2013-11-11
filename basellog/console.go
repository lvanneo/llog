package basellog

//控制台日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com

import (
	"github.com/lvanneo/llog/config"
	"log"
	"os"
)

//控制台日志结构体
type ConseleLog struct {
	logger    *log.Logger
	Level     int
	ShortFile bool
}

//控制台日志初始化
//注册控制台日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func init() {
	Register(AdapterConsole, NewConseLog) //注册控制台日志
}

//创建控制台日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func NewConseLog() LlogInterface {
	clog := new(ConseleLog)
	clog.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	clog.Level = LevelALL
	return clog
}

//初始化控制台日志
//通过配置信息初始化控制台日志
//configInfo 	配置信息
//level 		返回配置信息的日志等级及错误信息，错误则返回最低等级及错误信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *ConseleLog) InitLog(configInfo []byte) (level int, err error) {

	conf, err := config.GetConsoleConfig(configInfo)
	if err != nil {
		return LevelALL, err
	}

	/*
		this.ShortFile = conf.ShortFile
		if this.ShortFile {
			this.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			this.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
		}
	*/

	this.Level, err = CheckLevel(conf.Level)

	return this.Level, err

}

//写日志信息
//level 日志等级
//msg 	日志信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *ConseleLog) WriteLog(level int, msg string) (err error) {
	if this.Level > level {
		return nil
	}

	this.logger.Println(msg)

	return nil

}

//关闭日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *ConseleLog) CloseLog() {

}
