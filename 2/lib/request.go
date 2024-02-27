package lib

type Insert struct {
	Amount    int32
	Timestamp int32
}

func InsertFromMessage(message Message) Insert {
	return Insert{
		Amount:    message.Right(),
		Timestamp: message.Left(),
	}
}

type Query struct {
	MinTime int32
	MaxTime int32
}

func QueryFromMessage(message Message) Query {
	return Query{
		MaxTime: message.Right(),
		MinTime: message.Left(),
	}
}
