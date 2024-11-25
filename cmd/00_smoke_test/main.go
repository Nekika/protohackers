package main

import (
	"io"
	"net"

	"github.com/Nekika/protohackers/pkg/boilerplate"
)

const port = ":6543"

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

func run() {
	boilerplate.ListenTCP(port, handler)
}

func main() {
	run()
}
