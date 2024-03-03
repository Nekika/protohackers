package main

import (
	"log"
	"net"

	"github.com/nekika/protohackers/3/lib"
)

func main() {
	listener, err := net.Listen("tcp", ":42789")
	if err != nil {
		panic(err)
	}

	log.Fatal(lib.Serve(listener))
}
