package session

import (
	"net"

	"github.com/hashicorp/yamux"
)

type Session struct {
	tcpMux yamux.Session
}
type Authenticator interface {
	verityClient(clientId string, token string) bool
}

func (s *Session) Accept(conn *net.Conn, auth Authenticator) error {

}
