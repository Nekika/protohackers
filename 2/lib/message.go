package lib

import "encoding/binary"

const (
	MessageKindInsert rune = 'I'
	MessageKindQuery  rune = 'Q'
)

type Message struct {
	kind  byte
	left  []byte
	right []byte
}

func ParseMessage(encoded []byte) Message {
	return Message{
		kind:  encoded[0],
		left:  encoded[1:5],
		right: encoded[5:9],
	}
}

func (m Message) Kind() rune {
	return rune(m.kind)
}

func (m Message) Left() int32 {
	return int32(binary.BigEndian.Uint32(m.left))
}

func (m Message) Right() int32 {
	return int32(binary.BigEndian.Uint32(m.right))
}
