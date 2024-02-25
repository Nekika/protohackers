package main

import (
	"net"
	"testing"
)

func TestMain(m *testing.M) {
	go serve()
	m.Run()
}

type Request struct {
	ToSend    string
	ToReceive string
}

type TestCase struct {
	Description string
	Requests    []Request
	ShouldAbort bool
}

func TestSolution(t *testing.T) {
	cases := []TestCase{
		{
			Description: "Single conforming request",
			Requests: []Request{
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":7}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":true}\n",
				},
			},
			ShouldAbort: false,
		},
		{
			Description: "Multiple conforming request",
			Requests: []Request{
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":7}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":true}\n",
				},
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":12}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":false}\n",
				},
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":33}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":false}\n",
				},
			},
			ShouldAbort: false,
		},
		{
			Description: "Single malformed request from method",
			Requests: []Request{
				{
					ToSend:    "{\"method\":\"UNKNOWN\",\"number\":7}\n",
					ToReceive: "{\"error\":\"invalid request\"}\n",
				},
			},
			ShouldAbort: false,
		},
		{
			Description: "Single malformed request from number",
			Requests: []Request{
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":\"sixteen\"}\n",
					ToReceive: "{\"error\":\"invalid request\"}\n",
				},
			},
			ShouldAbort: false,
		},
		{
			Description: "Multiple requests with first malformed",
			Requests: []Request{
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":\"sixteen\"}\n",
					ToReceive: "{\"error\":\"invalid request\"}\n",
				},
				{
					ToSend: "{\"method\":\"isPrime\",\"number\":7}\n",
				},
				{
					ToSend: "{\"method\":\"isPrime\",\"number\":12}\n",
				},
			},
			ShouldAbort: true,
		},
		{
			Description: "Multiple requests with last malformed",
			Requests: []Request{
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":7}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":true}\n",
				},
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":12}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":false}\n",
				},
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":\"sixteen\"}\n",
					ToReceive: "{\"error\":\"invalid request\"}\n",
				},
			},
			ShouldAbort: false,
		},
		{
			Description: "Multiple requests with one malformed",
			Requests: []Request{
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":7}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":true}\n",
				},
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":12}\n",
					ToReceive: "{\"method\":\"isPrime\",\"prime\":false}\n",
				},
				{
					ToSend:    "{\"method\":\"isPrime\",\"number\":\"sixteen\"}\n",
					ToReceive: "{\"error\":\"invalid request\"}\n",
				},
				{
					ToSend: "{\"method\":\"isPrime\",\"number\":36}\n",
				},
			},
			ShouldAbort: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			conn, err := net.Dial("tcp", ":43240")
			if err != nil {
				t.Fatal(err)
			}
			defer conn.Close()

			var count int

			for _, req := range c.Requests {
				var (
					err error

					read    int
					reqdata = []byte(req.ToSend)
					resdata = make([]byte, len(req.ToReceive))
				)

				if _, err = conn.Write(reqdata); err != nil {
					break
				}

				if read, err = conn.Read(resdata); err != nil || read == 0 {
					break
				}

				if string(resdata) != req.ToReceive {
					t.Fatalf("responses don't match: expected %s but got %s", req.ToReceive, string(resdata))
				}

				count += 1
			}

			if !c.ShouldAbort && count != len(c.Requests) {
				t.Fatalf("received less responses than expected: expected %d but got %d", len(c.Requests), count)
			}

			if c.ShouldAbort && (count == len(c.Requests)) {
				t.Fatal("expected the connection to be aborted sooner by the server")
			}
		})
	}
}
