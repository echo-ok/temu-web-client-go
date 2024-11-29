package main

import (
	"github.com/bestk/temu-helper/browser"
	config "github.com/bestk/temu-helper/config"
)

func main() {
	cfg := config.TemuBrowserConfig{
		// 设置配置项
	}
	client := browser.New(cfg)
	// ... 使用 client
}
