package config

import (
	"log"
	"time"
)

type TemuBrowserConfig struct {
	Debug                bool          `json:"debug"`                                                                                                                                      // 是否为调试模式
	BaseUrl              string        `json:"base_url"`                                                                                                                                   // 基础 URL
	SellerCentralBaseUrl string        `json:"seller_central_base_url"`                                                                                                                    // 卖家中心基础 URL
	Timeout              time.Duration `json:"timeout"`                                                                                                                                    // 超时时间（秒）
	VerifySSL            bool          `json:"verify_ssl"`                                                                                                                                 // 是否验证 SSL
	UserAgent            string        `json:"user_agent" default:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"` // User Agent
	Proxy                string        `json:"proxy"`                                                                                                                                      // 代理
	Logger               *log.Logger   `json:"-"`                                                                                                                                          // 日志
}
