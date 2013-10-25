package basellog

//文件日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com

import (
	"github.com/lvanneo/llog/config"
	"log"
	"os"
	"sync"
)

//文件日志结构体
type FileLog struct {
	logger    *log.Logger
	mwfile    *MutexWriter
	Level     int
	filename  string
	ShortFile bool
}

//文件操作结构体
type MutexWriter struct {
	sync.Mutex
	file *os.File
}

//文件日志初始化
//注册文件日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func init() {
	Register(AdapterFile, NewFileLog) //注册文件日志
}

//创建文件日志
//返回已创建的文件日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func NewFileLog() LlogInterface {
	flog := new(FileLog)
	flog.Level = LevelALL
	flog.filename = ""
	flog.mwfile = new(MutexWriter)
	flog.logger = log.New(flog.mwfile, "", log.Ldate|log.Ltime)
	return flog
}

//文件操作结构体写文件方法
//msg 日志信息
//返回写的字节数及错误信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *MutexWriter) Write(msg []byte) (int, error) {
	this.Lock()
	defer this.Unlock()

	return this.file.Write(msg)
}

//文件操作结构体设置文件对象方法
//file 文件对象
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *MutexWriter) SetFile(file *os.File) {
	if this.file != nil {
		this.file.Close()
	}

	this.file = file
}

//根据配置信息初始化文件日志
//通过配置信息初始化文件日志
//configInfo 	配置信息
//level 		返回配置信息的日志等级及错误信息，错误则返回最低等级及错误信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) InitLog(configInfo []byte) (level int, err error) {

	conf, err := config.GetFileConfig(configInfo)
	if err != nil {
		return LevelALL, err
	}

	this.filename = conf.FileName

	this.ShortFile = conf.ShortFile
	if this.ShortFile {
		this.logger = log.New(this.mwfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		this.logger = log.New(this.mwfile, "", log.Ldate|log.Ltime)
	}

	this.Level, err = CheckLevel(conf.Level)

	err = this.initFile()

	return this.Level, err

}

//写日志信息
//level 日志等级
//msg 	日志信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) WriteLog(level int, msg string) (err error) {
	if this.Level > level {
		return nil
	}

	this.logger.Println(msg)

	return nil

}

//创建日志文件
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) createLlogFile() (*os.File, error) {
	return os.OpenFile(this.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
}

//日志文件初始化
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) initFile() error {
	file, err := this.createLlogFile()
	if err != nil {
		return err
	}

	this.mwfile.SetFile(file)

	return nil
}

//关闭日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) CloseLog() {

}
