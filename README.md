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
