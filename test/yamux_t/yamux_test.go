package yamux_t

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/hashicorp/yamux"
)

func TestClient(t *testing.T) {
	connOri, err := net.Dial("tcp", "127.0.0.1:9002")
	if err != nil {
		t.Fatal(err)
	}
	defer connOri.Close()
	session, err := yamux.Client(connOri, yamux.DefaultConfig())
	for {
		conn, err := session.Open()
		if err != nil {
			t.Fatal(err)
		}
		conn.Write([]byte("hello world"))
		time.Sleep(time.Second)
	}
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:9002")
	if err != nil {
		t.Fatal(err)
	}
	defer lis.Close()
	conn, err := lis.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	session, err := yamux.Server(conn, yamux.DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err = session.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(conn.RemoteAddr().String(), string(b[:n]))
		}()
	}
}
