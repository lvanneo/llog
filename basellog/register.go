package basellog

import (
	"fmt"
)

//定义返回方法类型 LlogType ，方法返回值为 LlogInterface 接口类型
type LlogType func() LlogInterface

//声明适配器的Map集合，已存放注册的适配器
var registeredAdapters = make(map[string]LlogType)

//适配器注册
//将已实现的具体的日志适配器进行注册
//adaptername	注册的适配器名称
//llog			待注册的适配器
//
//2013-10-24
//李林(LvanNeo)
//lvan_software@foxmail.com
func Register(adaptername string, llog LlogType) {
	if llog == nil {
		fmt.Println("日志注册失败！")
		return
	}

	_, adp := registeredAdapters[adaptername]
	if adp {
		fmt.Printf("日志注册失败，系统已存在日志：%q \n", adaptername)
		return
	} else {
		registeredAdapters[adaptername] = llog
	}

}
