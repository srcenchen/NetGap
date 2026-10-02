package sdk

import (
	"bufio"
	"context"
	"errors"
	"net"
	"strings"

	"github.com/sirupsen/logrus"
)

type Server struct {
	ServerOptions
}
type ServerOptions struct {
	TunnelAddr string // 隧道监听地址，客户端连入
}

type ServerOption func(*ServerOptions)

func WithServerTunnelAddr(addr string) ServerOption {
	return func(o *ServerOptions) {
		o.TunnelAddr = addr
	}
}

func NewServer(opts ...ServerOption) (*Server, error) {
	o := &ServerOptions{
		TunnelAddr: "",
	}
	for _, opt := range opts {
		opt(o)
	}
	server := &Server{
		ServerOptions: *o,
	}
	if o.TunnelAddr == "" {
		return nil, errors.New("tunnel 监听地址为空")
	}
	return server, nil
}

func (s *Server) Run(ctx context.Context) error {
	// 监听 tcp 端口
	lc := net.ListenConfig{}
	ln, err := lc.Listen(ctx, "tcp", s.TunnelAddr)
	logrus.Infof("NetGap 服务运行在 %s", ln.Addr())
	if err != nil {
		return err
	}
	defer func() {
		_ = ln.Close()
		logrus.Info("NetGap 服务已经停止")
	}()
	go func() {
		<-ctx.Done()
		_ = ln.Close()
	}()
	for {
		conn, err := ln.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			continue
		}
		go func() {
			r := bufio.NewReader(conn)
			for {
				line, err := r.ReadString('\n')
				if err != nil {
					return
				}
				logrus.Info(strings.TrimSpace(line))
			}
		}()
		_ = conn
	}
}
