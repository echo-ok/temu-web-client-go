// Temu 浏览器客户端
package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
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
	ProductService     productService
}

type Client struct {
	Debug                bool
	Logger               resty.Logger
	Services             services
	TimeLocation         *time.Location
	BaseUrl              string
	SellerCentralBaseUrl string
	SellerCentralClient  *resty.Client // SellerCentral专用客户端
	BgClient             *resty.Client // BgAuth专用客户端
	MallId               int
}

func New(config config.TemuBrowserConfig) *Client {
	logger := config.Logger

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
					logger.Errorf("重新获取 Anti-Content 失败: %v", err)
					return false
				}
				response.Request.SetHeader("Anti-Content", antiContent)

				logger.Debugf("重试请求，URL: %s", response.Request.URL)
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
		ProductService:     productService{xService, client},
	}

	client.BgClient = httpClient

	sellerCentralClient := resty.New().
		SetLogger(logger).
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
					logger.Errorf("重新获取 Anti-Content 失败: %v", err)
					return false
				}
				response.Request.SetHeader("Anti-Content", antiContent)

				logger.Debugf("重试请求，URL: %s", response.Request.URL)
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
		// 对于非2xx响应，手动解析错误信息
		var errorResult normal.Response
		if err := json.Unmarshal(resp.Body(), &errorResult); err != nil {
			return fmt.Errorf("failed to parse error response: %v", err)
		}

		if errorResult.ErrorMessage != "" {
			return errors.New(errorResult.ErrorMessage)
		}

		var errorResult2 normal.Response2
		if err := json.Unmarshal(resp.Body(), &errorResult2); err != nil {
			return fmt.Errorf("failed to parse error response: %v", err)
		}

		if errorResult2.ErrorMessage != "" {
			return errors.New(errorResult2.ErrorMessage)
		}

		return errors.New("unknown error")
	}

	if !result.Success {
		if result.ErrorCode == entity.ErrorNeedSMSCode {
			return normal.ErrNeedSMSCode
		}

		if result.ErrorCode == entity.ErrorNeedVerifyCode {
			return normal.ErrNeedVerifyCode
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

func (c *Client) SetMallId(mallId int) {
	c.MallId = mallId
}

func (c *Client) CheckMallId() error {
	if c.MallId == 0 {
		return errors.New("mall ID is not set")
	}
	return nil
}

func (c *Client) SetCookie(cookies []*http.Cookie, clearOld bool) {
	if clearOld {
		c.SellerCentralClient.Cookies = []*http.Cookie{}
	}

	c.SellerCentralClient.SetCookies(cookies)
}

func (c *Client) GetCookie() []*http.Cookie {
	return c.SellerCentralClient.Cookies
}

func (c *Client) SetAccountCookie(cookies []*http.Cookie, clearOld bool) {
	if clearOld {
		c.BgClient.Cookies = []*http.Cookie{}
	}

	c.BgClient.SetCookies(cookies)
}

func (c *Client) GetAccountCookie() []*http.Cookie {
	return c.BgClient.Cookies
}

func (c *Client) Clone() *Client {
	newClient := &Client{
		Debug:                c.Debug,
		Logger:               c.Logger,
		TimeLocation:         c.TimeLocation,
		BaseUrl:              c.BaseUrl,
		SellerCentralBaseUrl: c.SellerCentralBaseUrl,
		MallId:               c.MallId,
	}

	// 克隆 http client
	newClient.SellerCentralClient = c.SellerCentralClient.Clone()
	newClient.BgClient = c.BgClient.Clone()

	// 初始化服务
	xService := service{
		debug:      c.Debug,
		logger:     c.Logger,
		httpClient: newClient.BgClient,
	}

	// 重新初始化服务层
	newClient.Services = services{
		RecentOrderService: recentOrderService{xService, newClient},
		BgAuthService:      bgAuthService{xService, newClient},
		StockService:       stockService{xService, newClient},
		ProductService:     productService{xService, newClient},
	}

	return newClient
}

func (c *Client) IsBgSessionInvalid() bool {
	_, err := c.Services.BgAuthService.GetMallInfoByKuangjianmaihuo(context.Background())
	return err != nil
}

func (c *Client) IsSellerCentralSessionInvalid() bool {
	_, err := c.Services.BgAuthService.GetUserInfo(context.Background())
	return err != nil
}
