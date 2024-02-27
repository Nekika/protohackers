package main

import (
	"encoding/binary"
	"github.com/nekika/protohackers/2/lib"
	"net"
)

const Address = ":42567"

func handle(conn net.Conn) {
	defer conn.Close()

	var (
		buff       = make([]byte, 9)
		err        error
		message    lib.Message
		repository = lib.NewRepository()
	)

	for {
		if _, err = conn.Read(buff); err != nil {
			break
		}

		message = lib.ParseMessage(buff)

		switch message.Kind() {
		case lib.MessageKindInsert:
			handleInsert(conn, repository, lib.InsertFromMessage(message))
		case lib.MessageKindQuery:
			handleQuery(conn, repository, lib.QueryFromMessage(message))
		}
	}
}

func handleInsert(_ net.Conn, repository *lib.Repository, insert lib.Insert) {
	repository.Insert(insert.Amount, insert.Timestamp)
}

func handleQuery(conn net.Conn, repository *lib.Repository, query lib.Query) {
	avg := repository.Average(query.MinTime, query.MaxTime)

	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, uint32(avg))

	_, _ = conn.Write(buff)
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
