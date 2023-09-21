package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	Conn   net.Conn
	ReadCh chan []byte
}

// NewClient creates a new TCP client
func NewClient(uri string) (*Client, error) {
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	client := &Client{
		Conn:   conn,
		ReadCh: make(chan []byte),
	}

	go client.readLoop()

	return client, nil
}

// readLoop - reads data from the connection and send it to the channel
func (c *Client) readLoop() {
	reader := bufio.NewReader(c.Conn)

	for {
		msg, err := reader.ReadBytes('\n')
		if len(msg) > 0 {
			c.ReadCh <- msg
		}

		if err != nil {
			close(c.ReadCh)
			return
		}
	}
}

// Close - closes the connection
func (c *Client) Close() error {
	return c.Conn.Close()
}
