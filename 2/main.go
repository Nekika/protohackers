package main

import (
	"net"
)

const Address = "42567"

func handle(conn net.Conn) {
	defer conn.Close()
	panic("not implemented")
}

func serve() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handle(conn)
	}
}

func main() {
	serve()
}
