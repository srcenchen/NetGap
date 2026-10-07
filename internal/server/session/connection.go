package session

import (
	"NetGap/internal/protocol"
	"NetGap/internal/server/session/mux"
	"context"
	"net"

	"github.com/sirupsen/logrus"
)

type Session struct {
	clientId    string
	tcpMux      *mux.TcpMux
	controlConn net.Conn
	RelayMap    map[string]*RelaySession
}
type verityClient func(clientId string, token string) bool

// AcceptHandshake 握手判定
func AcceptHandshake(conn net.Conn, auth verityClient) (*Session, error) {
	jsonCodec := protocol.NewJSONCodec()
	hs := protocol.HandshakeReq{}
	jsonCodec.Decode(conn, &hs)
	authStatus := auth(hs.ClientID, hs.Token)
	err := handshakeResp(conn, authStatus)
	if err != nil {
		return nil, err
	}
	if !authStatus {
		logrus.Errorf("客户端身份校验失败 %s, ClientId %s", conn.RemoteAddr().String(), hs.ClientID)
		return nil, ErrAuthFailed
	}
	// 身份校验通过 构建 Session 并升级多路复用连接
	logrus.Infof("客户端连接成功 %s, ClientId %s", conn.RemoteAddr().String(), hs.ClientID)
	s := &Session{
		clientId: hs.ClientID,
		tcpMux:   mux.UpgradeMuxClient(conn),
		RelayMap: make(map[string]*RelaySession),
	}
	s.controlConn, err = s.CreateVirtualConn()
	return s, err
}

// RunNewRelay 启用新的中转会话
func (s *Session) RunNewRelay(tunnelId string, exposePort, targetAddr string) error {
	vConn, err := s.CreateVirtualConn()
	if err != nil {
		return err
	}
	rs := NewReplySession(tunnelId, vConn)
	s.RelayMap[tunnelId] = rs
	err = rs.Listen(context.Background(), exposePort)
	return err
}

func (s *Session) CreateVirtualConn() (net.Conn, error) {
	conn, err := s.tcpMux.Session.Open()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// CloseChan 存活检测 CloseChan 透传
func (s *Session) CloseChan() <-chan struct{} {
	return s.tcpMux.Session.CloseChan()
}

// handshakeResp 将会返回握手结果给客户端
// conn 是客户端的连接，status 是握手结果
func handshakeResp(conn net.Conn, status bool) error {
	jc := protocol.NewJSONCodec()
	if !status {
		return jc.Encode(conn, &protocol.HandshakeResp{
			Code: 1,
			Msg:  "客户端身份信息校验失败",
		})
	}
	return jc.Encode(conn, &protocol.HandshakeResp{
		Code: 0,
		Msg:  "success",
	})
}

func (s *Session) GetClientId() string {
	return s.clientId
}
