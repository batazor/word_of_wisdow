package tcp

import (
	"fmt"
	"net"
)

type Client struct {
	Conn net.Conn
}

// NewClient creates a new TCP client
func NewClient(uri string) (*Client, error) {
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return &Client{Conn: conn}, nil
}

// Close - closes the connection
func (c *Client) Close() error {
	return c.Conn.Close()
}
