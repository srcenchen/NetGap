package client

import (
	"NetGap/internal/mux"
	"NetGap/internal/protocol"
	"io"
	"net"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:6500")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	handshake := protocol.HandshakeReq{
		ClientID: "dab1955a-0163-4302-95ed-0131dcf27ec5",
		Token:    "5a4019ff-fa08-4c33-b47b-427234fea421",
	}
	jsonCodec := protocol.NewJSONCodec()
	jsonCodec.Encode(conn, handshake)
	hsResp := protocol.HandshakeResp{}
	jsonCodec.Decode(conn, &hsResp)
	t.Log(hsResp.Msg)
	t.Log("升级连接到yamux")
	m := mux.UpgradeMuxServer(conn)

	for {
		vconn, err := m.Session.Accept()
		if err != nil {
			logrus.Error(err)
		}
		// 做连接转发
		// 后续第一个包先确定目标连接设备
		go func() {
			// 现场连接 targetConn
			targetConn, err := net.Dial("tcp", "127.0.0.1:3000")
			if err != nil {
				logrus.Error(err)
				_ = vconn.Close()
				return
			}
			// 做io copy
			defer targetConn.Close()
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				io.Copy(targetConn, vconn)
				halfClose(targetConn)
			}()
			go func() {
				defer wg.Done()
				io.Copy(vconn, targetConn)
				halfClose(vconn)
			}()
			wg.Wait()
		}()
	}
}

func halfClose(c net.Conn) {
	if tcpConn, ok := c.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
	} else {
		c.Close()
	}
}
