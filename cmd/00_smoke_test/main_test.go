package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Nekika/protohackers/pkg/boilerplate"
)

func TestMain(m *testing.M) {
	go run()
	m.Run()
}

func TestE2E(t *testing.T) {
	testCases := []struct {
		Description      string
		Message          []byte
		ExpectedResponse []byte
	}{
		{
			Description:      "Normal usage 1",
			Message:          []byte("Hello, World!"),
			ExpectedResponse: []byte("Hello, World!"),
		},
		{
			Description:      "Normal usage 2",
			Message:          []byte("ABCDEFG12345"),
			ExpectedResponse: []byte("ABCDEFG12345"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			const host = "localhost" + port

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
