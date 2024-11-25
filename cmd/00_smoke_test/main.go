package main

import (
	"io"
	"net"

	"github.com/Nekika/protohackers/pkg/boilerplate"
)

func handler(conn *net.TCPConn) error {
	var (
		data []byte
		err  error
	)

	if data, err = io.ReadAll(conn); err == nil {
		_, err = conn.Write(data)
	}

	return err
}

func main() {
	boilerplate.ListenTCP(":6543", handler)
}
