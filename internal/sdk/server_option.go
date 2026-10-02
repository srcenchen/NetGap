package sdk

type ServerOptions struct {
	TunnelAddr string // 隧道监听地址，客户端连入
}

type ServerOption func(*ServerOptions)

func WithServerTunnelAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.TunnelAddr = addr
	}
}
