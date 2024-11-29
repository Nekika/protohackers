package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Nekika/protohackers/pkg/boilerplate"
)

func TestMain(m *testing.M) {
	go run()
	m.Run()
}

const (
	SessionMessageEmitterClient = 0
	SessionMessageEmitterServer = 1
)

type SessionMessage struct {
	Data    []byte
	Emitter int
}

func TestE2E(t *testing.T) {
	testCases := []struct {
		Description string
		Session     []SessionMessage
	}{
		{
			Description: "Normal session",
			Session: []SessionMessage{
				{Data: []byte{0x49, 0x00, 0x00, 0x30, 0x39, 0x00, 0x00, 0x00, 0x65}, Emitter: SessionMessageEmitterClient},
				{Data: []byte{0x49, 0x00, 0x00, 0x30, 0x3a, 0x00, 0x00, 0x00, 0x66}, Emitter: SessionMessageEmitterClient},
				{Data: []byte{0x49, 0x00, 0x00, 0x30, 0x3b, 0x00, 0x00, 0x00, 0x64}, Emitter: SessionMessageEmitterClient},
				{Data: []byte{0x49, 0x00, 0x00, 0xa0, 0x00, 0x00, 0x00, 0x00, 0x05}, Emitter: SessionMessageEmitterClient},
				{Data: []byte{0x51, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x40, 0x00}, Emitter: SessionMessageEmitterClient},
				{Data: []byte{0x00, 0x00, 0x00, 0x65}, Emitter: SessionMessageEmitterServer},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			const host = "localhost" + port

			conn := boilerplate.MustDialTCP(host)
			defer conn.Close()

			for _, msg := range tc.Session {
				if msg.Emitter == SessionMessageEmitterClient {
					n, err := conn.Write(msg.Data)
					assert.Nil(t, err)
					assert.Equal(t, len(msg.Data), n)
				} else {
					data := make([]byte, 4)
					n, err := conn.Read(data)
					assert.Nil(t, err)
					assert.Equal(t, len(data), n)
					assert.Equal(t, msg.Data, data)
				}
			}
		})
	}
}
