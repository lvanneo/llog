package basellog

import (
	"errors"
)

//日志等级常量
const (
	LevelALL = iota
	LevelDEBUG
	LevelINFO
	LevelWARN
	LevelERROR
	LevelFATAL
	LevelOFF
)

//可注册适配器常量
const (
	AdapterConsole = "CONSOLE"
	AdapterFile    = "FILE"
)

//日志适配器检查
//adaptername 适配器名称
//返回对应的日志适配器常量名称，如果非指定适配器名称则返回其原名称及错误信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func CheckAdapter(adaptername string) (string, error) {
	switch adaptername {
	case AdapterConsole:
		fallthrough
	case "console":
		fallthrough
	case "Console":
		return AdapterConsole, nil
	case AdapterFile:
		fallthrough
	case "file":
		fallthrough
	case "File":
		return AdapterFile, nil
	default:
		return adaptername, errors.New("非系统指定适配器！")
	}
}

//日志等级转换成对应的 [D] 形式
//level 日志等级
//返回转换后对应的字符串，如何不存在则返回空字符串及错误信息
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func ChangeLevel(level int) (string, error) {
	switch level {
	case LevelALL:
		return "[A]", nil
	case LevelDEBUG:
		return "[D]", nil
	case LevelINFO:
		return "[I]", nil
	case LevelWARN:
		return "[W]", nil
	case LevelERROR:
		return "[E]", nil
	case LevelFATAL:
		return "[F]", nil
	case LevelOFF:
		return "[O]", nil
	default:
		return "", errors.New("日志等级不存在！")
	}
}

//日志等级检查及转换成系统定义的常量等级
//level 日志等级，为字符串表示形式
//返回系统定义的常量日志等级
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func CheckLevel(level string) (int, error) {
	switch level {
	case "ALL":
		fallthrough
	case "All":
		fallthrough
	case "all":
		fallthrough
	case "0":
		return LevelALL, nil
	case "DEBUG":
		fallthrough
	case "Debug":
		fallthrough
	case "debug":
		fallthrough
	case "1":
		return LevelDEBUG, nil
	case "INFO":
		fallthrough
	case "Info":
		fallthrough
	case "info":
		fallthrough
	case "2":
		return LevelINFO, nil
	case "WARN":
		fallthrough
	case "Warn":
		fallthrough
	case "warn":
		fallthrough
	case "3":
		return LevelWARN, nil
	case "ERROR":
		fallthrough
	case "Error":
		fallthrough
	case "error":
		fallthrough
	case "4":
		return LevelERROR, nil
	case "FATAL":
		fallthrough
	case "Fatal":
		fallthrough
	case "fatal":
		fallthrough
	case "5":
		return LevelFATAL, nil
	case "OFF":
		fallthrough
	case "Off":
		fallthrough
	case "off":
		fallthrough
	case "6":
		return LevelOFF, nil
	default:
		return LevelALL, errors.New("日志等级无效！将采用最低等级！")
	}

}
