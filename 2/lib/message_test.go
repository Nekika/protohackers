package lib

import "testing"

func TestMessage_Kind(t *testing.T) {
	var (
		message  = Message{kind: 0x49}
		expected = 'I'
	)

	if message.Kind() != expected {
		t.Fatalf("kinds don't match: expected %v but got %v", expected, message.Kind())
	}
}

func TestMessage_Left(t *testing.T) {
	var (
		message        = Message{left: []byte{0x00, 0x00, 0x30, 0x39}}
		expected int32 = 12345
	)

	if message.Left() != expected {
		t.Fatalf("lefts don't match: expected %v but got %v", expected, message.Left())
	}
}

func TestMessage_Right(t *testing.T) {
	var (
		message        = Message{right: []byte{0x00, 0x00, 0x00, 0x65}}
		expected int32 = 101
	)

	if message.Right() != expected {
		t.Fatalf("rights don't match: expected %v but got %v", expected, message.Right())
	}
}
