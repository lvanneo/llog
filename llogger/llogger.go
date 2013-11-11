package llogger

//Llog日志工具集
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com

import (
	"fmt"
	"github.com/lvanneo/llog/basellog"
	"os"
)

//Llog 版本信息
const (
	Version = "Llog 1.0"
	Author  = "李林(LvanNeo)"
	Eamil   = "lvan_software@foxmail.com"
	VDate   = "2013-10-25"
)

//日志适配器接口
var Llogger basellog.AdapterInterface

//Llog日志初始化
//初始化Llog日志系统，并调用配置方法对Llog进行初始化配置
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func init() {
	Llogger = new(basellog.Llogger)
	Llogger.InitLlogger(10000)

	err := initload()
	if err != nil {
		err = defaultConfig()
		if err != nil {
			fmt.Println("Llog 日志缺省加载失败，Llog日志无法启动，请检查。。。\n", err)
			return
		} else {
			Info("Llog 日志已成功加载缺省配置")
		}
	} else {
		Info("Llog 日志已成功加载默认配置")
	}

	Info("Llog 日志已启动")
}

//Llog日志初始化加载配置信息
//通过设置默认的配置文件信息路径来加载配置信息
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com
func initload() error {
	paths := [6]string{"./llog.xml", "./conf/llog.xml", "./config/llog.xml", "./llog.json", "./conf/llog.json", "./config/llog.json"}

	var err error
	for _, ph := range paths {
		err = LoadConfig(ph)
		if err == nil {
			return nil
		}
	}
	return err
}

//配置文件缺省时初始化 Llog 日志
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com
func defaultConfig() error {
	config := `
<?xml version="1.0" encoding="UTF-8"?>
<Config>
	<Console>
		<Level>ALL</Level>
	</Console>
</Config>`

	return Llogger.SetLlogger("console", []byte(config))
}

//通过配置信息配置 Llog 日志
//可通过配置信息中的配置初始化多个日志适配器并进行配置
//config 配置信息，配置信息中的适配器名称必须为已注册的适配器名称
//返回配置的错误信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func SetLloggers(config []byte) error {
	err := Llogger.SetLloggers(config)
	if err != nil {
		Errorf("配置错误：%s", err)
		return fmt.Errorf("配置错误：%s", err)
	}

	return nil
}

//通过配置信息配置 Llog 日志
//通过指定适配器名称初始化日志适配器并进行配置
//adaptername 	适配器名称，必须为已注册的适配器名称
//config 		配置信息
//返回配置的错误信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func SetLlogger(adaptername string, config []byte) error {
	err := Llogger.SetLlogger(adaptername, config)
	if err != nil {
		Errorf("%q 配置错误：%s", adaptername, err)
		return fmt.Errorf("%q 配置错误：%s", adaptername, err)
	}

	Infof("Llog 日志 %q 适配器已成功加载配置", adaptername)

	return nil
}

//打开配置文件并读取配置信息
//configPath 	配置文件路径
//返回文件操作错误信息
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com
func LoadConfig(configPath string) error {
	configfile, err := os.OpenFile(configPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer configfile.Close()

	var conf []byte
	buf := make([]byte, 1024*10)
	i := 1
	for {
		n, _ := configfile.Read(buf)

		if 0 == n {
			break
		}

		conf = append(conf, buf[0:n]...)
		i++
	}

	SetLloggers(conf)
	return nil
}

//打印 Llog 版本信息
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com
func GetVersion() {
	fmt.Printf("\n版本：%s \n作者：%s \n邮箱：%s\n日期：%s\n\n", Version, Author, Eamil, VDate)
}

func Debug(val ...interface{}) {
	Llogger.Debug(val...)
}

func Info(val ...interface{}) {
	Llogger.Info(val...)
}

func Warn(val ...interface{}) {
	Llogger.Warn(val...)
}

func Error(val ...interface{}) {
	Llogger.Error(val...)
}

func Fatal(val ...interface{}) {
	Llogger.Fatal(val...)
}

func Debugf(format string, val ...interface{}) {
	Llogger.Debugf(format, val...)
}

func Infof(format string, val ...interface{}) {
	Llogger.Infof(format, val...)
}

func Warnf(format string, val ...interface{}) {
	Llogger.Warnf(format, val...)
}

func Errorf(format string, val ...interface{}) {
	Llogger.Errorf(format, val...)
}

func Fatalf(format string, val ...interface{}) {
	Llogger.Fatalf(format, val...)
}
