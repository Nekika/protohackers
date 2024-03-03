package lib

import (
	"slices"
	"sync"
)

type Room struct {
	mu      sync.Mutex
	clients []*Client
}

func NewRoom() *Room {
	return &Room{
		mu:      sync.Mutex{},
		clients: make([]*Client, 0),
	}
}

func (r *Room) Add(client *Client) {
	defer r.mu.Unlock()
	r.mu.Lock()

	r.clients = append(r.clients, client)
}

func (r *Room) Clients() []*Client {
	defer r.mu.Unlock()
	r.mu.Lock()

	return r.clients
}

func (r *Room) Remove(client *Client) {
	defer r.mu.Unlock()
	r.mu.Lock()

	r.clients = slices.DeleteFunc(r.clients, func(c *Client) bool {
		return c == client
	})
}

func (r *Room) Send(data []byte) {
	defer r.mu.Unlock()
	r.mu.Lock()

	for _, client := range r.clients {
		_, _ = client.Conn.Write(data)
	}
}
