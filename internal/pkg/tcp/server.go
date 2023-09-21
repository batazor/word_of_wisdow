package tcp

import (
	"context"
	"net"

	"github.com/batazor/word_of_wisdow/internal/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	ReadCh chan []byte
	log    *logger.Logger
}

// NewServer creates a new TCP server
func NewServer(ctx context.Context, uri string, log *logger.Logger) (*Server, error) {
	s := &Server{
		log:    log,
		ReadCh: make(chan []byte),
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
			if opErr, ok := err.(*net.OpError); ok && opErr.Op == "accept" {
				if opErr.Err.Error() == "use of closed network connection" {
					s.log.Info("Listener has been closed, stopping accept loop.")
					return
				}
			}

			s.log.Error("s.Listener.Accept", zap.Error(err))
			return
		}

		// Start reading data from the connection
		go s.readLoop(conn)
	}
}

// readLoop - reads data from the connection
func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		// Read data from the connection
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				return
			}

			s.log.Error("s.Conn.Read", zap.Error(err))
			return
		}

		// Send data to the channel
		s.ReadCh <- buf[:n]
	}
}
