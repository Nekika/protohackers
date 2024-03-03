package lib

const (
	EventKindClientJoined = iota + 1
	EventKindClientLeft
)

type Event struct {
	Kind   int
	Client *Client
}
