package main

import (
	"net"
	"slices"
	"testing"
)

func TestMain(m *testing.M) {
	go serve()
	m.Run()
}

func TestSolution(t *testing.T) {
	type Communication struct {
		ToSend    []byte
		ToReceive []byte
	}

	testCases := []struct {
		Description string
		Session     []Communication
	}{
		{
			Description: "Normal session",
			Session: []Communication{
				{
					ToSend:    []byte{0x49, 0x00, 0x00, 0x30, 0x39, 0x00, 0x00, 0x00, 0x65},
					ToReceive: nil,
				},
				{
					ToSend:    []byte{0x49, 0x00, 0x00, 0x30, 0x3a, 0x00, 0x00, 0x00, 0x66},
					ToReceive: nil,
				},
				{
					ToSend:    []byte{0x49, 0x00, 0x00, 0x30, 0x3b, 0x00, 0x00, 0x00, 0x64},
					ToReceive: nil,
				},
				{
					ToSend:    []byte{0x49, 0x00, 0x00, 0xa0, 0x00, 0x00, 0x00, 0x00, 0x05},
					ToReceive: nil,
				},
				{
					ToSend:    []byte{0x51, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x40, 0x00},
					ToReceive: []byte{0x00, 0x00, 0x00, 0x65},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			conn, err := net.Dial("tcp", Address)
			if err != nil {
				t.Fatal(err)
			}
			defer conn.Close()

			for _, communication := range tc.Session {
				if _, err := conn.Write(communication.ToSend); err != nil {
					t.Fatal(err)
				}

				if communication.ToReceive == nil {
					continue
				}

				buff := make([]byte, 4)
				t.Log("reading from conn")
				if _, err := conn.Read(buff); err != nil {
					t.Fatal(err)
				}

				if !slices.Equal(buff, communication.ToReceive) {
					t.Fatalf("responses don't match: expected %v but got %v", communication.ToReceive, buff)
				}
			}
		})
	}
}
