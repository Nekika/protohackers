package main

import (
	"encoding/binary"
	"errors"
	"net"
	"time"

	"github.com/Nekika/protohackers/pkg/boilerplate"
)

const port = ":6543"

const (
	messageKindInsert byte = 'I'
	messageKindQuery  byte = 'Q'
)

const messageLength = 9

func handler(conn *net.TCPConn) error {
	var (
		h = new(Historical)

		b   = make([]byte, messageLength)
		err error
	)

	for {
		if _, err = conn.Read(b); err != nil {
			break
		}

		kind, data := b[0], b[1:]

		switch kind {
		case messageKindInsert:
			handleInsert(conn, h, data)
		case messageKindQuery:
			handleQuery(conn, h, data)
		default:
			err = errors.New("unknown command")
		}

		if err != nil {
			break
		}
	}

	return err
}

func handleInsert(_ *net.TCPConn, h *Historical, data []byte) {
	entry := Entry{
		Price:     parsePrice(data[4:]),
		Timestamp: parseTimestamp(data[:4]),
	}
	h.Insert(entry)
}

func handleQuery(conn *net.TCPConn, h *Historical, data []byte) error {
	mintime := parseTimestamp(data[:4])
	maxtime := parseTimestamp(data[4:])
	price := h.Query(mintime, maxtime)
	encoded := encodePrice(price)
	_, err := conn.Write(encoded)
	return err
}

func encodePrice(p int32) []byte {
	return binary.BigEndian.AppendUint32([]byte{}, uint32(p))
}

func parsePrice(b []byte) int32 {
	p := binary.BigEndian.Uint32(b)
	return int32(p)
}

func parseTimestamp(b []byte) *time.Time {
	seconds := binary.BigEndian.Uint32(b)
	timestamp := time.Unix(int64(seconds), 0)
	return &timestamp
}

func run() {
	boilerplate.ListenTCP(port, handler)
}

func main() {
	run()
}
