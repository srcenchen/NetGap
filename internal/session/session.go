package session

import (
	"NetGap/internal/mux"
	"NetGap/internal/protocol"
	"net"

	"github.com/sirupsen/logrus"
)

type Session struct {
	ClientId string
	TcpMux   *mux.TcpMux
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
		ClientId: hs.ClientID,
		TcpMux:   mux.UpgradeMuxClient(conn),
	}
	return s, nil
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
