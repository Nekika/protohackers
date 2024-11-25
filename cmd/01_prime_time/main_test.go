package main

import (
	"bufio"
	"testing"

	"github.com/Nekika/protohackers/pkg/boilerplate"
	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	testCases := []struct {
		Description      string
		Message          []byte
		ExpectedResponse []byte
	}{
		{
			Description:      "Conforming request  Non prime number  123",
			Message:          []byte("{\"method\":\"isPrime\",\"number\":123}\n"),
			ExpectedResponse: []byte("{\"method\":\"isPrime\",\"prime\":false}\n"),
		},
		{
			Description:      "Conforming request  Non prime number  36",
			Message:          []byte("{\"method\":\"isPrime\",\"number\":36}\n"),
			ExpectedResponse: []byte("{\"method\":\"isPrime\",\"prime\":false}\n"),
		},
		{
			Description:      "Conforming request  Prime number  11",
			Message:          []byte("{\"method\":\"isPrime\",\"number\":11}\n"),
			ExpectedResponse: []byte("{\"method\":\"isPrime\",\"prime\":true}\n"),
		},
		{
			Description:      "Conforming request  Prime number  101",
			Message:          []byte("{\"method\":\"isPrime\",\"number\":11}\n"),
			ExpectedResponse: []byte("{\"method\":\"isPrime\",\"prime\":true}\n"),
		},
		{
			Description:      "Malformed request  Non JSON payload",
			Message:          []byte("isPrime 112\n"),
			ExpectedResponse: []byte("{\"error\":\"malformed request\"}\n"),
		},
		{
			Description:      "Malformed request  Invalid method",
			Message:          []byte("{\"method\":\"checkIfPrime\",\"number\":11}\n"),
			ExpectedResponse: []byte("{\"error\":\"malformed request\"}\n"),
		},
		{
			Description:      "Malformed request  Invalid number",
			Message:          []byte("{\"method\":\"isPrime\",\"number\":\"twelve\"}\n"),
			ExpectedResponse: []byte("{\"error\":\"malformed request\"}\n"),
		},
	}

	const port = ":6543"
	const host = "localhost" + port

	go func() {
		h := boilerplate.TCPListener{Handler: handler, ErrorHandler: errorHandler}
		h.Listen(":6543")
	}()

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			conn := boilerplate.MustDialTCP(host)

			n, err := conn.Write(tc.Message)
			assert.Nil(t, err)
			assert.Equal(t, len(tc.Message), n)

			reader := bufio.NewReader(conn)

			data, err := reader.ReadBytes('\n')
			assert.Nil(t, err)
			assert.Equal(t, tc.ExpectedResponse, data)
		})
	}
}
