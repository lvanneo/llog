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
	"strconv"
	"sync"
)

//文件日志结构体
type FileLog struct {
	logger    *log.Logger
	mwfile    *MutexWriter
	level     int
	filename  string
	shortFile bool
	maxsize   int64 //文件最大值
	nowsize   int64 //文件当前大小
	check     sync.Mutex
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
	flog.level = LevelALL
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
	this.maxsize = conf.FileSize * 1024 * 1024

	/*
		this.shortFile = conf.ShortFile
		if this.shortFile {
			this.logger = log.New(this.mwfile, "", log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			this.logger = log.New(this.mwfile, "", log.Ldate|log.Ltime)
		}
	*/

	this.level, err = CheckLevel(conf.Level)

	err = this.initFile()

	return this.level, err

}

//写日志信息
//level 日志等级
//msg 	日志信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) WriteLog(level int, msg string) (err error) {
	if this.level > level {
		return nil
	}

	msgsize := int64(len(msg))
	this.nowsize += msgsize

	this.checkLogFile()

	this.logger.Println(msg)

	if 0 == this.nowsize {
		this.nowsize += msgsize
	}

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

	stat, err := file.Stat()
	if err != nil {
		this.nowsize = 0
	} else {
		this.nowsize = stat.Size()
	}

	this.mwfile.SetFile(file)

	return nil
}

//日志文件检查
//
//2013-10-28
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) checkLogFile() {
	if this.nowsize < this.maxsize {
		return
	}

	this.check.Lock()

	this.CloseLog()

	for i := 0; ; i++ {
		newname := this.filename + strconv.Itoa(i)
		err := os.Rename(this.filename, newname)
		if err != nil {
			continue
		} else {
			break
		}
	}

	this.initFile()

	this.check.Unlock()

}

//关闭日志
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func (this *FileLog) CloseLog() {
	this.mwfile.file.Close()
}
