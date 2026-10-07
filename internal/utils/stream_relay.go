package utils

import (
	"io"
	"net"
	"sync"
)

func RelayStream(sourceConn net.Conn, targetConn net.Conn) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	// target -> source
	go func() {
		defer wg.Done()
		_, err := io.Copy(sourceConn, targetConn)
		if err != nil {
			if tcpConn, ok := sourceConn.(*net.TCPConn); ok {
				_ = tcpConn.CloseWrite()
			}
		}
	}()
	// source -> target
	go func() {
		defer wg.Done()
		_, err := io.Copy(targetConn, sourceConn)
		if err != nil {
			if tcpConn, ok := targetConn.(*net.TCPConn); ok {
				_ = tcpConn.CloseWrite()
			}
		}
	}()
	wg.Wait()
}
