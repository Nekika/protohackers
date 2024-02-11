package main

import (
	"net"
	"testing"
)

func TestMain(m *testing.M) {
	go serve()
	m.Run()
}

func TestSolution(t *testing.T) {
	cases := []struct {
		Message string
	}{
		{Message: "Hello\n"},
		{Message: "Holà\n"},
		{Message: "Ciao\n"},
		{Message: "Bonjour\n"},
		{Message: "Priviet\n"},
	}

	for _, c := range cases {
		t.Run(c.Message, func(t *testing.T) {
			conn, err := net.Dial("tcp", ":43240")
			if err != nil {
				t.Fatal(err)
			}
			defer conn.Close()

			message := []byte(c.Message)
			if _, err := conn.Write(message); err != nil {
				t.Fatal("err writing to server", err)
			}

			buff := make([]byte, len(message))
			if _, err := conn.Read(buff); err != nil {
				t.Fatal("err reading from server:", err)
			}

			if c.Message != string(buff) {
				t.Fatalf("messages don't match: expected %#v but got %#v", c.Message, string(buff))
			}
		})

	}
}
