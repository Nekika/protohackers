package main

import (
	"io"
	"log"
	"net"
)

func serve() {
	lister, err := net.Listen("tcp", ":43240")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lister.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer conn.Close()

			if _, err := io.Copy(conn, conn); err != nil {
				log.Println("err copying:", err)
			}
		}()
	}
}

func main() {
	serve()
}
