package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn, eventch chan<- Event, messagech chan<- Message) {
	defer conn.Close()

	var (
		client  = &Client{Conn: conn}
		err     error
		scanner = bufio.NewScanner(conn)
	)

	if _, err = conn.Write([]byte("Input username:\n")); err != nil {
		log.Printf("(%s) failed to write to conn: %s\n", conn.RemoteAddr().String(), err)
		return
	}

	if !scanner.Scan() {
		log.Printf("(%s) failed to read from conn: %s\n", conn.RemoteAddr().String(), scanner.Err())
		return
	}

	client.Username = scanner.Text()

	eventch <- Event{
		Kind:   EventKindClientJoined,
		Client: client,
	}

	for scanner.Scan() {
		messagech <- Message{
			Client: client,
			Text:   fmt.Sprintf("[%s] %s\n", client.Username, scanner.Text()),
		}
	}

	log.Printf("(%s) err scanning: %s\n", conn.RemoteAddr().String(), scanner.Err())

	eventch <- Event{
		Kind:   EventKindClientLeft,
		Client: client,
	}
}

func Serve(listener net.Listener) error {
	var (
		eventch   = make(chan Event)
		messagech = make(chan Message)
		room      = NewRoom()
	)

	go func() {
		for event := range eventch {
			switch event.Kind {
			case EventKindClientJoined:
				usernames := Map(room.Clients(), func(client *Client, _ int) string { return client.Username })
				_, _ = event.Client.Conn.Write([]byte(fmt.Sprintf("* Connected users : %s\n", strings.Join(usernames, ", "))))
				room.Send([]byte(fmt.Sprintf("* %s connected.\n", event.Client.Username)))
				room.Add(event.Client)
			case EventKindClientLeft:
				room.Send([]byte(fmt.Sprintf("* %s disconnected.\n", event.Client.Username)))
				room.Remove(event.Client)
				_ = event.Client.Conn.Close()
			}

		}
	}()

	go func() {
		for message := range messagech {
			for _, client := range room.Clients() {
				if client == message.Client {
					continue
				}

				_, _ = client.Conn.Write([]byte(message.Text))
			}
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go handle(conn, eventch, messagech)
	}
}
