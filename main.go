package main

import (
	"fmt"
	"go-server/core"
	"go-server/global"
)

func main() {
	// 初始化配置文件
	core.InitialConfig()
	fmt.Println(global.Config)
	// 连接数据库
	core.Gorm()

	fmt.Println(global.DB)
}
