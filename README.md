Llog
====

Llog，一个简单实用的Golang日志模块。  <br>
可直接使用。如果存在Llog的配置文件在conf、config及当前目录的任何目录下即可自动加载该配置文件。也可通过代码加载指定路径的配置文件。

## 示例
```
package main

import (
	"bufio"
	"github.com/lvanneo/llog/llogger"
	"os"
)

func main() {
	llogger.Debug("Debug: Llog 示例")
	llogger.Info("Info: Llog 示例")
	llogger.Warn("Warn: Llog 示例")
	llogger.Error("Error: Llog 示例")
	llogger.Fatal("Fatal: Llog 示例")

	llogger.Debugf("Debug: %s %s ", "Llog", "示例")
	llogger.Infof("Info: %s %s ", "Llog", "示例")
	llogger.Warnf("Error: %s %s ", "Llog", "示例")
	llogger.Errorf("Error: %s %s ", "Llog", "示例")
	llogger.Fatalf("Fatal: %s %s ", "Llog", "示例")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()
}
```
## Llog实现功能
1.日志分为7个级别，如下：
	ALL
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
2.支持输出日志到控制台及文件。
3.每个级别支持两种写日志函数：字符串默认拼接和自定义格式拼接。
4.在等于或高于指定日志输出等级时才进行字符串拼接。
5.支持日志文件名称及大小设定。
6.支持控制台与文件日志输出等级的分离设定。
7.支持xml、json格式配置。
8.日志输出文件可关闭，默认打开日志输出控制台。

## Llog日志实现方式
主要分为三层：
上层：封装启动和使用日志的相关工具方法。
中层：封装日志处理核心。
底层：写日志的具体执行。如：控制台、文件

底层与中层分离，采用适配器的方式，每实现一个新的日志输出形式，只需主动向中层日志核心提交注册即可使用该实现。所以很便于扩展其他日志输出实现。