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
