package mux

import (
	"net"

	"github.com/hashicorp/yamux"
)

type TcpMux struct {
	Session *yamux.Session
}

// UpgradeMuxServer 将普通TCP连接升级为多路复用连接
func UpgradeMuxServer(conn net.Conn) *TcpMux {
	session, _ := yamux.Server(conn, yamux.DefaultConfig())
	return &TcpMux{
		Session: session,
	}
}

// UpgradeMuxClient 客户端中 将TCP连接升级为多路复用连接
func UpgradeMuxClient(conn net.Conn) *TcpMux {
	session, _ := yamux.Client(conn, yamux.DefaultConfig())
	return &TcpMux{
		Session: session,
	}
}
