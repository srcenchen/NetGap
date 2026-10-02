package sdk

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Client struct {
	ClientOptions
}
type ClientOptions struct {
	ServerAddr string        // 服务端地址
	Timeout    time.Duration // 超时时间
	ClientId   string        // 客户端 Id
	Token      string        // 校验密钥
	CryptoType CryptoType    // 加密类型
}

type ClientOption func(*ClientOptions)

func WithServerAddr(addr string) ClientOption {
	return func(o *ClientOptions) {
		o.ServerAddr = addr
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *ClientOptions) {
		o.Timeout = timeout
	}
}

func WithToken(token string) ClientOption {
	return func(o *ClientOptions) {
		o.Token = token
	}
}

func WithCryptoType(ct CryptoType) ClientOption {
	return func(o *ClientOptions) {
		o.CryptoType = ct
	}
}

// NewClient 穿透客户端
// 警告：必须传入 ServerAddr 和 Token，否则无法启动
func NewClient(opts ...ClientOption) (*Client, error) {
	o := &ClientOptions{
		ServerAddr: "",
		Timeout:    0,
		Token:      "",
		CryptoType: CryptoNone,
	}
	for _, opt := range opts {
		opt(o)
	}
	client := &Client{
		ClientOptions: *o,
	}
	if o.ServerAddr == "" || o.Token == "" {
		return nil, errors.New("ServerAddr 或者 Token 为空")
	}
	if o.CryptoType == CryptoNone {
		logrus.Warn("[安全警告]: 当加密协议不配置时，您的数据将可能处于被监听状态")
	}
	return client, nil
}
