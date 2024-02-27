package lib

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	MessageKindInsert = rune('I')
	MessageKindQuery  = rune('Q')
)

type Message struct {
	Kind  rune
	Left  int32
	Right int32
}

func ParseMessage(raw []byte) (Message, error) {
	var (
		message Message
		err     error
	)

	if message.Kind, err = parseMember(raw[:1]); err != nil {
		return Message{}, err
	}

	if message.Kind != MessageKindInsert && message.Kind != MessageKindQuery {
		return Message{}, errors.New("invalid kind")
	}

	if message.Left, err = parseMember(raw[1:5]); err != nil {
		return Message{}, err
	}

	if message.Right, err = parseMember(raw[5:9]); err != nil {
		return Message{}, err
	}

	return message, nil
}

func parseMember(member []byte) (int32, error) {
	var (
		binary string
		parsed int64

		err error
	)

	for _, b := range member {
		binary += fmt.Sprintf("%02d", b)
	}

	if parsed, err = strconv.ParseInt(binary, 16, 32); err != nil {
		return 0, err
	}

	return int32(parsed), nil
}
