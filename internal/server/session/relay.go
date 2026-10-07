package session

import (
	"NetGap/internal/utils"
	"context"
	"net"
)

type RelaySession struct {
	id         string
	targetConn net.Conn
}

func NewReplySession(id string, conn net.Conn) *RelaySession {
	return &RelaySession{
		id:         id,
		targetConn: conn,
	}
}

func (s *RelaySession) Id() string {
	return s.id
}

func (s *RelaySession) Close() {
	_ = s.targetConn.Close()
}

func (s *RelaySession) Listen(ctx context.Context, exposePort string) error {
	ln, err := net.Listen("tcp", ":"+exposePort)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = ln.Close()
	}()
	for {
		sourceConn, err := ln.Accept()
		if err != nil {
			return err
		}
		s.relay(sourceConn)
	}
}

// relay 中转数据流
func (s *RelaySession) relay(source net.Conn) {
	utils.RelayStream(source, s.targetConn)
}
