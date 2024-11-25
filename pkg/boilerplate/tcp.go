package boilerplate

import (
	"fmt"
	"log"
	"net"
)

type TCPListener struct {
	Handler      func(*net.TCPConn) error
	ErrorHandler func(error, *net.TCPConn)
}

func (h TCPListener) Listen(addr string) {
	if h.Handler == nil {
		log.Fatal("No callback provided")
	}

	log.Println("Trying to create a new listener")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to create a listener: %s\n", err.Error())
	}
	defer listener.Close()

	log.Printf("Listener created and listening at %s\n", addr)

	for {
		log.Println("Waiting for a new connection")

		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept new connection: %s\n", err.Error())
		}

		log.Printf("Accepted new connection from %s\n", conn.RemoteAddr())

		go func() {
			defer conn.Close()
			if err := h.Handler(conn.(*net.TCPConn)); err != nil {
				log.Printf("[%s] %s", conn.RemoteAddr(), err.Error())
				if h.ErrorHandler != nil {
					h.ErrorHandler(err, conn.(*net.TCPConn))
				}
			}
		}()
	}
}

// ListenTCP creates a new TCPListener with the given handler
// then starts it on the given address.
func ListenTCP(addr string, handler func(*net.TCPConn) error) {
	listener := TCPListener{Handler: handler}
	listener.Listen(addr)
}

// MustDial connects to the address on the named network and panics if it fails.
func MustDial(network, addr string) net.Conn {
	conn, err := net.Dial(network, addr)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

// MustDialTCP connects to the adress on the TCP network and panic if it fails.
func MustDialTCP(addr string) *net.TCPConn {
	return MustDial("tcp", addr).(*net.TCPConn)
}

func WriteStringf(conn net.Conn, format string, args ...any) (int, error) {
	message := fmt.Sprintf(format, args...)
	return conn.Write([]byte(message))
}
