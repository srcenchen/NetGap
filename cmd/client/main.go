package main

import (
	"NetGap/internal/protocol"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6500")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	handshake := protocol.Handshake{
		ClientID: "dab1955a-0163-4302-95ed-0131dcf27ec5",
		Token:    "5a4019ff-fa08-4c33-b47b-427234fea421",
	}
	jsonCodec := protocol.NewJSONCodec()
	jsonCodec.Encode(conn, handshake)
}
