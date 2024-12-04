// Temu 浏览器客户端
package client

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/bestk/temu-helper/config"
	"github.com/bestk/temu-helper/entity"
	"github.com/bestk/temu-helper/normal"
	"github.com/bestk/temu-helper/utils"
	"github.com/go-resty/resty/v2"
)

type service struct {
	debug      bool          // Is debug mode
	logger     resty.Logger  // Log
	httpClient *resty.Client // HTTP client
}

type services struct {
	RecentOrderService recentOrderService
	BgAuthService      bgAuthService
	StockService       stockService
}

type Client struct {
	Debug                bool
	Logger               resty.Logger
	Services             services
	TimeLocation         *time.Location
	BaseUrl              string
	SellerCentralBaseUrl string
	SellerCentralClient  *resty.Client // SellerCentral专用客户端
	MallId               uint64
}

// 添加自定义 Logger 结构体
type customLogger struct {
	*log.Logger
}

// 实现 resty.Logger 接口所需的方法
func (l *customLogger) Errorf(format string, v ...interface{}) { l.Printf("ERROR "+format, v...) }
func (l *customLogger) Warnf(format string, v ...interface{})  { l.Printf("WARN "+format, v...) }
func (l *customLogger) Debugf(format string, v ...interface{}) { l.Printf("DEBUG "+format, v...) }

func createLogger() *customLogger {
	return &customLogger{log.New(os.Stdout, "[ Temu ] ", log.LstdFlags|log.Llongfile)}
}

func New(config config.TemuBrowserConfig) *Client {
	logger := config.Logger
	if logger == nil {
		logger = createLogger()
	}
	client := &Client{
		Debug:                config.Debug,
		Logger:               logger,
		BaseUrl:              config.BaseUrl,
		SellerCentralBaseUrl: config.SellerCentralBaseUrl,
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logger.Errorf("load location error: %v", err)
	}
	client.TimeLocation = loc

	httpClient := resty.New().
		SetLogger(logger).
		SetDebug(config.Debug).
		EnableTrace().
		SetBaseURL(config.BaseUrl).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).
		SetHeader("User-Agent", config.UserAgent).
		SetAllowGetMethodPayload(true).
		SetTimeout(config.Timeout * time.Second).
		SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.VerifySSL},
			DialContext: (&net.Dialer{
				Timeout: config.Timeout * time.Second,
			}).DialContext,
		}).
		SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			if req.Response.StatusCode == 302 {
				return nil
			}
			return http.ErrUseLastResponse
		})).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			values := make(map[string]any)
			if request.Body != nil {
				b, e := json.Marshal(request.Body)
				if e != nil {
					return e
				}

				e = json.Unmarshal(b, &values)
				if e != nil {
					return e
				}
			}
			// 设置请求头中的Anti-Content
			antiContent, err := utils.GetAntiContent()
			if err != nil {
				return err
			}
			request.SetHeader("Anti-Content", antiContent)
			return nil
		}).
		SetRetryCount(3).
		SetRetryWaitTime(time.Duration(500) * time.Millisecond).
		SetRetryMaxWaitTime(time.Duration(1) * time.Second).
		AddRetryCondition(func(response *resty.Response, err error) bool {
			if response == nil {
				return false
			}

			retry := response.StatusCode() == http.StatusTooManyRequests
			if !retry {
				r := struct {
					Success   bool   `json:"success"`
					ErrorCode int    `json:"errorCode"`
					ErrorMsg  string `json:"errorMsg"`
				}{}
				retry = json.Unmarshal(response.Body(), &r) == nil &&
					!r.Success &&
					r.ErrorCode == 4000000 &&
					strings.EqualFold(r.ErrorMsg, "SYSTEM_EXCEPTION")
			}
			if retry {
				// 重新设置 Anti-Content
				antiContent, err := utils.GetAntiContent()
				if err != nil {
					logger.Printf("重新获取 Anti-Content 失败: %v", err)
					return false
				}
				response.Request.SetHeader("Anti-Content", antiContent)

				logger.Printf("重试请求，URL: %s", response.Request.URL)
			}
			return retry
		})
	if config.Proxy != "" {
		httpClient.SetProxy(config.Proxy)
	}
	httpClient.JSONMarshal = json.Marshal
	httpClient.JSONUnmarshal = json.Unmarshal
	xService := service{
		debug:      config.Debug,
		logger:     logger,
		httpClient: httpClient,
	}
	client.Services = services{
		RecentOrderService: recentOrderService{xService, client},
		BgAuthService:      bgAuthService{xService, client},
		StockService:       stockService{xService, client},
	}

	sellerCentralClient := resty.New().
		SetDebug(config.Debug).
		EnableTrace().
		SetBaseURL(config.SellerCentralBaseUrl).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).
		SetHeader("User-Agent", config.UserAgent).
		SetAllowGetMethodPayload(true).
		SetTimeout(config.Timeout * time.Second).
		SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.VerifySSL},
			DialContext: (&net.Dialer{
				Timeout: config.Timeout * time.Second,
			}).DialContext,
		}).
		SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			if req.Response.StatusCode == 302 {
				return nil
			}
			return http.ErrUseLastResponse
		})).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			values := make(map[string]any)
			if request.Body != nil {
				b, e := json.Marshal(request.Body)
				if e != nil {
					return e
				}

				e = json.Unmarshal(b, &values)
				if e != nil {
					return e
				}
			}
			// 设置请求头中的Anti-Content
			antiContent, err := utils.GetAntiContent()
			if err != nil {
				return err
			}
			request.SetHeader("Anti-Content", antiContent)
			return nil
		}).
		SetRetryCount(3).
		SetRetryWaitTime(time.Duration(500) * time.Millisecond).
		SetRetryMaxWaitTime(time.Duration(1) * time.Second).
		AddRetryCondition(func(response *resty.Response, err error) bool {
			if response == nil {
				return false
			}

			retry := response.StatusCode() == http.StatusTooManyRequests
			if !retry {
				r := struct {
					Success   bool   `json:"success"`
					ErrorCode int    `json:"errorCode"`
					ErrorMsg  string `json:"errorMsg"`
				}{}
				retry = json.Unmarshal(response.Body(), &r) == nil &&
					!r.Success &&
					r.ErrorCode == 4000000 &&
					strings.EqualFold(r.ErrorMsg, "SYSTEM_EXCEPTION")
			}
			if retry {
				// 重新设置 Anti-Content
				antiContent, err := utils.GetAntiContent()
				if err != nil {
					logger.Printf("重新获取 Anti-Content 失败: %v", err)
					return false
				}
				response.Request.SetHeader("Anti-Content", antiContent)

				logger.Printf("重试请求，URL: %s", response.Request.URL)
			}
			return retry
		})
	if config.Proxy != "" {
		sellerCentralClient.SetProxy(config.Proxy)
	}
	sellerCentralClient.JSONMarshal = json.Marshal
	sellerCentralClient.JSONUnmarshal = json.Unmarshal
	xService.httpClient = sellerCentralClient
	client.SellerCentralClient = sellerCentralClient

	return client
}

func recheckError(resp *resty.Response, result normal.Response, e error) (err error) {
	if e != nil {
		return e
	}

	if resp.IsError() {
		errorMessage := strings.TrimSpace(result.ErrorMessage)

		return errors.New(errorMessage)
	}

	if !result.Success {
		if result.ErrorCode == entity.ErrorNeedSMSCode {
			return normal.ErrNeedSMSCode
		}
		return errors.New(result.ErrorMessage)
	}
	return nil
}

func parseResponseTotal(currentPage, pageSize, total int) (n, totalPages int, isLastPage bool) {
	if currentPage == 0 {
		currentPage = 1
	}

	totalPages = (total / pageSize) + 1
	return total, totalPages, currentPage >= totalPages
}

func (c *Client) SetMallId(mallId uint64) {
	c.MallId = mallId
}

func (c *Client) CheckMallId() error {
	if c.MallId == 0 {
		return errors.New("mall ID is not set")
	}
	return nil
}

func (c *Client) SetCookie(cookies []*http.Cookie) {
	c.SellerCentralClient.SetCookies(cookies)
}

func (c *Client) GetCookie() []*http.Cookie {
	// 输出所有cookie
	url, err := url.Parse(c.SellerCentralBaseUrl)
	if err != nil {
		c.Logger.Printf("解析 SellerCentralBaseURL失败: %v", err)
		return nil
	}
	cookies := c.SellerCentralClient.GetClient().Jar.Cookies(url)
	c.Logger.Printf("cookies: %v", cookies)
	return cookies
}
