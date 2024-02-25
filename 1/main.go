package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

var (
	ErrorResponse = []byte("{\"error\":\"invalid request\"}\n")
)

func handle(conn net.Conn) {
	defer func() {
		log.Println("closing")
		conn.Close()
	}()

	var (
		err     error
		reqdata []byte
		resdata []byte
		req     IsPrimeRequest
		res     IsPrimeResponse

		scanner = bufio.NewScanner(conn)
	)

	var n = 1
	for scanner.Scan() {
		log.Println("handling request number", n)

		reqdata = scanner.Bytes()
		if err != nil {
			break
		}

		if err = json.Unmarshal(reqdata, &req); err != nil || !req.Valid() {
			_, _ = conn.Write(ErrorResponse)
			break
		}

		res = NewIsPrimeResponse(req.Number)
		if resdata, err = json.Marshal(res); err != nil {
			break
		}
		resdata = append(resdata, '\n')

		if _, err = conn.Write(resdata); err != nil {
			break
		}

		n++
	}
}

func serve() {
	listener, err := net.Listen("tcp", ":43240")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting conn:", err)
			continue
		}

		go handle(conn)
	}
}

func main() {
	serve()
}
