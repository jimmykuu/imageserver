package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var config Config

func init() {
	// 读取配置文件
	file, err := os.Open(strings.Join([]string{"conf", "settings.json"}, string(os.PathSeparator)))
	if err != nil {
		log.Fatal("没有找到配置文件 conf/settings.json ", err)
		panic(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&config)
	if err != nil {
		log.Fatal("conf/settings.json 解析失败", err)
	}

	// 检查目录是否存在
	if _, err := os.Stat(config.ImagePath); os.IsNotExist(err) {
		log.Fatal("图片目录：", config.ImagePath, "不存在")
	}
}
