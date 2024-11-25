package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Nekika/protohackers/pkg/boilerplate"
)

func TestE2E(t *testing.T) {
	testCases := []struct {
		Description string
		Message     []byte
	}{
		{
			Description: "Normal usage",
			Message:     []byte("Hello, World!"),
		},
	}

	const port = ":6543"
	const host = "localhost" + port

	go boilerplate.ListenTCP(port, handler)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {

			conn := boilerplate.MustDialTCP(host)

			n, err := conn.Write(tc.Message)
			assert.Nil(t, err)
			assert.Equal(t, len(tc.Message), n)

			err = conn.CloseWrite()
			assert.Nil(t, err)

			message, err := io.ReadAll(conn)
			assert.Nil(t, err)
			assert.Equal(t, tc.Message, message)
		})
	}
}
