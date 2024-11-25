package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"math/big"
	"net"

	"github.com/Nekika/protohackers/pkg/boilerplate"
)

const port = ":6543"

var (
	errMalformedRequest = errors.New("malformed request")
)

type request struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

func (r request) valid() bool {
	return r.Method == "isPrime"
}

func isPrime(number int) bool {
	return big.NewInt(int64(number)).ProbablyPrime(0)
}

func parse(data []byte) (int, error) {
	var (
		req request
		err error
	)

	if err = json.Unmarshal(data, &req); err != nil || !req.valid() {
		err = errMalformedRequest
	}

	return req.Number, err
}

func handler(conn *net.TCPConn) error {
	var (
		reader = bufio.NewReader(conn)
		err    error
	)

	for {
		var (
			data   []byte
			number int
		)

		if data, err = reader.ReadBytes('\n'); err != nil {
			break
		}

		if number, err = parse(data); err != nil {
			break
		}

		if _, err = boilerplate.WriteStringf(conn, "{\"method\":\"isPrime\",\"prime\":%v}\n", isPrime(number)); err != nil {
			break
		}
	}

	return err
}

func errorHandler(err error, conn *net.TCPConn) {
	var message string

	if errors.Is(err, errMalformedRequest) {
		message = err.Error()
	}

	if message != "" {
		boilerplate.WriteStringf(conn, "{\"error\":\"%s\"}\n", message)
	}
}

func run() {
	h := boilerplate.TCPListener{Handler: handler, ErrorHandler: errorHandler}
	h.Listen(port)
}

func main() {
	run()
}
