package session

import (
	"NetGap/internal/protocol"
	"net"

	"github.com/hashicorp/yamux"
	"github.com/sirupsen/logrus"
)

type Session struct {
	tcpMux yamux.Session
}
type verityClient func(clientId string, token string) bool

// AcceptHandshake 握手判定
func AcceptHandshake(conn net.Conn, auth verityClient) (*Session, error) {
	jsonCodec := protocol.NewJSONCodec()
	hs := protocol.Handshake{}
	jsonCodec.Decode(conn, &hs)
	authStatus := auth(hs.ClientID, hs.Token)
	logrus.Infof("auth status: %v", authStatus)
	return nil, nil
}
