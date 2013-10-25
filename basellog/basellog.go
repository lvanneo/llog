package basellog

//Llog适配器
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com

import (
	"fmt"
	"github.com/lvanneo/llog/config"
	"sync"
)

//Llog日志结构体
type Llogger struct {
	llogadapters map[string]LlogInterface //可用的适配器
	lowestlevel  int                      //所有可用适配器中的最低日志等级
	msgchannel   chan *LlogMSG            //与写日志的协程传递日志信息的通道
	lock         sync.Mutex               //锁
}

//Llog日志信息结构体
type LlogMSG struct {
	level int    //日志等级
	msg   string //日志信息
}

//Llog日志适配器初始化
//channelLength 传送日志信息的 channel 大小
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *Llogger) InitLlogger(channelLength int64) {
	this.llogadapters = make(map[string]LlogInterface)
	this.msgchannel = make(chan *LlogMSG, channelLength)
	this.lowestlevel = LevelALL

	go this.runLlog()

}

//通过配置信息配置Llog日志
//configinfo Llog日志配置信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *Llogger) SetLloggers(configinfo []byte) error {
	adps, err := config.GetAppender(configinfo)
	if err != nil {
		return err
	}

	for _, adp := range adps {
		this.SetLlogger(adp, configinfo)
	}

	return nil

}

//指定Llog适配器，通过配置信息配置Llog日志
//configinfo Llog日志配置信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *Llogger) SetLlogger(adaptername string, configinfo []byte) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	aptname, err := CheckAdapter(adaptername)
	if err != nil {
		return err
	}

	adapter, ok := registeredAdapters[aptname]
	if ok {
		log := adapter()
		lelel, err := log.InitLog(configinfo)
		if err != nil {
			return err
		}

		if this.lowestlevel > lelel {
			this.lowestlevel = lelel
		}

		this.llogadapters[aptname] = log

		return nil

	} else {
		return fmt.Errorf("未注册的日志适配器：%q", adaptername)
	}

	/*
		adapter, ok := this.llogadapters[aptname]
		if ok {
			level, err := adapter.InitLog(configinfo)
			if err != nil {
				return err
			}
			if this.lowestlevel > level {
				this.lowestlevel = level
			}
		} else {
			switch aptname {
			case llog.AdapterConsole:
				adapter = impl.NewConseLog()
				adapter.InitLog(configinfo)
				this.llogadapters[llog.AdapterConsole] = adapter

			}
		}

		return nil
	*/
	return nil
}

//写日志
//将日志信息格式化，并通过 channel 将格式化后的日志信息传送给另一个协程进行写日志
//level	待写的日志等级
//val 	不定参的日志信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *Llogger) writeLog(level int, val ...interface{}) error {
	if this.lowestlevel > level {
		return nil
	}

	msg := fmt.Sprint(val...)

	lev, err := ChangeLevel(level)
	if err != nil {
		return err
	}
	msg = fmt.Sprintln(lev, msg)

	llogmsg := new(LlogMSG)
	llogmsg.level = level
	llogmsg.msg = msg
	this.msgchannel <- llogmsg

	return nil

}

//格式化写日志
//将日志信息按指定格式进行格式化，并通过 channel 将格式化后的日志信息传送给另一个协程进行写日志
//level	 	待写的日志等级
//format	格式化信息
//val 		不定参的日志信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *Llogger) writeLogf(level int, format string, val ...interface{}) error {
	if this.lowestlevel > level {
		return nil
	}

	msg := fmt.Sprintf(format, val...)

	lev, err := ChangeLevel(level)
	if err != nil {
		return err
	}
	msg = fmt.Sprintln(lev, msg)

	llogmsg := new(LlogMSG)
	llogmsg.level = level
	llogmsg.msg = msg
	this.msgchannel <- llogmsg

	return nil
}

//运行日志
//从 channel 中获取日志信息，并调用已注册的可用的日志适配器进行写日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *Llogger) runLlog() {
	for {
		select {
		case llogmsg := <-this.msgchannel:
			for _, logsty := range this.llogadapters {
				logsty.WriteLog(llogmsg.level, llogmsg.msg)
			}
		}
	}

}

func (this *Llogger) Debug(val ...interface{}) {
	this.writeLog(LevelDEBUG, val...)
}

func (this *Llogger) Info(val ...interface{}) {
	this.writeLog(LevelINFO, val...)
}

func (this *Llogger) Warn(val ...interface{}) {
	this.writeLog(LevelWARN, val...)
}

func (this *Llogger) Error(val ...interface{}) {
	this.writeLog(LevelERROR, val...)
}

func (this *Llogger) Fatal(val ...interface{}) {
	this.writeLog(LevelFATAL, val...)
}

func (this *Llogger) Debugf(format string, val ...interface{}) {
	this.writeLogf(LevelDEBUG, format, val...)
}

func (this *Llogger) Infof(format string, val ...interface{}) {
	this.writeLogf(LevelINFO, format, val...)
}

func (this *Llogger) Warnf(format string, val ...interface{}) {
	this.writeLogf(LevelWARN, format, val...)
}

func (this *Llogger) Errorf(format string, val ...interface{}) {
	this.writeLogf(LevelERROR, format, val...)
}

func (this *Llogger) Fatalf(format string, val ...interface{}) {
	this.writeLogf(LevelFATAL, format, val...)
}
