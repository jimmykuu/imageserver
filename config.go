package main

type Config struct {
	Addr      string `json:"addr"`       // 图片服务器地址
	ImagePath string `json:"image_path"` // 图片路径
}
