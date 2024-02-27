package lib

type Insert struct {
	Price     int32
	Timestamp int32
}

func InsertFromMessage(message Message) Insert {
	return Insert{
		Price:     message.Left,
		Timestamp: message.Right,
	}
}

type Query struct {
	MaxTime int32
	MinTime int32
}

func QueryFromMessage(message Message) Query {
	return Query{
		MaxTime: message.Right,
		MinTime: message.Left,
	}
}
