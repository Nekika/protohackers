package lib

import "net"

type Client struct {
	Conn     net.Conn
	Username string
}
