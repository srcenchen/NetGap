package sdk

import (
	"NetGap/internal/biz"
	"NetGap/internal/server/data"
	"NetGap/internal/session"
	"context"
	"errors"
	"net"

	"github.com/sirupsen/logrus"
)

type Server struct {
	ServerOptions
	sessionManager *SessionManager
}

func NewServer(opts ...ServerOption) (*Server, error) {
	o := &ServerOptions{
		TunnelAddr: "",
	}
	for _, opt := range opts {
		opt(o)
	}
	sm := NewSessionManager()
	server := &Server{
		ServerOptions:  *o,
		sessionManager: sm,
	}
	if o.TunnelAddr == "" {
		return nil, errors.New("tunnel 监听地址为空")
	}
	// 初始化 数据库
	if err := data.InitServerDb("serverData.db"); err != nil {
		return nil, err
	}
	return server, nil
}

func (s *Server) Run(ctx context.Context) error {
	// 监听 tcp 端口
	lc := net.ListenConfig{}
	ln, err := lc.Listen(ctx, "tcp", s.TunnelAddr)
	if err != nil {
		return err
	}
	logrus.Infof("NetGap Tunnel 服务运行在 %s", ln.Addr())
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
			sess, err := session.AcceptHandshake(conn, biz.VerityClient)
			if err != nil {
				logrus.Errorf("客户端握手失败 %s", err.Error())
			}
			s.sessionManager.SetSession(sess.ClientId, sess)
		}()
	}
}
