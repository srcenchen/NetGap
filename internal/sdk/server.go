package sdk

import (
	"NetGap/internal/server/biz"
	"NetGap/internal/server/data"
	"NetGap/internal/server/handler"
	session2 "NetGap/internal/server/session"
	"context"
	"errors"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ServerOptions
	sessionManager *session2.Manager
}

func NewServer(opts ...ServerOption) (*Server, error) {
	o := &ServerOptions{
		TunnelAddr: "",
	}
	for _, opt := range opts {
		opt(o)
	}
	sm := session2.NewSessionManager()
	server := &Server{
		ServerOptions:  *o,
		sessionManager: sm,
	}
	if o.TunnelAddr == "" {
		return nil, errors.New("tunnel 监听地址为空")
	}
	if o.HttpAddr == "" {
		return nil, errors.New("http-api 监听地址为空")
	}
	// 初始化 数据库
	if err := data.InitServerDb("serverData.db"); err != nil {
		return nil, err
	}
	return server, nil
}

func (s *Server) Run(ctx context.Context) error {
	// 启动 Tunnel 隧道服务 server-client 段
	lc := net.ListenConfig{}
	tunnelLn, err := lc.Listen(ctx, "tcp", s.TunnelAddr)
	if err != nil {
		return err
	}
	logrus.Infof("NetGap Tunnel 服务运行在 %s", tunnelLn.Addr())
	defer func() {
		_ = tunnelLn.Close()
		logrus.Info("NetGap 服务已经停止")
	}()
	go s.tunnelAccept(tunnelLn)

	// 启动 GIN-HTTP 服务器
	//gin.SetMode(gin.ReleaseMode)
	ginSrv := gin.New()
	ginSrv.Use(gin.Recovery())
	authorized := ginSrv.Group("/")
	r := handler.NewRouter(s.sessionManager)
	r.AuthNeededRouter(authorized) // 路由注册
	go ginSrv.Run(s.HttpAddr)
	// LOOP
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Server) tunnelAccept(tunnelLn net.Listener) {
	for {
		conn, err := tunnelLn.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			logrus.Errorf("[TunnelErr]: %v", err)
			continue
		}
		go func() {
			sess, err := session2.AcceptHandshake(conn, biz.VerityClient)
			if err != nil {
				logrus.Errorf("客户端握手失败 %s", err.Error())
			}
			s.sessionManager.SetSession(sess.GetClientId(), sess) // 纳入到 sessionManager 管理
		}()
	}
}
