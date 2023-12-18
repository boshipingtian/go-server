package core

import (
	"fmt"
	"go-server/config"
	"go-server/global"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// InitialConfig 读取yaml文件。初始化配置
func InitialConfig() {
	const configPath = "settings.yaml"
	conf := &config.Config{}
	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("get yaml config err %s", err))
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		panic(fmt.Errorf("config Unmarshal err %v", err))
	}
	log.Println("config yaml file load init success.")
	global.Config = conf
}
