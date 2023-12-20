package main

import (
	"go-server/core"
	"go-server/global"
)

func main() {
	// 初始化配置文件
	core.InitialConfig()
	// 初始化logrus
	global.Log = core.InitLogger()
	global.Log.Warn("1234")
	global.Log.Debug("1234")
	global.Log.Error("1234")
	global.Log.Info("1234")
	// 连接数据库
	core.Gorm()
}
