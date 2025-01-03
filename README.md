# TEMU Web Client

## 项目简介

这是一个用 Go 语言实现的 TEMU 网页客户端,用于与 TEMU API 进行交互。支持卖家中心的各项功能操作。

## 主要特性

-   完整的 API 封装
-   自动处理认证和会话管理
-   内置请求重试和错误处理
-   支持代理配置
-   详细的调试日志

## 技术栈

-   Go 1.19+
-   resty (HTTP 客户端)
-   goja (JavaScript 运行时)

## 配置选项

| 配置项               | 说明              | 默认值     |
| -------------------- | ----------------- | ---------- |
| Debug                | 是否启用调试模式  | false      |
| BaseUrl              | TEMU API 基础地址 | -          |
| SellerCentralBaseUrl | 卖家中心 API 地址 | -          |
| Timeout              | 请求超时时间      | 30s        |
| VerifySSL            | 是否验证 SSL 证书 | true       |
| UserAgent            | 自定义 User-Agent | Chrome 120 |
| Proxy                | 代理服务器地址    | -          |

## 主要功能

-   订单管理
-   商品管理
-   库存管理
-   物流跟踪
-   数据报表
-   账户设置

## 错误处理

客户端会自动处理常见错误:

-   网络超时自动重试
-   Token 过期自动刷新
-   请求限流自动等待
-   详细的错误信息返回

## 开发说明

### 目录结构

```
├── config/         # 配置相关
├── utils/          # 工具函数
├── models/         # 数据模型
└── services/       # API 服务
```

### 运行测试

```bash
go test ./...
```

## 注意事项

1. 请合理控制请求频率
2. 建议在生产环境使用代理
3. 定期检查 API 更新

## 贡献指南

欢迎提交 Pull Request 或 Issue。

## 许可证

MIT License

## 免责声明

本项目仅用于正常的商业用途,请遵守 TEMU 的使用条款和政策。
