package sdk

import "time"

type ClientOptions struct {
	ServerAddr string        // 服务端地址
	Timeout    time.Duration // 超时时间
	
}
