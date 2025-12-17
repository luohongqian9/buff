package core

import (
	"gopkg.in/yaml.v3"
	"log"
	"server/config"
	"server/utils"
)

func InitConfig() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYaml()
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	if err := yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return c
}
