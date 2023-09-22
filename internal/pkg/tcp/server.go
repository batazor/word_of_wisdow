package tcp

import (
	"context"
	"errors"
	"io"
	"net"
	"sync"

	"github.com/batazor/word_of_wisdom/internal/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	mu      sync.Mutex
	clients map[net.Conn]struct{}
	ReadCh  chan []byte
	log     *logger.Logger
}

// NewServer creates a new TCP server
func NewServer(ctx context.Context, uri string, log *logger.Logger) (*Server, error) {
	s := &Server{
		log:     log,
		clients: make(map[net.Conn]struct{}),
		ReadCh:  make(chan []byte),
	}

	// Create a new TCP listener
	listener, err := net.Listen("tcp", uri)
	if err != nil {
		return nil, err
	}

	// Accept connections
	go s.acceptConnections(ctx, listener)

	return s, nil
}

// acceptConnections - accepts connections
func (s *Server) acceptConnections(ctx context.Context, listener net.Listener) {
	defer close(s.ReadCh)

	go func() {
		<-ctx.Done()
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				s.log.Info("Listener has been closed, stopping accept loop.")
				return
			}

			s.log.Error("s.Listener.Accept", zap.Error(err))
			return
		}

		// Add to the list of connections
		s.mu.Lock()
		s.clients[conn] = struct{}{}
		s.mu.Unlock()

		// Start reading data from the connection
		go s.readLoop(conn)
	}
}

// readLoop - reads data from the connection
func (s *Server) readLoop(conn net.Conn) {
	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()

	buf := make([]byte, 1024)

	for {
		// Read data from the connection
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}

			s.log.Error("s.Conn.Read", zap.Error(err))
			return
		}

		// Send data to the channel
		s.ReadCh <- buf[:n]
	}
}

// Send - sends data to all connections
func (s *Server) Send(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.clients {
		_, err := conn.Write(data)
		if err != nil {
			s.log.Error("conn.Write", zap.Error(err))
		}
	}

	return nil
}
