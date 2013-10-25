package config

//Llog日志配置信息工具集
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com

import (
	"encoding/xml"
)

type XmlConfig struct {
	XMLName     xml.Name      `xml:"Config"`
	ConsoleInfo ConsoleConfig `xml:"Console"`
	FileInfo    FileConfig    `xml:"File"`
	Appender    []string      `xml:"Appender"`
}

type ConsoleConfig struct {
	Level     string `xml:"Level"`
	ShortFile bool   `xml:"ShortFile"`
}

type FileConfig struct {
	Level     string `xml:"Level"`
	FileName  string `xml:"FileName"`
	FileSize  int    `xml:"FileSize"`
	MaxLine   int    `xml:"MaxLine"`
	ShortFile bool   `xml:"ShortFile"`
}

//获取配置信息中的Llog适配器名称
//configInfo	Llog配置信息
//返回Llog适配器名称集合及错误信息
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com
func GetAppender(configInfo []byte) ([]string, error) {
	conf := new(XmlConfig)
	err := xml.Unmarshal(configInfo, &conf)
	if err != nil {
		return conf.Appender, err
	}

	return conf.Appender, nil

}

//获取配置信息中Llog控制台日志的配置信息
//configInfo	Llog配置信息
//返回Llog控制台日志配置信息及错误信息
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com
func GetConsoleConfig(configInfo []byte) (ConsoleConfig, error) {
	conf := new(XmlConfig)
	err := xml.Unmarshal(configInfo, &conf)
	if err != nil {
		return conf.ConsoleInfo, err
	}

	return conf.ConsoleInfo, nil

}

//获取配置信息中Llog文件日志的配置信息
//configInfo	Llog配置信息
//返回Llog文件日志配置信息及错误信息
//
//2013-10-25
//李林(LvanNeo)
//lvan_software@foxmail.com
func GetFileConfig(configInfo []byte) (FileConfig, error) {
	conf := new(XmlConfig)
	err := xml.Unmarshal(configInfo, &conf)
	if err != nil {
		return conf.FileInfo, err
	}

	return conf.FileInfo, nil
}
