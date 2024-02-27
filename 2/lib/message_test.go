package lib

import (
	"reflect"
	"testing"
)

func TestParseMessage(t *testing.T) {
	testCases := []struct {
		Description string
		Encoded     []byte
		Message     Message
	}{
		{
			Description: "Normal usage",
			Encoded:     []byte{49, 00, 00, 30, 39, 00, 00, 00, 65},
			Message: Message{
				Kind:  MessageKindInsert,
				Left:  12345,
				Right: 101,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			message, err := ParseMessage(tc.Encoded)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(message, tc.Message) {
				t.Fatalf("mesages don't match: expected %#v but got %#v", tc.Message, message)
			}
		})
	}
}
